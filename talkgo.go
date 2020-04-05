package main

import (
	"log"
)

// BecomeAContributor become a contributor
func BecomeAContributor(username string) bool {
	log.Printf("Congratulations！@%s, you have been a contributor to night-reading-go.", username)
	return true
}
