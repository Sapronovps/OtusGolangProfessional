package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		envName      string
		env          EnvValue
		expectedCode int
	}{
		{
			envName: "EMPTY",
			env: EnvValue{
				Value:      "",
				NeedRemove: false,
			},
			expectedCode: 0,
		},
		{
			envName: "UNSET",
			env: EnvValue{
				Value:      "",
				NeedRemove: true,
			},
			expectedCode: 0,
		},
		{
			envName: "HELLO",
			env: EnvValue{
				Value:      "\"hello\"",
				NeedRemove: false,
			},
			expectedCode: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.envName, func(t *testing.T) {
			mapEnv := make(map[string]EnvValue)
			mapEnv[tc.envName] = tc.env

			code := RunCmd([]string{"echo", "test test"}, mapEnv)
			env := os.Getenv(tc.envName)
			require.Equal(t, tc.env.Value, env)
			require.Equal(t, tc.expectedCode, code)
		})
	}
}

func TestRunEmptyCmd(t *testing.T) {
	t.Run("empty cmd", func(t *testing.T) {
		mapEnv := make(map[string]EnvValue)
		code := RunCmd([]string{}, mapEnv)
		require.Equal(t, 1, code)
	})
}

func TestRunNotExistsCmd(t *testing.T) {
	t.Run("not exists cmd", func(t *testing.T) {
		mapEnv := make(map[string]EnvValue)
		code := RunCmd([]string{"echo22", "test test"}, mapEnv)
		require.Equal(t, 1, code)
	})
}
