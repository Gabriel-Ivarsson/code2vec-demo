package releases

import (
	"github.com/spf13/cobra"

	"github.com/superfly/flyctl/internal/command/apps"
)

// TODO: deprecate
func New() *cobra.Command {
	return apps.NewReleases()
}
