package app

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Repo struct {
	Owner string
	Name  string
}

type Issue struct {
	Number int
	Repo   *Repo
}

type GitHub struct {
	Client *githubv4.Client
}

func NewGitHub(token string) *GitHub {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	return &GitHub{Client: client}
}

func (g *GitHub) UpdateIssue(issue *Issue, title string, body string) error {
	ctx := context.Background()

	variables := map[string]interface{}{
		"name":   githubv4.String(issue.Repo.Name),
		"owner":  githubv4.String(issue.Repo.Owner),
		"number": githubv4.Int(issue.Number),
	}
	var query struct {
		Repository struct {
			Issue struct {
				ID githubv4.ID
			} `graphql:"issue(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	err := g.Client.Query(ctx, &query, variables)
	if err != nil {
		return err
	}

	var mutation struct {
		UpdateIssue struct {
			Issue struct {
				ID githubv4.ID
			}
		} `graphql:"updateIssue(input: $input)"`
	}
	input := githubv4.UpdateIssueInput{
		ID:    query.Repository.Issue.ID,
		Title: githubv4.NewString(githubv4.String(title)),
		Body:  githubv4.NewString(githubv4.String(body)),
	}
	err = g.Client.Mutate(ctx, &mutation, input, nil)
	if err != nil {
		return err
	}
	return nil
}
