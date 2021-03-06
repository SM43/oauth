// Code generated by goa v3.2.3, DO NOT EDIT.
//
// api HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sm43/oauth/api/design

package client

import (
	api "github.com/sm43/oauth/api/gen/api"
)

// BuildAuthenticatePayload builds the payload for the api Authenticate
// endpoint from CLI flags.
func BuildAuthenticatePayload(apiAuthenticateCode string) (*api.AuthenticatePayload, error) {
	var code string
	{
		code = apiAuthenticateCode
	}
	v := &api.AuthenticatePayload{}
	v.Code = code

	return v, nil
}

// BuildDetailsPayload builds the payload for the api Details endpoint from CLI
// flags.
func BuildDetailsPayload(apiDetailsToken string) (*api.DetailsPayload, error) {
	var token string
	{
		token = apiDetailsToken
	}
	v := &api.DetailsPayload{}
	v.Token = token

	return v, nil
}
