package tunnel

import (
	"fmt"

	"github.com/6za/cluster-tools/internal/controller/tunnel"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const IN_CLUSTER_ATLANTIS_ADDRESS = "http://ngrok-agent.ngrok-agent.svc.cluster.local:4040/api/tunnels"

func executeTunnel(cmd *cobra.Command, args []string) error {

	return tunnel.ExecuteTunnel(owner, repo)
}

func validateTunnel(cmd *cobra.Command, args []string) error {
	if len(repo) < 1 || len(owner) < 1 {
		return fmt.Errorf("both repo(%s) and owner(%s) must be provided in order for webhookupdater to work as expected", repo, owner)
	}
	log.Info().Msgf("Validation: Success repo(%s) and owner(%s) provided as epxected", repo, owner)
	return nil
}
