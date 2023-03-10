package ingressgen

import (
	"fmt"
	"os"
	"strings"

	"github.com/6za/cluster-tools/internal/controller/ingressgen"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func executeIngressGen(cmd *cobra.Command, args []string) error {
	ingressgen.GenerateArtifacts(hostDomain, appList, repo, path)
	ingressgen.GenerateHostsLine(hostDomain, appList, ip)
	return nil
}

func validateIngressGen(cmd *cobra.Command, args []string) error {
	hasError := false
	if hostDomain == "" {
		log.Error().Msgf("Missing flag host-domain:  %s ", hostDomain)
		hasError = true

	}
	if apps == "" {
		log.Error().Msgf("Missing list of on flag apps:  %s ", apps)
		hasError = true

	} else {
		appList = strings.Split(apps, ",")
		log.Info().Msgf("List of apps:  %s ", appList)
	}
	if repo == "" {
		log.Error().Msgf("Missing flag repo:  %s ", repo)
		hasError = true

	}
	if repo == "" {
		log.Error().Msgf("Missing flag repo:  %s ", repo)
		hasError = true

	}
	if !strings.HasPrefix(repo, "https://") {
		log.Error().Msgf("Not supported repo:  %s ", repo)
		hasError = true
	}

	if !strings.HasSuffix(repo, ".git") {
		log.Error().Msgf("Not supported repo:  %s ", repo)
		hasError = true
	}
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Error().Msgf("Missing env variable GITHUB_AUTH_TOKEN.")
		hasError = true

	}
	if hasError {
		return fmt.Errorf("missing flags")
	}

	return nil
}
