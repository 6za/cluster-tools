package tunnel

import (
	loggerUtil "github.com/6za/cluster-tools/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	owner string
	repo  string
)

// NewCommand is the parent command of the folder
// The standard explored, this command is called `NewCommand` to help identify its role.
func NewCommand() *cobra.Command {

	tunnelCheckCmd := &cobra.Command{
		Use:     "tunnel-check",
		Short:   "verify and update atlanits tunnel on github",
		Long:    "",
		PreRunE: validateTunnel,
		RunE:    executeTunnel,
		// PostRunE: runPostAction,
	}

	tunnelCheckCmd.Flags().StringVar(&owner, "owner", "", "repository that will observed fro changes on tunnels")
	tunnelCheckCmd.Flags().StringVar(&repo, "repo", "", "owner of repository that will observed fro changes on tunnels, organization or user")

	return tunnelCheckCmd
}

func init() {
	log.Logger = loggerUtil.ZerologSetup(zerolog.InfoLevel)

}
