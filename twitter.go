package tip_moya4

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

type Twitter struct {
	Api *anaconda.TwitterApi
}

func NewTwitter(consumerKey, ConsumerSecret, accessToken, accessTokenSecret string) *Twitter {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	return &Twitter{
		Api: anaconda.NewTwitterApi(accessToken, accessTokenSecret),
	}
}

func (t *Twitter) Post(message string) (anaconda.Tweet, error){
	tweet, err := t.Api.PostTweet(message, nil)
	return tweet, err
}


func (t *Twitter) AccountActivity(v url.Values) {
	t.Api.GetActivityWebhooks(v)
}

func (t *Twitter) UserStream(v url.Values) *anaconda.Stream {
	return t.Api.UserStream(v)
}