package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func exists(name string) bool {

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func exe(jar string) {

	args := append([]string{"-jar", jar}, os.Args[1:]...)

	cmd := exec.Command("java", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func main() {

	imagePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	jar := imagePath + ".jar"
	if exists(jar) {
		exe(jar)

	} else {
		ext := filepath.Ext(imagePath)
		if ext != "" {

			jar = imagePath[:len(imagePath)-len(ext)] + ".jar"
			if exists(jar) {
				exe(jar)
			}
		}
	}

	fmt.Printf("UNSUCCESS: Properly named executable jar file was not found in %s directory\n", filepath.Dir(imagePath))
	os.Exit(1)
}
