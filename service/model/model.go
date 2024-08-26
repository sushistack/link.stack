package model

type HostingService int

const (
	Github HostingService = iota
	GitLab
	Bitbucket
)

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

type Article struct {
	title       string
	description string
	content     string
}

type Post struct {
	id          string
	filePath    string
	fileName    string
	extension   string
	publisherId string
}

type GitAccount struct {
	id             string
	username       string
	hostingService HostingService
}

type GitRepository struct {
	id    string
	name  string
	onwer *GitAccount
}

type Order struct {
	id        string
	orderType OrderType
	url       string
	status    OrderStatus
}

type LinkNode struct {
	id           string
	tier         int
	url          string
	repository   GitRepository
	post         *Post
	order        *Order
	parentNodeId string
}

type Comment struct {
	id            string
	publisherType PublisherType
	postUrl       string
}

type StaticWebPage struct {
	id            string
	publisherType PublisherType
	domain        string
	provider      ServiceProviderType
	repository    *GitRepository
}
