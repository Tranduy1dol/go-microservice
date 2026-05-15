package grpc

import (
	"context"
	"log"

	searchpb "github.com/Tranduy1dol/kotoba-press-core/proto/grpc_service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

	log.Printf("[gRPC] Connected to search engine at %s", addr)
	return &SearchClient{
		client: searchpb.NewSearchServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *SearchClient) Close() error {
	return c.conn.Close()
}

func (c *SearchClient) Search(ctx context.Context, query string, limit int) (*searchpb.SearchResponse, error) {
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
	req := &searchpb.IndexDocumentRequest{
		DocId:       docID,
		Title:       title,
		Text:        text,
		ContentType: contentType,
		Level:       int32(level),
	}

	_, err := c.client.IndexDocument(ctx, req)
	return err
}

func (c *SearchClient) BulkIndex(ctx context.Context, docs []*searchpb.IndexDocumentRequest) (uint32, uint32, error) {
	stream, err := c.client.BulkIndex(ctx)
	if err != nil {
		return 0, 0, err
	}

	for _, doc := range docs {
		if err := stream.Send(doc); err != nil {
			return 0, 0, err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return 0, 0, err
	}

	return resp.IndexedCount, resp.FailedCount, nil
}

func (c *SearchClient) GetStats(ctx context.Context) (uint32, float64, error) {
	resp, err := c.client.GetStats(ctx, &searchpb.GetStatsRequest{})
	if err != nil {
		return 0, 0, err
	}

	return resp.TotalDocuments, resp.AvgDocumentLength, nil
}
