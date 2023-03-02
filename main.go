package main

import (
	"fmt"

	"github.com/biancarosa/linuxtips-github-tags/git"
)

func main() {
	fmt.Println("#VAAAAAAAAI")
	tag := "v1.26.2"
	donoDoRepositorio := "kubernetes"
	repo := "kubernetes"

	git.BuscaGitTag(tag, donoDoRepositorio, repo)

	go git.BuscaGitTag(tag, donoDoRepositorio, repo)
	fmt.Println("Sai do programa")
}
