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

	owner := "talkgo"
	repo := "night"

	// close old issues first
	closeOldIssues(client, owner, repo, issueInfo.Title)

	req := &github.IssueRequest{
		Title:     &issueInfo.Title,
		Labels:    &issueInfo.Labels,
		Assignees: &issueInfo.Assignee,
		Body:      &issueInfo.Body,
	}

	// create new issues
	issue, resp, err := client.Issues.Create(context.Background(), owner, repo, req)
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

	_ = resp.Body.Close()

	fmt.Printf("Created new issue %d %s \n", issue.GetNumber(), issue.GetTitle())
}

func closeOldIssues(client *github.Client, owner, repo string, name string) {

	issues, resp, err := client.Issues.List(context.Background(), false, &github.IssueListOptions{
		State: "open",
	})
	if err != nil {
		if resp != nil {
			data, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fmt.Println("github api get issues list response :", string(data))
			}
		}
		panic(err)
	}

	_ = resp.Body.Close()

	c := "closed"
	for _, issue := range issues {

		if issue.GetTitle() != name {
			continue
		}

		num := issue.GetNumber()

		//@all-contributors
		createPRByAllContributorsBot(client, owner, repo, num)

		_, _, err := client.Issues.Edit(context.Background(), owner, repo, num, &github.IssueRequest{
			State: &c,
		})

		if err != nil {
			panic(err)
		}

		fmt.Printf("close issue %d %s \n", num, issue.GetTitle())
	}
}

func createPRByAllContributorsBot(client *github.Client, owner, repo string, num int) {

	commentsList, resp, err := client.Issues.ListComments(context.Background(), owner, repo, num, &github.IssueListCommentsOptions{})
	if err != nil {
		if resp != nil {
			data, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fmt.Println("github api get issues list response :", string(data))
			}
		}
		panic(err)
	}

	_ = resp.Body.Close()

	var userNames string
	for _, v := range commentsList {
		userNames += "@" + *v.User.Login + " "
	}

	commentTempleate := "@all-contributors please add %s to doc."
	body := fmt.Sprintf(commentTempleate, userNames)
	_, resp, err = client.Issues.CreateComment(context.Background(), owner, repo, num, &github.IssueComment{
		Body: &body,
	})
	if err != nil {
		if resp != nil {
			data, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fmt.Println("github api get issues list response :", string(data))
			}
		}
		panic(err)
	}

	_ = resp.Body.Close()

	fmt.Printf("add users to contributors %s \n", userNames)
}
