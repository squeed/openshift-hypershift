package infra

import (
	"github.com/spf13/cobra"

	"github.com/openshift/hypershift/cmd/infra/aws"
	"github.com/openshift/hypershift/cmd/infra/azure"
)

func NewDestroyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "infra",
		Short:        "Commands for destroying HyperShift infra resources",
		SilenceUsage: true,
	}

	cmd.AddCommand(aws.NewDestroyCommand())
	cmd.AddCommand(azure.NewDestroyCommand())

	return cmd
}
