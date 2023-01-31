package ingressgen

import (
	loggerUtil "github.com/6za/cluster-tools/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	hostDomain string
	apps       string
	repo       string
	path       string
	ip         string
	appList    []string
)

// NewCommand is the parent command of the folder
// The standard explored, this command is called `NewCommand` to help identify its role.
func NewCommand() *cobra.Command {

	ingressGenCmd := &cobra.Command{
		Use:     "ingress-gen",
		Short:   "generate ingress routes",
		Long:    "",
		PreRunE: validateIngressGen,
		RunE:    executeIngressGen,
		// PostRunE: runPostAction,
	}

	ingressGenCmd.Flags().StringVar(&hostDomain, "host-domain", "", "the suffix name to be used on etc hosts, for example: mylocal.cloud.internal")
	ingressGenCmd.Flags().StringVar(&apps, "apps", "minio,minio-console,vault,atlantis,chartmuseum,argo,kubefirst,argocd", "comma separated list fo apps names to be used on the generation")
	ingressGenCmd.Flags().StringVar(&repo, "repo", "", "gitops repo address")
	ingressGenCmd.Flags().StringVar(&path, "path", "registry", "path at the repository for the application to be installed")
	ingressGenCmd.Flags().StringVar(&ip, "ip", "127.0.0.1", "LAN IP of the local installation machine")

	return ingressGenCmd
}

func init() {
	log.Logger = loggerUtil.ZerologSetup(zerolog.InfoLevel)

}
