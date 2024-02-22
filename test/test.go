package main

import (
	"context"

	"github.com/seipan/loghook"
)

func eToeTest(discordWebhookURL string, slackWebhookURL string) error {
	testNosend()
	testDiscordSend(discordWebhookURL)
	testSendSlack(slackWebhookURL)
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
	logger.Debugf("test %s", "debug")
	logger.Info("test")
	logger.InfoContext(context.Background(), "test")
	logger.Infof("test %s", "info")
	logger.Warn("test")
	logger.WarnContext(context.Background(), "test")
	logger.Warnf("test %s", "warn")
	logger.Error("test")
	logger.ErrorContext(context.Background(), "test")
	logger.Errorf("test %s", "error")
	logger.Fatal("test")
	logger.FatalContext(context.Background(), "test")
	logger.Fatalf("test %s", "fatal")
}

func testSendSlack(slackWebhookURL string) {
	logger := loghook.NewLogger("", "test", "slack", slackWebhookURL)
	logger.SetLevel(loghook.DebugLevel)
	logger.SetWebhook("")

	logger.Debug("test")
	logger.DebugContext(context.Background(), "test")
	logger.Debugf("test %s", "debug")
	logger.Info("test")
	logger.InfoContext(context.Background(), "test")
	logger.Infof("test %s", "info")
	logger.Warn("test")
	logger.WarnContext(context.Background(), "test")
	logger.Warnf("test %s", "warn")
	logger.Error("test")
	logger.ErrorContext(context.Background(), "test")
	logger.Errorf("test %s", "error")
	logger.Fatal("test")
	logger.FatalContext(context.Background(), "test")
	logger.Fatalf("test %s", "fatal")
}
