package tasks

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// 执行 command
func RunExecute(name string, arg ...string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command(name, arg...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

/*
	go 创建项目

	1. 建目录
		mkdir -p $GOPATH/src/github.com/hans007/xxx

	2. git clone
		git clone https://github.com/hans007/xxx $GOPATH/src/github.com/hans007/xxx

	3. vscode 打开
		code $GOPATH/src/github.com/hans007/xxx

*/

// 建目录
func GoMkdir(owner string, repo string) {
	path := fmt.Sprintf("%s/src/github.com/%s/%s", os.Getenv("GOPATH"), owner, repo)
	fmt.Println(fmt.Sprintf("mkdir -p %s", path))
	RunExecute("mkdir", "-p", path)
}

// git clone
func GoGitClone(owner string, repo string) {
	url := fmt.Sprintf("https://github.com/%s/%s", owner, repo)
	path := fmt.Sprintf("%s/src/github.com/%s/%s", os.Getenv("GOPATH"), owner, repo)
	fmt.Println(fmt.Sprintf("git clone %s %s", url, path))
	RunExecute("git", "clone", url, path)
}

// vscode open
func GoVSCodeOpen(owner string, repo string) {
	path := fmt.Sprintf("%s/src/github.com/%s/%s", os.Getenv("GOPATH"), owner, repo)
	fmt.Println(fmt.Sprintf("code %s", path))
	RunExecute("code", path)
}
