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

type Slack struct {
	// This is the webhook url of the channel for which you want to send notifications to the slack.
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

func NewSlack(webhook string) *Slack {
	return &Slack{Webhook: webhook}
}

func (s *Slack) SetWebhook(webhook string) {
	s.Webhook = webhook
}

func (s *Slack) SetDebugWebhook(webhook string) {
	s.DebugWebhook = webhook
}

func (s *Slack) SetInfoWebhook(webhook string) {
	s.InfoWebhook = webhook
}

func (s *Slack) SetWarnWebhook(webhook string) {
	s.WarnWebhook = webhook
}

func (s *Slack) SetErrorWebhook(webhook string) {
	s.ErrorWebhook = webhook
}

func (s *Slack) SetPanicWebhook(webhook string) {
	s.PanicWebhook = webhook
}

func (s *Slack) SetFatalWebhook(webhook string) {
	s.FatalWebhook = webhook
}
