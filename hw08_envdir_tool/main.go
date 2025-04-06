package main

import (
	"fmt"
	"os"
)

func main() {
	mapEnv, err := ReadDir(os.Args[1])
	if err != nil {
		fmt.Println(fmt.Errorf("read env dir error: %w", err))
		return
	}

	code := RunCmd(os.Args[2:], mapEnv)
	os.Exit(code)
}
