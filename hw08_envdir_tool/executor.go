package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		fmt.Println(fmt.Errorf("error: empty command"))
		return 1
	}

	if err := prepareEnvironment(env); err != nil {
		fmt.Println(fmt.Errorf("environment preparation error: %w", err))
		return 1
	}

	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Env = os.Environ()

	if err := command.Run(); err != nil {
		fmt.Println(fmt.Errorf("command execution error: %w", err))
		return 1
	}

	return 0
}

// Подготовка переменных окружения.
func prepareEnvironment(env Environment) error {
	for name, params := range env {
		if params.NeedRemove {
			if err := os.Unsetenv(name); err != nil {
				return fmt.Errorf("failed to unset %q: %w", name, err)
			}
			continue
		}

		if err := os.Setenv(name, params.Value); err != nil {
			return fmt.Errorf("failed to set %q: %w", name, err)
		}
	}
	return nil
}
