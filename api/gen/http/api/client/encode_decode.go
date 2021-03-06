// Code generated by goa v3.2.3, DO NOT EDIT.
//
// api HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/sm43/oauth/api/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	api "github.com/sm43/oauth/api/gen/api"
	goahttp "goa.design/goa/v3/http"
)

// BuildAuthenticateRequest instantiates a HTTP request object with method and
// path set to call the "api" service "Authenticate" endpoint
func (c *Client) BuildAuthenticateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AuthenticateAPIPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api", "Authenticate", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAuthenticateRequest returns an encoder for requests sent to the api
// Authenticate server.
func EncodeAuthenticateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*api.AuthenticatePayload)
		if !ok {
			return goahttp.ErrInvalidType("api", "Authenticate", "*api.AuthenticatePayload", v)
		}
		values := req.URL.Query()
		values.Add("code", p.Code)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeAuthenticateResponse returns a decoder for responses returned by the
// api Authenticate endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeAuthenticateResponse may return the following errors:
//	- "invalid-code" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeAuthenticateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AuthenticateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api", "Authenticate", err)
			}
			err = ValidateAuthenticateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api", "Authenticate", err)
			}
			res := NewAuthenticateResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body AuthenticateInvalidCodeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api", "Authenticate", err)
			}
			err = ValidateAuthenticateInvalidCodeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api", "Authenticate", err)
			}
			return nil, NewAuthenticateInvalidCode(&body)
		case http.StatusInternalServerError:
			var (
				body AuthenticateInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api", "Authenticate", err)
			}
			err = ValidateAuthenticateInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api", "Authenticate", err)
			}
			return nil, NewAuthenticateInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api", "Authenticate", resp.StatusCode, string(body))
		}
	}
}

// BuildDetailsRequest instantiates a HTTP request object with method and path
// set to call the "api" service "Details" endpoint
func (c *Client) BuildDetailsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DetailsAPIPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api", "Details", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDetailsRequest returns an encoder for requests sent to the api Details
// server.
func EncodeDetailsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*api.DetailsPayload)
		if !ok {
			return goahttp.ErrInvalidType("api", "Details", "*api.DetailsPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDetailsResponse returns a decoder for responses returned by the api
// Details endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDetailsResponse may return the following errors:
//	- "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "invalid-token" (type *goa.ServiceError): http.StatusUnauthorized
//	- error: internal error
func DecodeDetailsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DetailsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api", "Details", err)
			}
			err = ValidateDetailsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api", "Details", err)
			}
			res := NewDetailsUserOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DetailsInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api", "Details", err)
			}
			err = ValidateDetailsInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api", "Details", err)
			}
			return nil, NewDetailsInternalError(&body)
		case http.StatusUnauthorized:
			var (
				body DetailsInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api", "Details", err)
			}
			err = ValidateDetailsInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api", "Details", err)
			}
			return nil, NewDetailsInvalidToken(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api", "Details", resp.StatusCode, string(body))
		}
	}
}
