package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dyweb/gommon/util/httputil"
	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
)

type Issue struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Labels   []string `json:"labels"`
	Assignee []string `json:"assignee"`
}

func main() {

	var cfg string

	flag.StringVar(&cfg, "c", "", "")
	flag.Parse()

	data, err := ioutil.ReadFile(cfg)
	if err != nil {
		panic(err)
	}

	issueInfo := Issue{}
	err = json.Unmarshal(data, &issueInfo)
	if err != nil {
		panic(err)
	}

	var token = os.Getenv("GITHUB_TOKEN")

	hc := httputil.NewUnPooledClient()
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: token,
		})
		hc = oauth2.NewClient(context.Background(), ts)
	}

	client := github.NewClient(hc)

	req := &github.IssueRequest{
		Title:     &issueInfo.Title,
		Labels:    &issueInfo.Labels,
		Assignees: &issueInfo.Assignee,
		Body:      &issueInfo.Body,
	}

	issue, resp, err := client.Issues.Create(context.Background(), "talkgo", "night", req)
	if err != nil {

		// err != nil
		// maybe have response useful for debug
		if resp != nil {
			data, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fmt.Println("github api response :", string(data))
			}
		}

		panic(err)
	}

	fmt.Println("Created new issue %d %s", issue.GetNumber(), issue.GetTitle())

}
