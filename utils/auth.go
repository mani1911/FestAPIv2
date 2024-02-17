package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/delta/FestAPI/config"
)

type DAuthToken struct {
	AccessToken string
	IDToken     string
}

type DAuthUser struct {
	Email  string
	Name   string
	Gender string
	Phone  string
}

func GetDAuthToken(code string, site string) (*DAuthToken, error) {
	tokenEndpoint := "https://auth.delta.nitt.edu/api/oauth/token"
	values := url.Values{}
	if site == "tshirt" {
		values.Add("client_id", config.TshirtDAuthClientID)
		values.Add("client_secret", config.TshirtDAuthClientSecret)
		values.Add("redirect_uri", config.TshirtDAuthCallbackURL)
	} else {
		values.Add("client_id", config.DAuthClientID)
		values.Add("client_secret", config.DAuthClientSecret)
		values.Add("redirect_uri", config.DAuthCallbackURL)
	}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	query := values.Encode()

	req, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var DAuthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody, &DAuthTokenRes); err != nil {
		return nil, err
	}

	tokenBody := &DAuthToken{
		AccessToken: DAuthTokenRes["access_token"].(string),
		IDToken:     DAuthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

func GetDAuthUser(accessToken string) (*DAuthUser, error) {
	userEndpoint := "https://auth.delta.nitt.edu/api/resources/user"
	var TokenRes map[string]interface{}

	req, err := http.NewRequest("POST", userEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resBody, &TokenRes); err != nil {
		return nil, err
	}

	DAuthUser := &DAuthUser{
		Email:  TokenRes["email"].(string),
		Name:   TokenRes["name"].(string),
		Gender: strings.ToUpper(TokenRes["gender"].(string)),
		Phone:  TokenRes["phoneNumber"].(string),
	}

	return DAuthUser, nil
}
