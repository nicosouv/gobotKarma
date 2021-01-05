package gitlab

import (
	Lib "app/cmd/gobot/lib"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
)

func GrabProjects() []*gitlab.Project {
	git := gitlab.NewClient(nil, Lib.Getenv("GITLAB_PRIVATE_TOKEN"))
	git.SetBaseURL(Lib.Getenv("GITLAB_URL"))
	projects, _, _ := git.Projects.ListProjects(nil)

	return projects
}


func GrabMRsForAllProjects()  {
	log.Printf("-- Begin")
	git := gitlab.NewClient(nil, Lib.Getenv("GITLAB_PRIVATE_TOKEN"))
	git.SetBaseURL(Lib.Getenv("GITLAB_URL"))

	var mrs []*gitlab.MergeRequest
	//var pidMrs []*gitlab.MergeRequest
	var str = new(string)
	*str = "opened"

	var mrOpts = gitlab.ListMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: 100,
		},
		State:       str,
	}
	log.Printf("-- Grab projects")
	projects := GrabProjects()

	log.Printf("-- for")
	for _, element := range projects {
		var pidMrs []*gitlab.MergeRequest
		pidMrs, _, _ = git.MergeRequests.ListMergeRequests(element.ID, &mrOpts)
		fmt.Print("%v\n", pidMrs)
		mrs = append(mrs, pidMrs...)
	}
	log.Printf("-- endfor")
	fmt.Print("%v\n", mrs)

	//return mrs
}