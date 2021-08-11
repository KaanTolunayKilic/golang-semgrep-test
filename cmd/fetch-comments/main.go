package main

import (
	"flag"
	"fmt"
	"github.com/KaanTolunayKilic/golang-semgrep-test/pkg/api"
	print2 "github.com/KaanTolunayKilic/golang-semgrep-test/pkg/print"
	"os"
)

// Fetches all posts from user and all comments
func main() {
	userIDPtr := flag.Int("userId", -1, "User id *required")
	filterPtr := flag.String("filter", "", "Filter comments")
	flag.Parse()

	if *userIDPtr == -1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := api.NewClient(nil)
	fmt.Printf("Posts and comments for user with id %d:\n", *userIDPtr)
	for _, post := range client.Posts(*userIDPtr) {
		print2.Post(post)
		for _, comment := range client.Comments(&post) {
			if comment.Contains(*filterPtr) {
				print2.Comment(comment)
			}
		}
		print2.Divider()
	}
}