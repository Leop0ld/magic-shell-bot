package cmd

import (
	"math/rand"

	"github.com/sbstjn/hanu"
)

func init() {
	Register(
		"마법의 소라고둥님 (.*)",
		"대답",
		func(conv hanu.ConversationInterface) {
			answers := make([]string, 0)
			answers = append(answers,
				"해.",
				"하지마.",
				"그만둬.",
				"시작해.")
		
			randIndex := rand.Int() % len(answers)

			conv.Reply(answers[randIndex])
		},
	)
}
