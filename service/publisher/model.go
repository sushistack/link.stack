package publisher

import "github.com/sushistack/link.stack/service/git"

type PublisherType int

const (
	STATIC_WEB PublisherType = iota
	COMMENT
)

type ServiceProviderType int

const (
	PRIVATE_BLOG_NETWORK ServiceProviderType = iota
	CLOUD_BLOG_NETWORK
)

type Comment struct {
	id            string
	publisherType PublisherType
	postUrl       string
}

type StaticWebPage struct {
	Id            string
	PublisherType PublisherType
	Domain        string
	Provider      ServiceProviderType
	Repository    *git.GitRepository
}
