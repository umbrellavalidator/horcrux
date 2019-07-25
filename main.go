package main

import (
	"fmt"
	//	"os"
	"log"
	"os/exec"
)

func main() {
	p := "ls"
	cmd := exec.Command(p)

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}

	fmt.Println("\nThe End")

}
