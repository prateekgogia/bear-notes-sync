package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v26/github"
	"golang.org/x/oauth2"
)

type client struct {
	*github.Client
}

func NewClient() *client {

	oauth2Token := oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")}
	githubClient := github.NewClient(oauth2.NewClient(
		context.Background(), oauth2.StaticTokenSource(&oauth2Token)))
	return &client{Client: githubClient}
}

func (c *client) SaveFile(filename, content string) error {

	ctx := context.Background()
	baseRef, _, err := c.Git.GetRef(
		ctx,
		"prateekgogia",
		"notes",
		"refs/heads/master",
	)
	if err != nil {
		return err
	}

	newRef := &github.Reference{
		Ref:    github.String("refs/heads/master"),
		Object: &github.GitObject{SHA: baseRef.Object.SHA}}

	t := []github.TreeEntry{github.TreeEntry{
		Path: github.String(filename), Type: github.String("blob"),
		Content: github.String(content), Mode: github.String("100644")}}

	tree, _, err := c.Git.CreateTree(ctx, "prateekgogia", "notes", *newRef.Object.SHA, t)
	if err != nil {
		return err
	}

	parent, _, err := c.Repositories.GetCommit(ctx, "prateekgogia", "notes", newRef.Object.GetSHA())

	parent.Commit.SHA = parent.SHA
	commitMessage := fmt.Sprintf("test commit")
	time := time.Now()
	author := &github.CommitAuthor{
		Date:  &time,
		Name:  github.String("prateekgogia"),
		Email: github.String("prateekgogia42@gmail.com"),
	}

	commitInfo := &github.Commit{Author: author, Message: github.String(commitMessage),
		Tree: tree, Parents: []github.Commit{*parent.Commit}}

	commit, _, err := c.Git.CreateCommit(ctx, "prateekgogia", "notes", commitInfo)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully committed %+v\n", commit)
	return nil
}
