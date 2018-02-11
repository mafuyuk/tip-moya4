package bot

// Command はbotのコマンドを表す
//go:generate enumer -type=Command -transform=kebab command.go
type Command int

const (
	Balance Command = iota
	Tip
	Deposit
	Donate
	Withdraw
)
