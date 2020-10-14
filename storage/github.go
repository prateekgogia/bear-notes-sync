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
	baseRef, _, err := c.Git.GetRef(ctx, "githubName", "notes", "refs/heads/master")
	if err != nil {
		return err
	}

	newRef := &github.Reference{
		Ref:    github.String("refs/heads/master"),
		Object: &github.GitObject{SHA: baseRef.Object.SHA}}

	t := []github.TreeEntry{github.TreeEntry{
		Path: github.String(filename), Type: github.String("blob"),
		Content: github.String(content), Mode: github.String("100644")}}

	tree, _, err := c.Git.CreateTree(ctx, "githubName", "notes", *newRef.Object.SHA, t)
	if err != nil {
		return err
	}

	parent, _, err := c.Repositories.GetCommit(ctx, "githubName", "notes", *newRef.Object.SHA)
	if err != nil {
		return err
	}

	parent.Commit.SHA = parent.SHA
	commitMessage := fmt.Sprintf("test commit")
	curTime := time.Now()
	author := &github.CommitAuthor{
		Date:  &curTime,
		Name:  github.String("githubName"),
		Email: github.String("githubName@gmail.com"),
	}

	commitInfo := &github.Commit{Author: author, Message: github.String(commitMessage),
		Tree: tree, Parents: []github.Commit{*parent.Commit}}

	commit, _, err := c.Git.CreateCommit(ctx, "githubName", "notes", commitInfo)
	if err != nil {
		return err
	}

	newRef.Object.SHA = commit.SHA
	reference, _, err := c.Git.UpdateRef(ctx, "githubName", "notes", newRef, false)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully committed %+v\n", commit)
	fmt.Printf("Successfully committed reference is %+v\n", reference)
	return nil
}
