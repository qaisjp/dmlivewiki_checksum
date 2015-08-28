package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"path"
)

func createFile(filename string) *os.File {
	fmt.Println("Creating", filename+"...")
	f, err := os.Create(filename)
	if err != nil {
		panic(f)
	}
	return f
}

func removeFile(filename string) {
	fmt.Printf("Removing %s...", filename)

	err := os.Remove(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(" does not exist!")
			return
		}
		panic(err)
	}
	fmt.Println(" success!")
}

func isDirectory(filepath string) (bool, os.FileInfo, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, nil, err
	}
	return fileInfo.IsDir(), fileInfo, err
}

func getLastPathComponents(filepath string, depth int) (absPath string) {
	for i := 1; i < depth; i++ {
		absPath = path.Base(filepath) + "\\" + absPath
		filepath = path.Dir(filepath)
	}
	return
}

func checkCommandArgumentNumber(c *cli.Context, n int) bool {
	if len(c.Args()) != n {
		cli.ShowCommandHelp(c, c.Command.Name)
		return false
	}
	return true
}
