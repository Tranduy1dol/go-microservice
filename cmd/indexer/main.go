package main

import (
	"context"
	"log"

	"github.com/Tranduy1dol/kotoba-press-core/config"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/grpc"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/mongo"
	searchpb "github.com/Tranduy1dol/kotoba-press-core/proto/grpc_service/v1"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	_, db, err := mongo.NewClient(context.Background(), cfg.MongoDB.URI, cfg.MongoDB.Database)
	if err != nil {
		log.Fatalf("failed to connect mongo: %v", err)
	}

	searchClient, err := grpc.NewSearchClient(cfg.GRPC.SearchEngineAddr)
	if err != nil {
		log.Fatalf("failed to connect search engine: %v", err)
	}
	defer func() { _ = searchClient.Close() }()

	wordRepo := mongo.NewWordRepository(db)
	words, _, err := wordRepo.List(context.Background(), 100000, 0)
	if err != nil {
		log.Fatalf("failed to list words: %v", err)
	}

	log.Printf("found %d words to index", len(words))

	docs := make([]*searchpb.IndexDocumentRequest, len(words))
	for i, w := range words {
		title, text := w.SearchText()
		docs[i] = &searchpb.IndexDocumentRequest{
			DocId:       w.ID,
			Title:       title,
			Text:        text,
			ContentType: searchpb.ContentType_CONTENT_TYPE_WORD,
			Level:       int32(w.JLPT),
		}
	}

	indexed, failed, err := searchClient.BulkIndex(context.Background(), docs)
	if err != nil {
		log.Fatalf("bulk index failed: %v", err)
	}

	log.Printf("Indexed %d, Failed %d", indexed, failed)
}
