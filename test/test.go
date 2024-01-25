package main

import (
	"context"

	"github.com/seipan/loghook"
)

func eToeTest(discordWebhookURL string, slackWebhookURL string) error {
	testNosend()
	testDiscordSend(discordWebhookURL)

	return nil
}

func testNosend() {
	logger := loghook.NewLogger("", "test", "discord", "")
	logger.SetLevel(loghook.DebugLevel)
	logger.SetWebhook("")

	logger.NoSendDebug()
	logger.Debug("test")
	logger.DebugContext(context.Background(), "test")
	logger.NoSendInfo()
	logger.Infof("test %s", "info")
}

func testDiscordSend(discordWebhookURL string) {
	logger := loghook.NewLogger("", "test", "discord", discordWebhookURL)
	logger.SetLevel(loghook.DebugLevel)
	logger.SetWebhook(discordWebhookURL)

	logger.Debug("test")
	logger.DebugContext(context.Background(), "test")
	logger.Infof("test %s", "info")
}
