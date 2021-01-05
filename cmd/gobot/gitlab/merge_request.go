package gitlab

import (
	Lib "app/cmd/gobot/lib"
	"github.com/xanzy/go-gitlab"
)

func GrabProjects() []*gitlab.Project {
	git := gitlab.NewClient(nil, Lib.Getenv("GITLAB_PRIVATE_TOKEN"))
	git.SetBaseURL(Lib.Getenv("GITLAB_URL"))
	projects, _, _ := git.Projects.ListProjects(nil)

	return projects
}


func GrabMRsForAllProjects() []*gitlab.MergeRequest {
	git := gitlab.NewClient(nil, Lib.Getenv("GITLAB_PRIVATE_TOKEN"))
	git.SetBaseURL(Lib.Getenv("GITLAB_URL"))

	var mrs []*gitlab.MergeRequest
	var str = new(string)
	*str = "opened"

	var mrOpts = gitlab.ListMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: 100,
		},
		State:       str,
	}

	projects := GrabProjects()

	for _, element := range projects {
		var pidMrs []*gitlab.MergeRequest
		pidMrs, _, _ = git.MergeRequests.ListMergeRequests(element.ID, &mrOpts)
		mrs = append(mrs, pidMrs...)
	}

	return mrs
}