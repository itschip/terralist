package azure

import (
	"context"
	"fmt"
	"terralist/pkg/storage"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type Resolver struct {
	Container string
	Account   string
	Session   *azblob.Client
}

func (r *Resolver) Store(in *storage.StoreInput) (string, error) {
	key := fmt.Sprintf("%s/%s", in.KeyPrefix, in.FileName)

	_, err := r.Session.UploadBuffer(context.TODO(), r.Container, key, in.Content, nil)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return key, nil
}

func (r *Resolver) Find(key string) (string, error) {
	return "", nil
}

func (r *Resolver) Purge(key string) error {
	return nil
}
