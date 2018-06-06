package ghclient

import (
	"net/http"

	"github.com/google/go-github/github"
)

// Client is a github client used to get info from github.
type Client struct {
	owner string
	repo  string

	c *github.Client
}

// New creates a new client.
func New(tc *http.Client, owner, repo string) *Client {
	return &Client{
		owner: owner,
		repo:  repo,
		c:     github.NewClient(tc),
	}
}

// GetMergedPRsForMilestone returns a list of github issues that are merged PRs
// for this milestone.
func (c *Client) GetMergedPRsForMilestone(milestone string) []*github.Issue {
	return c.getMergedPRsForMilestone(milestone)
}

// GetMergedPRsForLabels returns a list of github issues that are merged PRs
// with the given label.
func (c *Client) GetMergedPRsForLabels(labels []string) []*github.Issue {
	return c.getMergedPRsForLabels(labels)
}

// GetOrgMembers returns a set of names of members in the org.
func (c *Client) GetOrgMembers(org string) map[string]struct{} {
	return c.getOrgMembers(org)
}

// CommitIDForMergedPR returns the commit id for pr.
//
// It returns "" if pr is not a merged PR.
func (c *Client) CommitIDForMergedPR(pr *github.Issue) string {
	return c.commitIDForMergedPR(pr)
}