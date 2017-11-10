package searcher

import "github.com/giantswarm/versionbundle"

type Response struct {
	ReleaseVersion string                 `json:"release_version"`
	VersionBundles []versionbundle.Bundle `json:"version_bundles"`
}

func DefaultResponse() Response {
	return Response{
		ReleaseVersion: "",
		VersionBundles: []versionbundle.Bundle{},
	}
}
