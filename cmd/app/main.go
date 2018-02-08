package main

import (

	. "github.com/mafuyuk/tip-moya4"
	"net/url"
	"github.com/ChimeraCoder/anaconda"
	"fmt"
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

	//values := url.Values{}
	//values.Add([クエリのキー], [値]
	//t.AccountActivity(values)
	//fmt.Println(tweet.Text)

	v := url.Values{}
	stream := t.UserStream(v) // 接続

	for {
		// 受信待ち
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				// Tweet を受信
				fmt.Printf("%#v", status.User.Id) //ユニークID
				fmt.Printf("%s: %s\n", status.User.ScreenName, status.Text) //@以降のID
				fmt.Printf("%s: %s\n", status.Text) //@以降のID
			case anaconda.StatusDeletionNotice:
			// pass
			default:
			}
		}
	}
}
