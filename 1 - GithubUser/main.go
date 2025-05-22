package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type UserGithub []struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarID   string `json:"gravatar_id"`
		URL          string `json:"url"`
		AvatarURL    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Ref          interface{} `json:"ref"`
		RefType      string      `json:"ref_type"`
		MasterBranch string      `json:"master_branch"`
		Description  string      `json:"description"`
		PusherType   string      `json:"pusher_type"`
	} `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	for _, username := range os.Args[1:] {
		req, err := http.Get("https://api.github.com/users/" + username + "/events")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data UserGithub
		err = json.Unmarshal(res, &data)

		if err != nil {
			println(err)
		}

		for k, v := range data {
			if k < 3 {
				fmt.Println("Últimos 3 repositorios do usuário: ", v.Repo.Name)
			}
		}
	}
}

// range os.Args[:1] -> Passar dados via CLI
// os.Stderr -> Para mostrar no terminal
