package release

import (
	"fmt"
	"strings"

	log "github.com/saulshanabrook/pypi-dockerhub/Godeps/_workspace/src/github.com/Sirupsen/logrus"

	"github.com/saulshanabrook/pypi-dockerhub/Godeps/_workspace/src/github.com/google/go-github/github"
)

type Release struct {
	Name    string
	Version string
	Time    int64
}

func (r *Release) DockerfilePath() string {
	return r.Name + "/Dockerfile"
}

func (r *Release) DockerfileContents() string {
	return fmt.Sprintf("FROM python\nCMD [ \"pip\", \"install\", \"%v==%v\" ]\n", r.Name, r.Version)
}

func (r *Release) GithubTagName() string {
	return fmt.Sprintf("%v@%v", strings.ToLower(r.Name), r.Version)
}

func (r *Release) GithubTagMessage(rcr *github.RepositoryContentResponse) string {
	return fmt.Sprintf(
		"Version %v for %v\n\nFor commit %v",
		r.Version,
		r.Name,
		rcr.SHA,
	)
}

func (r *Release) GithubCommitMessage() string {
	return fmt.Sprintf(
		"Adding version %v for %v\n\nAdded %v seconds after epoch",
		r.Version,
		r.Name,
		r.Time,
	)
}

func (r *Release) DockerhubName() string {
	return strings.ToLower(r.Name)
}

func (r *Release) DockerhubTag() string {
	return r.Version
}

func (r *Release) DockerhubRepoShortDescription() string {
	return fmt.Sprintf("%v PyPi package based on python:3", r.Name)
}

func (r *Release) DockerhubRepoFullDescription() string {
	return fmt.Sprintf(`# %[1]v

[Autogenerated build](https://github.com/saulshanabrook/pypi-dockerhub) of [%[1]v PyPi package](https://pypi.python.org/pypi/%[1]v), based on the `+"`python:3`"+` image.

Using the `+"`latest`"+` branch will pull in the last version built.

I recommend pinning to an explicit version number, because it is possible if an old version is built retroactively that it will be set to `+"`latest`"+`.

Check [the tags page](./tags/) for a list of all build version.`, r.Name)
}

func (r *Release) Log() *log.Entry {
	return log.WithFields(log.Fields{
		"name":    r.Name,
		"version": r.Version,
		"time":    r.Time,
	})
}
