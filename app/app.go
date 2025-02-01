package app

func Run() error {
	token, repo, err := GetEnv()
	if err != nil {
		return err
	}

	flags, err := NewFlag()
	if err != nil {
		return err
	}

	err = IsValidUrl(flags.URL)
	if err != nil {
		return err
	}

	title, err := FetchTitle(flags.URL)
	if err != nil {
		return err
	}

	issue := &Issue{
		Number: flags.Number,
		Repo: &Repo{
			Owner: repo.Owner,
			Name:  repo.Name,
		},
	}
	gh := NewGitHub(token)
	if err := gh.UpdateIssue(issue, title, flags.URL); err != nil {
		return err
	}

	return nil
}
