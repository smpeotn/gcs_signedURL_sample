package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	bucket := "public-image-smpeotn"
	object := "IMG_3766.JPG"
	expiresTime := 3 * time.Minute

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(fmt.Errorf("storage.NewClient: %w", err))
		return
	}
	defer client.Close()

	u, err := client.Bucket(bucket).SignedURL(object, &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  http.MethodGet,
		Expires: time.Now().Add(expiresTime),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err))
		return
	}
	fmt.Println(u)
}
