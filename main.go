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

	b := git.BuscadorGit{
		DonoDoRepositorio: donoDoRepositorio,
		Repo:              repo,
	}
	b.BuscaGitTag(tag)

	fmt.Println("Sai do programa")
}
