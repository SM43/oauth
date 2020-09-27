// Code generated by goa v3.2.3, DO NOT EDIT.
//
// api HTTP server types
//
// Command:
// $ goa gen github.com/sm43/oauth/api/design

package server

import (
	api "github.com/sm43/oauth/api/gen/api"
	goa "goa.design/goa/v3/pkg"
)

// AuthenticateResponseBody is the type of the "api" service "Authenticate"
// endpoint HTTP response body.
type AuthenticateResponseBody struct {
	// JWT
	Token string `form:"token" json:"token" xml:"token"`
}

// DetailsResponseBody is the type of the "api" service "Details" endpoint HTTP
// response body.
type DetailsResponseBody struct {
	// ID is the unique id of User
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of User
	Name string `form:"name" json:"name" xml:"name"`
	// GitHub ID
	GithubID string `form:"githubID" json:"githubID" xml:"githubID"`
}

// AuthenticateInvalidCodeResponseBody is the type of the "api" service
// "Authenticate" endpoint HTTP response body for the "invalid-code" error.
type AuthenticateInvalidCodeResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// AuthenticateInternalErrorResponseBody is the type of the "api" service
// "Authenticate" endpoint HTTP response body for the "internal-error" error.
type AuthenticateInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DetailsInternalErrorResponseBody is the type of the "api" service "Details"
// endpoint HTTP response body for the "internal-error" error.
type DetailsInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DetailsInvalidTokenResponseBody is the type of the "api" service "Details"
// endpoint HTTP response body for the "invalid-token" error.
type DetailsInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewAuthenticateResponseBody builds the HTTP response body from the result of
// the "Authenticate" endpoint of the "api" service.
func NewAuthenticateResponseBody(res *api.AuthenticateResult) *AuthenticateResponseBody {
	body := &AuthenticateResponseBody{
		Token: res.Token,
	}
	return body
}

// NewDetailsResponseBody builds the HTTP response body from the result of the
// "Details" endpoint of the "api" service.
func NewDetailsResponseBody(res *api.User) *DetailsResponseBody {
	body := &DetailsResponseBody{
		ID:       res.ID,
		Name:     res.Name,
		GithubID: res.GithubID,
	}
	return body
}

// NewAuthenticateInvalidCodeResponseBody builds the HTTP response body from
// the result of the "Authenticate" endpoint of the "api" service.
func NewAuthenticateInvalidCodeResponseBody(res *goa.ServiceError) *AuthenticateInvalidCodeResponseBody {
	body := &AuthenticateInvalidCodeResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAuthenticateInternalErrorResponseBody builds the HTTP response body from
// the result of the "Authenticate" endpoint of the "api" service.
func NewAuthenticateInternalErrorResponseBody(res *goa.ServiceError) *AuthenticateInternalErrorResponseBody {
	body := &AuthenticateInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDetailsInternalErrorResponseBody builds the HTTP response body from the
// result of the "Details" endpoint of the "api" service.
func NewDetailsInternalErrorResponseBody(res *goa.ServiceError) *DetailsInternalErrorResponseBody {
	body := &DetailsInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDetailsInvalidTokenResponseBody builds the HTTP response body from the
// result of the "Details" endpoint of the "api" service.
func NewDetailsInvalidTokenResponseBody(res *goa.ServiceError) *DetailsInvalidTokenResponseBody {
	body := &DetailsInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAuthenticatePayload builds a api service Authenticate endpoint payload.
func NewAuthenticatePayload(code string) *api.AuthenticatePayload {
	v := &api.AuthenticatePayload{}
	v.Code = code

	return v
}

// NewDetailsPayload builds a api service Details endpoint payload.
func NewDetailsPayload(token string) *api.DetailsPayload {
	v := &api.DetailsPayload{}
	v.Token = token

	return v
}