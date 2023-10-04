// MIT License

// Copyright (c) 2023 Yamasaki Shotaro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package loghook

import (
	"context"
	"testing"
)

var (
	// DiscordWebhookURL is a webhook url for discord.
	DiscordWebhookURL = "https://discord.com/api/webhooks/xxxxxxxx/xxxxxxxx"

	// SlackWebhookURL is a webhook url for slack.
	SlackWebhookURL = "https://hooks.slack.com/services/xxxxx/xxxx/xxxxxxxxx"
)

func TestDiscordExample(t *testing.T) {
	logger := NewLogger("", "test", "discord", DiscordWebhookURL)
	logger.SetLevel(DebugLevel)
	logger.SetWebhook(DiscordWebhookURL)

	logger.NoSendDebug()
	logger.Debug("test")
	logger.DebugContext(context.Background(), "test")
	logger.NoSendInfo()
	logger.Infof("test %s", "info")
	logger.InfoContext(context.Background(), "test")
}

func TestSlackExample(t *testing.T) {
	logger := NewLogger("", "test", "slack", SlackWebhookURL)
	logger.SetLevel(DebugLevel)
	logger.SetWebhook(SlackWebhookURL)

	logger.NoSendDebug()
	logger.Debugf("test %s", "debug")
	logger.DebugContext(context.Background(), "test")
}
