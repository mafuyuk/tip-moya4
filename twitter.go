package tip_moya4

import "github.com/ChimeraCoder/anaconda"

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
