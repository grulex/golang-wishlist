package image

import (
	"github.com/grulex/go-wishlist/pkg/file"
	"time"
)

type ID string

type Image struct {
	ID        ID
	FileLink  file.Link
	Width     uint
	Height    uint
	Hash      Hash
	CreatedAt time.Time
}

type Hash struct {
	AHash string
	DHash string
	PHash string
}