package cmd

import (
	"github.com/open-pomodoro/openpomodoro-cli/hook"
	"github.com/spf13/cobra"
)

func init() {
	command := &cobra.Command{
		Use:   "cancel",
		Short: "Cancel the current Pomodoro",
		RunE:  cancelCmd,
	}

	RootCmd.AddCommand(command)
}

func cancelCmd(cmd *cobra.Command, args []string) error {
  p, _ := client.Pomodoro()

	if err := hook.Run(client, "cancel", hook.ArgsFromPomodoro(p)); err != nil {
		return err
	}

	return client.Cancel()
}
