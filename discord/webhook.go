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

func SendLogToDiscord(whurl string, dw *discordWebhook) {
	j, err := json.Marshal(dw)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("new request err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	_, err = client.Do(req)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
}

func setWebhookStruct(name string, img string) *discordWebhook {
	dw := &discordWebhook{
		UserName:  name,
		AvatarURL: img,
	}
	return dw
}

func setWebfookMessage(dis *discordWebhook, message string, user string, level string) *discordWebhook {
	dis.Content = user
	dis.Embeds = []discordEmbed{
		discordEmbed{
			Title: level,
			Desc:  message,
			Color: 0x550000,
		},
	}
	return dis
}
