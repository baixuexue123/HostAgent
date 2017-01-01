package main

import (
	"fmt"
	"os"
	"os/exec"
)

func dmidecode() string {
	cmd := exec.Command("sudo", "dmidecode")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s", err)
		return ""
	}
	return fmt.Sprintf("%s", buf)
}

func main() {
	// cmd := exec.Command("sudo", "dmidecode")
	// buf, err := cmd.Output()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "The command failed to perform: %s", err)
	// 	return
	// }
	// fmt.Fprintf(os.Stdout, "Result: %s", buf)
	// fmt.Println("")
	fmt.Println(dmidecode())
}
