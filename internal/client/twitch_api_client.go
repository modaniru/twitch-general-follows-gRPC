package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

// TODO documentation
// TODO refactor
const UserInfoURI = "https://api.twitch.tv/helix/users"
const UserGetFollowListURI = "https://api.twitch.tv/helix/users/follows"
const GetTokenURI = "https://id.twitch.tv/oauth2/token"
const ValidateTokenURI = "https://id.twitch.tv/oauth2/validate"

type Queries struct {
	clientId     string
	clientSecret string

	client *httpclient.Client
	token  string
}

func NewQueries(clientId, clientSecret string) *Queries {
	timeout := 10 * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	queries := Queries{
		client:       client,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
	token, err := queries.GetOauthToken()
	if err != nil {
		log.Fatal(err)
	}
	queries.token = token.AccessToken
	return &queries
}

func (q *Queries) GetOauthToken() (*OauthToken, error) {
	uri := fmt.Sprintf("%s?client_id=%s&client_secret=%s&grant_type=client_credentials", GetTokenURI, q.clientId, q.clientSecret)
	res, err := q.client.Post(uri, nil, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		message := fmt.Sprintf("%d status code", res.StatusCode)
		return nil, errors.New(message)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response OauthToken
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}

func (q *Queries) IsValid() (*ValidToken, error) {
	uri := ValidateTokenURI
	token := "OAuth " + q.token
	header := http.Header{}
	header.Add("Authorization", token)
	res, err := q.client.Get(uri, header)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		message := fmt.Sprintf("%d status code", res.StatusCode)
		return nil, errors.New(message)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response ValidToken
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}

func (q *Queries) GetUsersInfo(data []string, dataType string) ([]UserInfo, error) {
	iterations := len(data)/100 + 1
	channel := make(chan *UserCollection, 100)
	for len(data) > 100 {
		temp := data[:100]
		go func() {
			u, err := q.getUsersInfoRoutine(temp, dataType)
			if err != nil {
				channel <- nil
				return
			}
			channel <- u
		}()
		data = data[100:]
	}
	go func() {
		u, err := q.getUsersInfoRoutine(data, dataType)
		if err != nil {
			channel <- nil
			return
		}
		channel <- u
	}()
	res := []UserInfo{}
	for i := 0; i < iterations; i++ {
		userCollection := <-channel
		if userCollection == nil {
			return nil, errors.New("invalid user(s)")
		}
		res = append(res, userCollection.Data...)
	}
	return res, nil
}

func (q *Queries) getUsersInfoRoutine(users []string, t string) (*UserCollection, error) {
	uri := UserInfoURI
	for i, v := range users {
		symb := "&"
		if i == 0 {
			symb = "?"
		}
		uri = fmt.Sprintf("%s%s%s=%s", uri, symb, t, v)
	}
	header := http.Header{}
	token := "Bearer " + q.getToken()
	header.Add("Authorization", token)
	header.Add("Client-Id", q.clientId)
	res, err := q.client.Get(uri, header)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d status code", res.StatusCode)

	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response UserCollection
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}

func (q *Queries) getToken() string {
	_, err := q.IsValid()
	if err != nil {
		token, err := q.GetOauthToken()
		if err == nil {
			return ""
		}
		q.token = token.AccessToken
	}
	return q.token
}

// go routine
func (q *Queries) GetFollows(id string) ([]FollowInfo, error) {
	uri := fmt.Sprintf("%s?from_id=%s&first=%d", UserGetFollowListURI, id, 100)
	response, err := q.getFollowsWithoutPagination(uri)
	if err != nil {
		return nil, err
	}
	var result []FollowInfo
	result = append(result, response.Data...)

	for response.Pagination.Cursor != "" {
		uri2 := fmt.Sprintf("%s&after=%s", uri, response.Pagination.Cursor)
		response, err = q.getFollowsWithoutPagination(uri2)
		if err != nil {
			return nil, err
		}
		result = append(result, response.Data...)
	}
	return result, nil
}

func (q *Queries) getFollowsWithoutPagination(uri string) (*FollowsCollection, error) {
	header := http.Header{}
	token := "Bearer " + q.getToken()
	header.Add("Authorization", token)
	header.Add("Client-Id", q.clientId)
	res, err := q.client.Get(uri, header)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response FollowsCollection
	json.Unmarshal([]byte(string(body)), &response)
	return &response, nil
}
