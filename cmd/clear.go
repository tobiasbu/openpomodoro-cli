package cmd

import (
	"github.com/open-pomodoro/openpomodoro-cli/hook"
	"github.com/spf13/cobra"
)

func init() {
	command := &cobra.Command{
		Use:   "clear",
		Short: "Clear the current Pomodoro",
		RunE:  clearCmd,
	}

	RootCmd.AddCommand(command)
}

func clearCmd(cmd *cobra.Command, args []string) error {
  p, _ := client.Pomodoro()

	if err := hook.Run(client, "clear", hook.ArgsFromPomodoro(p)); err != nil {
		return err
	}

	return client.Clear()
}
