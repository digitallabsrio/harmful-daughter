package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"../../pkg/repo_manager"
)

func main() {
	fmt.Println("**** main() start")
	command := flag.String("command", "", "The git command")
	ignoreErros := flag.Bool(
		"ignore-errors",
		false,
		"Keep running after error if true")
	flag.Parse()

	// Get managed repos from environment variables
	root := os.Getenv("MG_ROOT")
	if root[len(root)-1] != '/' {
		root += "/"
	}

	repoNames := strings.Split(os.Getenv("MG_REPOS"), ",")

	repoManager, err := repo_manager.NewRepoManager(root, repoNames, *ignoreErros)
	if err != nil {
		log.Fatal(err)
	}

	output, _ := repoManager.Exec(*command)
	for repo, out := range output {
		fmt.Printf("[%s]: git %s\n", path.Base(repo), *command)
		fmt.Println(out)
	}

	fmt.Println("Done.")
}
