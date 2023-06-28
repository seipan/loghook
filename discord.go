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

type Discord struct {
	// This is the webhook url of the channel for which you want to send notifications to the discord.
	Webhook string

	// This is the webhook url for debug level log
	DebugWebhook string

	// This is the webhook url for info level log
	InfoWebhook string

	// This is the webhook url for warn level log
	WarnWebhook string

	// This is the webhook url for error level log
	ErrorWebhook string

	// This is the webhook url for panic level log
	PanicWebhook string

	// This is the webhook url for fatal level log
	FatalWebhook string
}

func NewDiscord(webhook string) *Discord {
	return &Discord{Webhook: webhook}
}

func (d *Discord) SetWebhook(webhook string) {
	d.Webhook = webhook
}

func (d *Discord) SetDebugWebhook(webhook string) {
	d.DebugWebhook = webhook
}

func (d *Discord) SetInfoWebhook(webhook string) {
	d.InfoWebhook = webhook
}

func (d *Discord) SetWarnWebhook(webhook string) {
	d.WarnWebhook = webhook
}

func (d *Discord) SetErrorWebhook(webhook string) {
	d.ErrorWebhook = webhook
}

func (d *Discord) SetPanicWebhook(webhook string) {
	d.PanicWebhook = webhook
}

func (d *Discord) SetFatalWebhook(webhook string) {
	d.FatalWebhook = webhook
}
