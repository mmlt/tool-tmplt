package azkv

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
)

// Azure Key Vault
// https://docs.microsoft.com/en-us/azure/key-vault/about-keys-secrets-and-certificates

// Authenticate via environment variables AZURE_TENANT_ID, AZURE_CLIENT_ID, AZURE_CLIENT_SECRET
// https://docs.microsoft.com/en-us/azure/go/azure-sdk-go-authorization#use-environment-based-authentication

func Login(url string) (*keyvault.BaseClient, error) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, err
	}

	basicClient := keyvault.New()
	basicClient.Authorizer = authorizer

	return &basicClient, nil
}

func Get(name string, client *keyvault.BaseClient, url string) (string, error) {
	secretResp, err := client.GetSecret(context.Background(), url, name, "")
	if err != nil {
		return "", fmt.Errorf("no secret %s: %v", name, err)
	}
	return *secretResp.Value, nil
}
