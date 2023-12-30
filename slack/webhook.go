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

package slack

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type SlackHandler struct {
	WebhookURL string

	message string
	level   string
	user    string
	img     string
}

func (sh *SlackHandler) Send(msg string) error {
	sl := setWebfookMessage(msg, sh.level, sh.user, sh.img)
	return sendLogToSlack(sh.WebhookURL, sl)
}

func NewSlackHandler(webhookURL string, message string, level string, user string, img string) *SlackHandler {
	return &SlackHandler{
		WebhookURL: webhookURL,
		message:    message,
		level:      level,
		user:       user,
		img:        img,
	}
}

func setWebfookMessage(message string, level string, user string, img string) *Payload {
	sl := &Payload{
		Username: user,
		IconUrl:  img,
		Text:     message,
	}
	return sl
}

func sendLogToSlack(webhookURL string, sw *Payload) error {
	p, err := json.Marshal(sw)
	if err != nil {
		return err
	}
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(p)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
