package model

import "strings"

// Post struct for post
type Post struct {
	ID     int
	UserID int
	Title  string
	Body   string
}

// Comment struct for comment
type Comment struct {
	ID     int
	PostID int
	Name   string
	Email  string
	Body   string
}

// Contains checks if body or email contains substring
func (c *Comment) Contains(substr string) bool {
	if strings.Contains(c.Body, substr) || strings.Contains(c.Email, substr) {
		return true
	}

	return false
}
