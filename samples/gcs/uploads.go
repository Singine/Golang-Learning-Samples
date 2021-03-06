package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	//projectID := "gcs-gubao"
	bucket := "gubao_bucket_test"
	object := "article_1.png"
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	f, err := os.Open("C:/Users/DZH/go/Golang-Learning-Samples/samples/gcs/upload/files/article_1.png")
	if err != nil {
		fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Errorf("Writer.Close: %v", err)
	}

}
