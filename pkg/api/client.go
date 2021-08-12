package api

import (
	"encoding/json"
	"fmt"
	"github.com/KaanTolunayKilic/golang-semgrep-test/pkg/model"
	"log"
	"net/http"
	"net/url"
)

const apiURL string = "https://jsonplaceholder.typicode.com"

// Client struct to make request
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

// NewClient creats new client
func NewClient(httpClient *http.Client) Client {
	url, _ := url.Parse(apiURL)
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return Client{httpClient: httpClient, BaseURL: url}
}

// ListUserPosts requests all posts from user
func (c *Client) ListUserPosts(userID int) ([]model.Post, error) {
	req, err := c.newRequest(http.MethodGet, "/posts")
	if err != nil {
		return nil, err
	}
	addQuery(req, "userId", fmt.Sprintf("%d", userID))

	var posts []model.Post
	_, err = c.do(req, &posts)
	return posts, err
}

// ListPostComments requests all comments for post
func (c *Client) ListPostComments(postID int) ([]model.Comment, error) {
	req, err := c.newRequest(http.MethodGet, "/comments")
	if err != nil {
		return nil, err
	}
	addQuery(req, "postId", fmt.Sprintf("%d", postID))

	var comments []model.Comment
	_, err = c.do(req, &comments)
	return comments, err
}

func (c *Client) Comments(p *model.Post) []model.Comment {
	comments, err := c.ListPostComments(p.ID)
	if err != nil {
		log.Fatalf("Could not request comments for post with id %d: %s", p.ID, err)
	}
	return comments
}

func (c *Client) Posts(userID int) []model.Post {
	posts, err := c.ListUserPosts(userID)
	if err != nil {
		log.Fatalf("Could not request posts for user with id %d: %s", userID, err)
	}
	return posts
}


func addQuery(r *http.Request, key, val string) {
	query := r.URL.Query()
	query.Add(key, val)
	r.URL.RawQuery = query.Encode()
}

func (c *Client) newRequest(method, path string) (*http.Request, error) {
	relPath := &url.URL{Path: path}
	url := c.BaseURL.ResolveReference(relPath).String()

	req, err := http.NewRequest(method, url, nil)
	log.Fatalf("error occured: %s", err)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
