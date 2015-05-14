package main

// Build information populated by the LDFLAGS on the compiler/linker.
var (
	buildDate       string
	buildCommitHash string
	buildBranch     string
	buildAppVersion string
	buildMachine    string
	buildGoVersion  string
	buildOS         string
	buildRelease    string
)

type BuildInfo struct {
	Date       string `json:"date"`
	CommitHash string `json:"commitHash"`
	Branch     string `json:"branch"`
	AppVersion string `json:"appVersion"`
	Machine    string `json:"machine"`
	GoVersion  string `json:"goVersion"`
	BuildOS    string `json:"buildOS"`
	Release    string `json:"release"` //dev, beta, RC1, stable
}

var buildInfo BuildInfo

func init() {
	buildInfo.Date = buildDate
	buildInfo.CommitHash = buildCommitHash
	buildInfo.Branch = buildBranch
	buildInfo.AppVersion = buildAppVersion
	buildInfo.Machine = buildMachine
	buildInfo.GoVersion = buildGoVersion
	buildInfo.BuildOS = buildOS
	buildInfo.Release = buildRelease
}
