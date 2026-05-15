package grpc

import (
	"context"

	"github.com/Tranduy1dol/kotoba-press-core/internal/logger"
	searchpb "github.com/Tranduy1dol/kotoba-press-core/proto/grpc_service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logger.New(logger.ComponentGRPC)

type SearchClient struct {
	client searchpb.SearchServiceClient
	conn   *grpc.ClientConn
}

func NewSearchClient(addr string) (*SearchClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	log.Info("connected to search engine", "addr", addr)
	return &SearchClient{
		client: searchpb.NewSearchServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *SearchClient) Close() error {
	return c.conn.Close()
}

func (c *SearchClient) Search(ctx context.Context, query string, limit int) (*searchpb.SearchResponse, error) {
	log.Debug("searching", "query", query, "limit", limit)

	req := &searchpb.SearchRequest{
		Query: query,
		TopK:  uint32(limit),
		Filter: &searchpb.ContentFilter{
			ContentTypes: []searchpb.ContentType{searchpb.ContentType_CONTENT_TYPE_WORD},
		},
	}

	return c.client.Search(ctx, req)
}

func (c *SearchClient) IndexDocument(ctx context.Context, docID, title, text string, contentType searchpb.ContentType, level int) error {
	log.Debug("indexing document", "doc_id", docID, "content_type", contentType, "level", level)

	req := &searchpb.IndexDocumentRequest{
		DocId:       docID,
		Title:       title,
		Text:        text,
		ContentType: contentType,
		Level:       int32(level),
	}

	_, err := c.client.IndexDocument(ctx, req)
	if err != nil {
		log.Error("failed to index document", "doc_id", docID, "error", err)
	}
	return err
}

func (c *SearchClient) BulkIndex(ctx context.Context, docs []*searchpb.IndexDocumentRequest) (uint32, uint32, error) {
	log.Info("bulk indexing documents", "count", len(docs))

	stream, err := c.client.BulkIndex(ctx)
	if err != nil {
		log.Error("failed to open bulk index stream", "error", err)
		return 0, 0, err
	}

	for _, doc := range docs {
		if err := stream.Send(doc); err != nil {
			log.Error("failed to send document in bulk stream", "doc_id", doc.DocId, "error", err)
			return 0, 0, err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Error("failed to close bulk index stream", "error", err)
		return 0, 0, err
	}

	log.Info("bulk indexing completed", "indexed", resp.IndexedCount, "failed", resp.FailedCount)
	return resp.IndexedCount, resp.FailedCount, nil
}

func (c *SearchClient) GetStats(ctx context.Context) (uint32, float64, error) {
	log.Debug("fetching search engine stats")

	resp, err := c.client.GetStats(ctx, &searchpb.GetStatsRequest{})
	if err != nil {
		log.Error("failed to fetch search engine stats", "error", err)
		return 0, 0, err
	}

	return resp.TotalDocuments, resp.AvgDocumentLength, nil
}
