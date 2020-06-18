package main

import (
	"testing"
)

func TestGenerateBlock(t *testing.T) {
	emojis = true
	emojiBlock := generateBlock("🚨 ERROR 🚨", "░")
	if emojiBlock != "░░░░░░░░░░░░░\n░🚨 ERROR 🚨░\n░░░░░░░░░░░░░" {
		t.Errorf(emojiBlock, "!= ░░░░░░░░░░░░░\n░🚨 ERROR 🚨░\n░░░░░░░░░░░░░")
	}

	emojis = false
	block := generateBlock("ERROR", "░")
	if block != "░░░░░░░\n░ERROR░\n░░░░░░░" {
		t.Errorf(block, "!= ░░░░░░░\n░ERROR░\n░░░░░░░")
	}
}
