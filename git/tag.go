package git

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Buscador interface {
	BuscaGitTag(tag, donoDoRepositorio, repo string)
}

type BuscadorGit struct {
	Repo              string
	DonoDoRepositorio string
}

type JSONDoGithub struct {
	FullName    string   `json:"full_name"`
	HTMLUrl     string   `json:"html_url"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
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

		// nao facam isso em casaaaaa! nao ignorem errors!
		// nao faca em producaaaao!!!!
		b, _ := io.ReadAll(resp.Body)
		// fmt.Println(string(b))

		j := new(JSONDoGithub)
		err := json.Unmarshal(b, j)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(j.Description)
		fmt.Println(j.FullName)
		fmt.Println(j.HTMLUrl)
		fmt.Println(j.Topics)
	}

}
