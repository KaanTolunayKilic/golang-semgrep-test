package api

import (
	"github.com/KaanTolunayKilic/golang-semgrep-test/pkg/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListUserPosts(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters

		if req.URL.String() == "/posts" {
			t.Errorf("Request failed expected /posts got: %s", req.URL.String())
		}

		if val, ok := req.URL.Query()["userId"]; !ok || val[0] != "1" {
			t.Errorf("Request query not as expected: userId=1 == %s", req.URL.RawPath)
		}

		// Send response to be tested
		rw.Write([]byte(`[{
			"userId": 1,
			"id": 1,
			"title": "Mr. Robot",
			"body": "Lorem Ipsum Dolor"
		  }]`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	url, _ := url.Parse(server.URL)
	client := Client{BaseURL: url, httpClient: server.Client()}
	posts, err := client.ListUserPosts(1)

	if err != nil {
		t.Errorf("Failed ListUserPosts: %s", err)
	}

	if len(posts) != 1 {
		t.Errorf("Expected only one post: len(posts) == %d", len(posts))
	}

	expect := model.Post{
		ID:     1,
		UserID: 1,
		Title:  "Mr. Robot",
		Body:   "Lorem Ipsum Dolor",
	}
	actual := posts[0]

	if actual.Body != expect.Body ||
		actual.Title != expect.Title {
		t.Errorf("Expect: %v\nActual: %v\n", expect, actual)
	}
}
