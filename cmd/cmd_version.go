package cmd

import "github.com/sbstjn/hanu"

func init() {
	Register(
		"version",
		"버전",
		func(conv hanu.ConversationInterface) {
			conv.Reply("I'm running with `%s`", Version)
		},
	)
}
