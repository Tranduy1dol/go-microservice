package main

import (
	"context"
	"log"
	"os"

	"github.com/Tranduy1dol/kotoba-press-core/config"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/jmdict"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/mongo"
	"github.com/spf13/cobra"
)

func main() {
	var filePath string

	cmd := &cobra.Command{
		Use:   "importer",
		Short: "Import JMdict data into MongoDB",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}

			client, db, err := mongo.NewClient(context.Background(), cfg.MongoDB.URI, cfg.MongoDB.Database)
			if err != nil {
				return err
			}
			defer func() { _ = mongo.Close(client) }()

			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer func() { _ = file.Close() }()

			log.Printf("parsing JMdict file...")
			words, err := jmdict.Parse(file)
			if err != nil {
				return err
			}

			wordRepo := mongo.NewWordRepository(db)
			batchSize := 500
			imported := 0

			for i := 0; i < len(words); i += batchSize {
				end := min(i+batchSize, len(words))
				batch := words[i:end]

				count, err := wordRepo.BulkCreate(context.Background(), batch)
				if err != nil {
					log.Printf("batch %d-%d failed: %v", i, end, err)
					continue
				}

				imported += int(count)
				log.Printf("imported %d/%d words...", imported, len(words))
			}

			log.Printf("done: %d words imported", imported)
			return nil
		},
	}

	cmd.Flags().StringVar(&filePath, "file", "", "path to jmdict JSON file")
	_ = cmd.MarkFlagRequired("file")
	_ = cmd.Execute()
}
