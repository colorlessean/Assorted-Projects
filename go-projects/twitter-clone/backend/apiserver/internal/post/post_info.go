package post

import (
	"time"

	"github.com/google/uuid"
)

// PostID is a type definition for a post's unique identifier
type PostID uuid.UUID

// NewPostID generates a new post ID
func NewPostID() PostID {
	return PostID(uuid.New())
}

// PostInfo is the basic type for a post
type PostInfo struct {
	ID      PostID    `json:"id"`
	Author  string    `json:"author"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
}

// New provides a way to create a new post from an author username and content string
func New(author, content string) (*PostInfo, error) {
	return &PostInfo{
		ID:      NewPostID(),
		Author:  author,
		Date:    time.Now(),
		Content: content,
	}, nil
}
