package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v28/github"
	"github.com/hans007/goGithubCreate/tasks"
	"golang.org/x/oauth2"
)

var (
	name          = flag.String("name", "", "仓库名称")
	description   = flag.String("des", "", "说明")
	private       = flag.Bool("private", false, "是否私有项目")
	repoHans      = flag.Bool("hans", true, "hans 账户")
	repoDucafecat = flag.Bool("ducafecat", false, "ducafecat 账户")
	execGo        = flag.Bool("go", false, "执行 go")
	execVue       = flag.Bool("vue", false, "执行 vue")
)

func main() {
	flag.Parse()

	fmt.Println(*name)

	token := os.Getenv("GITHUB_AUTH_TOKEN_HANS")
	if *repoDucafecat {
		token = os.Getenv("GITHUB_AUTH_TOKEN_DUCAFECAT")
	}

	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	r := &github.Repository{Name: name, Private: private, Description: description}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())

	owner := "hans007"
	if *repoDucafecat {
		owner = "ducafecat"
	}

	if *execGo {
		tasks.GoMkdir(owner, *name)      // 目录
		tasks.GoGitClone(owner, *name)   // git clone
		tasks.GoVSCodeOpen(owner, *name) // vscode open
	}

}
