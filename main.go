package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func scoop(out bool, args ...string) string {
	if out {
		cmd := exec.Command("scoop", args...)

		output, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}

		return strings.TrimSpace(string(output))
	}

	cmd := exec.Command("scoop", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	return ""
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <packages>\n", os.Args[0])
		os.Exit(0)
	}

	for i, a := range os.Args {
		if i == 0 {
			continue
		}

		deps := strings.Split(scoop(true, "depends", a), "\n")

		for _, d := range deps {
			if strings.TrimSpace(d) == "" {
				continue
			}

			scoop(false, "uninstall", "-p", d)
		}

		scoop(false, "uninstall", "-p", a)
	}
}
