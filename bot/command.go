package bot

//go:generate enumer stringer -type=Command -transform=kebab
type Command string

const (
	Balance Command = iota
	Tip
	Deposit
	Donate
	Withdraw
)
