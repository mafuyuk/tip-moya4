package main

import (
	"fmt"
	"log"

	. "github.com/mafuyuk/tip-moya4"
)

func main() {

	t := NewTwitter(
		"yDyBqCYT9UFbsqPjNvDuBWgTU",
		"bUugxUzPY7CZyWnyECYH2QqVfK1cStWUBck0deXGjqvdJvei5c",
		"787995333573775360-70dZWSrFukC01kco79KzltU4WhoN0HS",
		"DexvphD9VSeZmzxvbGvmMu71slXn32WSenJbXUfYPjt4q",
	)

	text := "test twitter API: Hello from Golang."
	tweet, err := t.Post(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tweet.Text)
}
