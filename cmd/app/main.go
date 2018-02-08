package main

import (
	"net/url"
	"fmt"

	. "github.com/mafuyuk/tip-moya4"
	"github.com/mafuyuk/tip-moya4/bot"

	"github.com/ChimeraCoder/anaconda"
)

func main() {

	t := NewTwitter(
		"yDyBqCYT9UFbsqPjNvDuBWgTU",
		"bUugxUzPY7CZyWnyECYH2QqVfK1cStWUBck0deXGjqvdJvei5c",
		"787995333573775360-70dZWSrFukC01kco79KzltU4WhoN0HS",
		"DexvphD9VSeZmzxvbGvmMu71slXn32WSenJbXUfYPjt4q",
	)

	//text := "test twitter API: Hello from Golang."
	//tweet, err := t.Post(text)
	//if err != nil {
	//	log.Fatal(err)
	//}

	v := url.Values{}
	stream := t.UserStream(v) // 接続

	for {
		// 受信待ち
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				// Tweet を受信
				fmt.Printf("%s: %s: %s\n", status.User.Id, status.User.ScreenName, status.Text)

				bitClient, err := bot.New(status.User.Id, status.User.ScreenName, status.Text)
				if err != nil {
					fmt.Printf("skip: %v", err)
					return
				}
        bot.Exec(bitClient)
			case anaconda.StatusDeletionNotice:
			// pass
			default:
			}
		}
	}
}
