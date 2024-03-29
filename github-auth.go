package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/google/go-github/v28/github"
	"golang.org/x/crypto/ssh/terminal"
)

func main2() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("GitHub Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("GitHub Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())
	ctx := context.Background()
	user, _, err := client.Users.Get(ctx, "")

	// Is this a two-factor auth error? If so, prompt for OTP and try again.
	if _, ok := err.(*github.TwoFactorAuthError); ok {
		fmt.Print("\nGitHub OTP: ")
		otp, _ := r.ReadString('\n')
		tp.OTP = strings.TrimSpace(otp)
		user, _, err = client.Users.Get(ctx, "")
	}

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Printf("\n%v\n", github.Stringify(user))
}
