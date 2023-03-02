package git

import "fmt"

type Buscador interface {
	BuscaGitTag(tag, donoDoRepositorio, repo string)
}

type BuscadorGit struct {
}

func (b *BuscadorGit) BuscaGitTag(tag, donoDoRepositorio, repo string) {
	fmt.Println("Buscando commits pertencentes a git tag...")
	fmt.Printf("%s %s %s\n", tag, donoDoRepositorio, repo)
}
