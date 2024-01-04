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

package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordHandler struct {
	WebhookURL string

	message string
	level   string
	user    string
	img     string
}

func NewDiscordHandler(webhookURL string, message string, level string, user string, img string) *DiscordHandler {
	return &DiscordHandler{
		WebhookURL: webhookURL,
		message:    message,
		level:      level,
		user:       user,
		img:        img,
	}
}

func (dh *DiscordHandler) Send(msg string) error {
	dw := setWebhookStruct(dh.user, dh.img)
	dw = setWebfookMessage(dw, msg, dh.level)
	return sendLogToDiscord(dh.WebhookURL, dw)
}

func sendLogToDiscord(whurl string, dw *discordWebhook) error {
	j, err := json.Marshal(dw)
	if err != nil {
		return fmt.Errorf("json marshal err: %w", err)
	}

	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	if err != nil {
		return fmt.Errorf("new request err: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("client do err: %w", err)
	}

	return nil
}

func setWebhookStruct(name string, img string) *discordWebhook {
	dw := &discordWebhook{
		UserName:  name,
		AvatarURL: img,
	}
	return dw
}

func setWebfookMessage(dis *discordWebhook, message string, level string) *discordWebhook {
	dis.Embeds = []discordEmbed{
		{
			Title: level,
			Desc:  message,
			Color: 0x550000,
		},
	}
	return dis
}
