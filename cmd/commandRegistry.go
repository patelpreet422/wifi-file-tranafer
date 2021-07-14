package cmd

type Command = string

type CommandRegistry struct {
	Receive Command
	Send    Command
	None    Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		Receive: "receive",
		Send:    "send",
		None:    "",
	}
}
