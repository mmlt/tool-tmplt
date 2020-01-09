package thycotic

import (
	"fmt"
)

// Login
func Login(url, username, passw, domain string) (*SSWebServiceSoap, string, error) {
	client := NewSSWebServiceSoap(fmt.Sprint(url, "/SecretServer/webservices/sswebservice.asmx"), true, nil) //TODO change true in false or make it a flag

	ar, err := client.Authenticate(
		&Authenticate{
			Username:     username,
			Password:     passw,
			Organization: "",
			Domain:       domain,
		})
	if err != nil {
		return nil, "", err
	}

	token := ar.AuthenticateResult.Token
	if token == "" {
		return nil, "", fmt.Errorf("unauthorized")
	}

	return client, token, nil
}

// Get returns a secret value from a Thycotic secret server based on secret id and item field name.
func Get(id int32, item string, client *SSWebServiceSoap, token string) (string, error) {
	response, err := client.GetSecret(&GetSecret{Token: token, SecretId: id})
	if err != nil {
		return "<unknown-secret>", fmt.Errorf("no secret %d: %s", id, err)
	}
	if response.GetSecretResult.Secret == nil {
		return "<unknown-secret>", fmt.Errorf("no secret %d", id)
	}

	for _, si := range response.GetSecretResult.Secret.Items.SecretItem {
		if si.FieldName == item {
			return si.Value, nil
		}
	}
	return "<unknown-secret-field>", fmt.Errorf("secret %d has no field %s", id, item)
}
