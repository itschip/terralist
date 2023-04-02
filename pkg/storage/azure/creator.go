package azure

import (
	"fmt"
	"terralist/pkg/storage"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type Creator struct{}

func (t *Creator) New(config storage.Configurator) (storage.Resolver, error) {
	cfg, ok := config.(*Config)

	if !ok {
		return nil, fmt.Errorf("unsupported configurator")
	}

	var creds *azidentity.DefaultAzureCredential = nil

	if !cfg.DefaultAzureCredential {
		var err error
		creds, err = azidentity.NewClientSecretCredential(cfg.TenantID, cfg.ClientID, cfg.ClientSecret, nil)
		if err != nil {
			return nil, err
		}
	}

	client, err := azblob.NewClient(cfg.Account, creds, nil)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		Container: cfg.Container,
		Account:   cfg.Account,
		Session:   client,
	}, nil
}
