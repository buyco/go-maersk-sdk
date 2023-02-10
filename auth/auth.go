package auth

import (
	"context"
	"time"

	"golang.org/x/xerrors"
)

const (
	authPath  = "/oauth2/access_token"
	grantType = "client_credentials"
)

type Result struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	IdToken     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type Client struct {
	httpClient   HTTPClient
	url          string
	clientId     string
	clientSecret string
}

func NewClient(httpClient HTTPClient, url, clientId, clientSecret string) *Client {
	return &Client{
		httpClient:   httpClient,
		url:          url,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (c Client) Token(ctx context.Context) (string, time.Duration, error) {
	result := Result{}
	errRes := ErrorResponse{}

	r, err := c.httpClient.NewRequest().
		SetResult(&result).
		SetError(&errRes).
		SetHeaders(c.buildHeaders()).
		SetContext(ctx).
		SetFormData(c.buildParams()).
		Post(c.url + authPath)
	if err != nil {
		return "", 0, err
	}

	if r.IsError() {
		return "", 0, xerrors.New(errRes.Error)
	} else if r.IsSuccess() {
		return result.AccessToken, time.Duration(result.ExpiresIn) * time.Second, nil
	}

	return "", 0, xerrors.Errorf("unknown error. status code %d; body: %s", r.GetStatusCode(), r.String())
}

func (a Client) buildParams() map[string]string {
	return map[string]string{
		"grant_type":    grantType,
		"client_id":     a.clientId,
		"client_secret": a.clientSecret,
	}
}

func (a Client) buildHeaders() map[string]string {
	return map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
}
