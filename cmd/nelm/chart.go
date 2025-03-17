package main

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/werf/common-go/pkg/cli"
)

func newChartCommand(ctx context.Context, afterAllCommandsBuiltFuncs map[*cobra.Command]func(cmd *cobra.Command) error) *cobra.Command {
	cmd := cli.NewGroupCommand(
		ctx,
		"chart",
		"Manage charts.",
		"Manage charts.",
		chartCmdGroup,
		cli.GroupCommandOptions{},
	)

	cmd.AddCommand(newChartRenderCommand(ctx, afterAllCommandsBuiltFuncs))
	cmd.AddCommand(newChartDependencyCommand(ctx, afterAllCommandsBuiltFuncs))
	cmd.AddCommand(newChartDownloadCommand(ctx, afterAllCommandsBuiltFuncs))
	cmd.AddCommand(newChartUploadCommand(ctx, afterAllCommandsBuiltFuncs))
	cmd.AddCommand(newChartPackCommand(ctx, afterAllCommandsBuiltFuncs))
	cmd.AddCommand(newChartLintCommand(ctx, afterAllCommandsBuiltFuncs))
	cmd.AddCommand(newChartSecretCommand(ctx, afterAllCommandsBuiltFuncs))

	return cmd
}
