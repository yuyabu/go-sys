package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	result, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(result))

	state := cmd.ProcessState
	fmt.Printf("%s \n", state.String())
	fmt.Printf("  pid: %d\n", state.Pid())
	fmt.Printf("  system: %v\n", state.SystemTime())
	fmt.Printf("  user: %v\n", state.UserTime())

}
