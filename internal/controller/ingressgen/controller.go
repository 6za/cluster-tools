package ingressgen

import "github.com/6za/cluster-tools/internal/generator"

func GenerateArtifacts(hostSuffix string, apps []string, repo string, path string) error {

	generator.GenerateIngressesFromApps(hostSuffix, apps)
	return nil
}
