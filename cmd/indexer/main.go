package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Tranduy1dol/kotoba-press-core/config"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/grpc"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/mongo"
	"github.com/Tranduy1dol/kotoba-press-core/internal/logger"
	searchpb "github.com/Tranduy1dol/kotoba-press-core/proto/grpc_service/v1"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Init(cfg.Server.Env, cfg.Server.LogLevel)

	_, db, err := mongo.NewClient(context.Background(), cfg.MongoDB.URI, cfg.MongoDB.Database)
	if err != nil {
		slog.Error("failed to connect mongo", "error", err)
		os.Exit(1)
	}

	searchClient, err := grpc.NewSearchClient(cfg.GRPC.SearchEngineAddr)
	if err != nil {
		slog.Error("failed to connect search engine", "error", err)
		os.Exit(1)
	}
	defer func() { _ = searchClient.Close() }()

	wordRepo := mongo.NewWordRepository(db)
	words, _, err := wordRepo.List(context.Background(), 100000, 0)
	if err != nil {
		slog.Error("failed to list words", "error", err)
		os.Exit(1)
	}

	slog.Info("found words to index", "count", len(words))

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
		slog.Error("bulk index failed", "error", err)
		os.Exit(1)
	}

	slog.Info("bulk indexing completed", "indexed", indexed, "failed", failed)
}
