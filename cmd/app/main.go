package main

import (
"fmt"
"github.com/ChimeraCoder/anaconda"
"log"
)

func main() {
	anaconda.SetConsumerKey("yDyBqCYT9UFbsqPjNvDuBWgTU")
	anaconda.SetConsumerSecret("bUugxUzPY7CZyWnyECYH2QqVfK1cStWUBck0deXGjqvdJvei5c")
	api := anaconda.NewTwitterApi("787995333573775360-70dZWSrFukC01kco79KzltU4WhoN0HS", "DexvphD9VSeZmzxvbGvmMu71slXn32WSenJbXUfYPjt4q")

	text := "test twitter API: Hello from Golang."
	tweet, err := api.PostTweet(text, nil)
	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println(tweet.Text)
}
