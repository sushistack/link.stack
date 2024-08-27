package git

type HostingService int

const (
	Github HostingService = iota
	GitLab
	Bitbucket
)

type GitAccount struct {
	Id             string
	Username       string
	HostingService HostingService
}

type GitRepository struct {
	Id    string
	Name  string
	Onwer *GitAccount
}
