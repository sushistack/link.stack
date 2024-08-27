package service

import "github.com/sushistack/link.stack/service/git"

type OrderType int

const (
	STANDARD OrderType = iota
	DELUXE
	PRIMIUM
)

type OrderStatus int

const (
	READY OrderStatus = iota
	STAGED
	PROCESSING
	DONE
)

type Article struct {
	Title       string
	Description string
	Content     string
}

type Post struct {
	Id          string
	FilePath    string
	FileName    string
	Extension   string
	PublisherId string
}

type Order struct {
	Id        string
	OrderType OrderType
	Url       string
	Status    OrderStatus
}

type LinkNode struct {
	Id           string
	Tier         int
	Url          string
	Repository   git.GitRepository
	Post         *Post
	Order        *Order
	ParentNodeId string
}

type ExecelSheet[T any] struct {
	Name    string
	Columns []string
	Data    []T
}
