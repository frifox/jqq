package main

import (
	"log/slog"
	"os"
	"os/exec"
)

func init() {
	_, err := exec.LookPath("jq")
	if err != nil {
		slog.Error("jq not found")
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]
	args = append(args, "-c", "--raw-input", "--raw-output", ". as $raw | try fromjson catch $raw")

	jq := exec.Command("jq", args...)

	jq.Stdin = os.Stdin
	jq.Stdout = os.Stdout
	jq.Stderr = os.Stderr

	err := jq.Start()
	if err != nil {
		slog.Error("jq start failed: %v", err)
		os.Exit(1)
	}

	jq.Wait()
}
