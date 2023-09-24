package cmd

import (
	"fmt"
	"time"

	"github.com/open-pomodoro/openpomodoro-cli/format"
	"github.com/open-pomodoro/openpomodoro-cli/hook"
	"github.com/spf13/cobra"
)

func init() {
	command := &cobra.Command{
		Use:   "finish",
		Short: "Finish the current Pomodoro",
		RunE:  finishCmd,
	}

	RootCmd.AddCommand(command)
}

func finishCmd(cmd *cobra.Command, args []string) error {
	p, err := client.Pomodoro()
	if err != nil {
		return err
	}

  now := time.Now()
	d := now.Sub(p.StartTime)
  runDuration := format.DurationAsTime(d)
	fmt.Println(runDuration)

  hookArgs := hook.ArgsFromPomodoro(p)
  hookArgs = append(hookArgs,
    fmt.Sprintf("--end-time='%s'", now.String()),
    fmt.Sprintf("--run-duration='%s'", runDuration),
  )

	if err := hook.Run(client, "finish", hookArgs); err != nil {
		return err
	}

	return client.Finish()
}
