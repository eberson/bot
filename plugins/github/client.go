package github

import (
	"context"

	"github.com/google/go-github/v31/github"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type Github struct {
	client *github.Client
}

func NewGithub(url, token string) (*Github, error) {
	c, err := newClient(url, token)
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Info("creating a new github client...")

	return &Github{
		client: c,
	}, nil
}

func newClient(url, token string) (*github.Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	tc := oauth2.NewClient(ctx, ts)

	cli, err := github.NewEnterpriseClient(url, url, tc)
	if err != nil {
		return nil, errors.Wrap(err, "attempt to connect to github failed")
	}
	return cli, nil
}

func (g *Github) ListPullRequestsByState(ctx context.Context, owner, repo, state string) ([]*github.PullRequest, error) {
	prs, _, err := g.client.PullRequests.List(
		ctx,
		owner,
		repo,
		&github.PullRequestListOptions{
			State: state,
		},
	)

	if err != nil {
		return nil, err
	}

	return prs, nil
}
