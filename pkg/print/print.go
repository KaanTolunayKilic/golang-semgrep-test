package print

import (
	"fmt"
	"github.com/KaanTolunayKilic/golang-semgrep-test/pkg/model"
	"strings"
)

// Post prints post on console
func Post(post model.Post) {
	fmt.Printf("Post: %s (userId=%d)\n", post.Title, post.UserID)
	fmt.Println(" >", removeLineBreaks(post.Body))
}

// Comment prints comment on console
func Comment(comment model.Comment) {
	fmt.Printf("    Comment: %s (by %s)\n", removeLineBreaks(comment.Body), comment.Email)
}

// Divider prints dividing line on console
func Divider() {
	fmt.Println("-----------------------------------------------------------------------")
}

func removeLineBreaks(text string) string {
	return strings.Replace(text, "\n", "", -1)
}
