package constants

var Const = struct {
	CETypesEtag                    string
	CETypesFile                    string
	CETypesURL                     string
	DockerLocaldevServerImage      string
	DockerNetwork                  string
	ExtClientExtensionDir          string
	ExtClientExtensionDirSpecified string
	ReleasesEtag                   string
	ReleasesFile                   string
	ReleasesURL                    string
	RepoDir                        string
	RepoRemote                     string
	RepoBranch                     string
	RepoSync                       string
	TlsLfrdevDomain                string
	Verbose                        int8
}{
	CETypesEtag:                    "cetypes.etag",
	CETypesFile:                    "cetypes.file",
	CETypesURL:                     "cetypes.url",
	DockerLocaldevServerImage:      "docker.localdev.server.image",
	DockerNetwork:                  "docker.network",
	ExtClientExtensionDir:          "extension.client-extension.dir",
	ExtClientExtensionDirSpecified: "extension.client-extension.dir-specified",
	ReleasesEtag:                   "releases.etag",
	ReleasesFile:                   "releases.file",
	ReleasesURL:                    "releases.url",
	RepoDir:                        "localdev.resources.dir",
	RepoRemote:                     "localdev.resources.remote",
	RepoBranch:                     "localdev.resources.branch",
	RepoSync:                       "localdev.resources.sync",
	TlsLfrdevDomain:                "tls.lfrdev.domain",
	Verbose:                        1,
}
