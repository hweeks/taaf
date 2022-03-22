package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func RunACommand(cmd string, timeout_base int) ([]byte, error) {
	if timeout_base <= 0 {
		fmt.Println("honestly, a time of zero or less seems unintentional")
		panic(1)
	}
	time_to_wait := time.Duration(timeout_base) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), time_to_wait)
	defer cancel()
	root_cmd := exec.CommandContext(ctx, "bash", "-c", cmd)
	cmd_out, err := root_cmd.Output()
	if err != nil {
		fmt.Println("> command failure")
		fmt.Println(string(cmd))
		return cmd_out, err
	}
	return cmd_out, nil
}

func ValidateLocalEnv() {
	fmt.Println("checking for ansible...")
	_, err := RunACommand("ansible --version", 5)
	if err != nil {
		fmt.Println("whelp, i will install ansible for you")
		final_cmd, err := RunACommand("pip3 install ansible", 500)
		fmt.Println(string(final_cmd))
		panic(err)
	}
	fmt.Println("ansible works, next!")
}
