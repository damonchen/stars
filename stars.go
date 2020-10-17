package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/google/go-github/v32/github"
)

func quit(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func main() {
	var username string

	var rootCmd = &cobra.Command{
		Use:   "stars",
		Short: "stars is a command line tool for you to search your github star repos",
		Long:  `stars is a command line tool for you to search your github star repos`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			client := github.NewClient(nil)
			opts := &github.ActivityListStarredOptions{
				ListOptions: github.ListOptions{
					Page:    1,
					PerPage: 100,
				},
			}

			type Repo struct {
				Name     string
				FullName string
			}
			var repos []Repo

			page := 1
			for {
				opts.Page = page
				starRepos, resp, err := client.Activity.ListStarred(context.Background(), username, opts)
				if err != nil {
					quit(err)
				}

				for _, repo := range starRepos {
					repos = append(repos, Repo{
						Name:     *repo.Repository.Name,
						FullName: *repo.Repository.FullName,
					})

				}

				if page >= resp.LastPage {
					break
				}
				page++
			}

			for _, repo := range repos {
				fmt.Printf("%v https://github.com/%v\n", repo.Name, repo.FullName)
			}
		},
	}
	rootCmd.PersistentFlags().StringVar(&username, "name", "damonchen", "github username")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

