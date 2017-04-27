package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/urbint/drone-sentry/sentry"
)

type Args struct {
	ApiKey    string `envconfig:"sentry_auth_token"`
	Org       string `envconfig:"sentry_organization"`
	Project   string `envconfig:"sentry_project"`
	Environ   string `envconfig:"sentry_release_environment"`
	Version   string `envconfig:"sentry_release_version"`
}

type DroneVars struct {
	BuildNumber   int    `envconfig:"build_number"`
	BuildFinished string `envconfig:"build_finished"`
	BuildStatus   string `envconfig:"build_status"`
	BuildLink     string `envconfig:"build_link"`
	CommitSha     string `envconfig:"commit_sha"`
	CommitBranch  string `envconfig:"commit_branch"`
	CommitAuthor  string `envconfig:"commit_author"`
	CommitLink    string `envconfig:"commit_link"`
	CommitMessage string `envconfig:"commit_message"`
	JobStarted    int64  `envconfig:"job_started"`
	Repo          string `envconfig:"build_link"`
	RepoLink      string `envconfig:"repo_link"`
	System        string
}

func main() {
	var (
		err   error
		vargs Args
		drone DroneVars
	)

	err = envconfig.Process("plugin", &vargs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = envconfig.Process("sentry", &drone)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create the Sentry client
	client := sentry.NewClient(vargs.ApiKey, vargs.Org, vargs.Project)

	// generate the Sentry objects
	release := sentry.Release{
		Version: vargs.Version,
		Ref:  "vargs.Version",
	}

	deploy := sentry.Deploy{
		Name: vargs.Version,
		Environment: vargs.Environ,
	}

	// sends the message
	if err := client.CreateRelease(&release); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err2 := client.CreateDeploy(&deploy); err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
