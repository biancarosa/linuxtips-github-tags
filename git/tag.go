package git

import (
	"fmt"
	"net/http"
)

type Buscador interface {
	BuscaGitTag(tag, donoDoRepositorio, repo string)
}

type BuscadorGit struct {
	Repo              string
	DonoDoRepositorio string
}

func (b *BuscadorGit) BuscaGitTag(tag string) {
	fmt.Println("Buscando commits pertencentes a git tag...")
	fmt.Printf("%s %s %s\n", tag, b.DonoDoRepositorio, b.Repo)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", b.DonoDoRepositorio, b.Repo)
	// Go não tem exceção!!!
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("oi")
		fmt.Println(resp.Body)
	}
	// panic("ME SALVA")

}
