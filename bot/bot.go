package bot

import (
	"strings"
	"errors"
	"fmt"
)

const moya4BotAccountID = "mafuyuk_m"

type Bot struct {
	Command

	RequesterTwitterID  uint64
	RequesterAccountID string

	ReceiverTwitterID  uint64
	ReceiverAddress  string
}

var ErrNotExecBot = errors.New("bot: not execute bot")

func Exec(bot *Bot) {
	switch bot.Command {
	case Balance:
		fmt.Println("exec balance")
		bot.Balance()
	case Tip:
		fmt.Println("exec tip")
	}
}

// NewはTweet内容がBotへのコマンド実行だった場合に必要な情報を構造体にして返す
func New(twitterID uint64, twitterAccountID,tweet string) (*Bot, error) {
	splitRes := strings.Split(tweet, " ")

	if splitRes[0] != "@" + moya4BotAccountID {
		return nil, fmt.Errorf("It is not a mention to bot: %s", ErrNotExecBot)
	}

	return &Bot{
		RequesterTwitterID: twitterID,
		RequesterAccountID: twitterAccountID,
	}, nil
}

func(*Bot) Balance() {
	fmt.Println("Balance!")
}