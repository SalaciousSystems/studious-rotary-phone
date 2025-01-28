package app

import "fmt"

var (
	buildVersion = "v0.0.1-nightly.1"
	buildBranch  = "stable" // [ stable | beta | alpha | nightly ]
)

func GetBuildMeta() string {
	return fmt.Sprintf("%s: %s", buildBranch, buildVersion)
}
