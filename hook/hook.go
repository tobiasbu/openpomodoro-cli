package hook

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/open-pomodoro/go-openpomodoro"
)

const hookFileName = "hooks"

// Run runs a hook with the given name.
func Run(client *openpomodoro.Client, cmdName string, args []string) error {
	filename := path.Join(client.Directory, "hooks", hookFileName)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}

  cmdArgs := []string{cmdName}
  cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command(filename, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Hook %q failed:\n\n", cmdName)
		return err
	}

	return nil
}

func ArgsFromPomodoro(p *openpomodoro.Pomodoro) []string {
  if p == nil {
    return []string{}
  }
   hookArgs := []string{
    fmt.Sprintf("--duration=%s", p.Duration),
    fmt.Sprintf("--start-time='%s'", p.StartTime),
  }

  if strings.Trim(p.Description, " ") != "" {
    hookArgs = append(hookArgs, fmt.Sprintf("--description='%s'", p.Description))
  }

  if len(p.Tags) != 0 {
    hookArgs = append(hookArgs, fmt.Sprintf("--tags=%s", strings.Join(p.Tags, ",")))
  }

  return hookArgs
}
