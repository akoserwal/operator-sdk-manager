package search

import (
	"context"
	"fmt"
	"github.com/google/go-github/v31/github"
	"github.com/spf13/cobra"
	"strings"
)

//https://api.github.com/repos/operator-framework/operator-sdk/releases
//https://api.github.com/repos/operator-framework/operator-sdk/tags

const (
	GITHUB_API_URL= "https://api.github.com/repos/operator-framework/operator-sdk/tags"
)

var (
	listTags string
)

func NewCmd() *cobra.Command {
	searchCMD := &cobra.Command{
		Use:   "search",
		Short: "search available version",
		Long:  "",
		RunE:SearchVersions,
	}
	return searchCMD
}

func ListAllVersions() error{
	client := github.NewClient(nil)
	opt := &github.ListOptions{Page:1}
	repositoryTagList, response , err := client.Repositories.ListTags(context.Background(), "operator-framework", "operator-sdk", opt)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode == 200 {
		for _, tag := range repositoryTagList {
			fmt.Println(*tag.Name)
		}
	}

	return nil
}

func SearchVersions(cmd *cobra.Command, args []string) error{
	if len(args) > 0 {
		version := strings.ToLower(args[0])
		repos, response := GetReleaseByTag(version)
		if response.StatusCode == 200 {
			fmt.Println(*repos.Name)
			fmt.Println(*repos.Body)
		} else {
			fmt.Println("Version is not available")
		}
	} else {
		fmt.Println("provide version to search")
	}
	return nil
}

func GetReleaseByTag(version string) (*github.RepositoryRelease, *github.Response) {
	client := github.NewClient(nil)
	repos, response, err := client.Repositories.GetReleaseByTag(context.Background(), "operator-framework", "operator-sdk", version)
	if err != nil {
		fmt.Println(version, "version not found")
	}
	return repos, response
}