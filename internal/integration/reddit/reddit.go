package reddit

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TODO: parameterize these urls, include in startup flags or env
const authorizationUrl = "https://www.reddit.com/api/v1/access_token"
const apiUrl = "https://oauth.reddit.com"
const subreddit = "/r/gaming"

// returns a userless bearer token
// TODO: implement retry logic instead of crashing
// TODO: define interface for integrations
func GetAuthToken(clientId string, clientSecret string) string {
	fmt.Print("getting auth token using values: " + clientId + " " + clientSecret + "\n")

	data := url.Values{}

	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, authorizationUrl, strings.NewReader(data.Encode()))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "DOskvig-api/0.1")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Authorization", "Basic "+basicAuth(clientId, clientSecret))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	var respData = struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&respData)

	if err != nil {
		log.Fatal(err)
	}

	return respData.AccessToken
}

func GetPostsForSubreddit(token string, before string, subreddit string) (*SubredditListing, error) {
	//TODO: implement response object interface
	var data SubredditListing
	//TODO: replace with URL builder
	url := apiUrl + subreddit + "/new"
	if before != "" {
		url = url + "?before=" + before
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("User-Agent", "DOskvig-api/0.1")

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		if res.StatusCode != http.StatusTooManyRequests {
			log.Fatal(err)
		}
	}

	return &data, err
}

// TODO: move to utils
func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
