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

// This structure slack holds the webhook url for logging to slack.
// This webhook assumes that the Incoming Webhook is used.
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

// Sets the webhook url.
// This url will be used if the webhook url for any level is not set
func (s *Slack) SetWebhook(webhook string) {
	s.Webhook = webhook
}

// Set the webhookurl for the Debug level.
func (s *Slack) SetDebugWebhook(webhook string) {
	s.DebugWebhook = webhook
}

// Set the webhookurl for the Info level.
func (s *Slack) SetInfoWebhook(webhook string) {
	s.InfoWebhook = webhook
}

// Set the webhookurl for the Warn level.
func (s *Slack) SetWarnWebhook(webhook string) {
	s.WarnWebhook = webhook
}

// Set the webhookurl for the Error level.
func (s *Slack) SetErrorWebhook(webhook string) {
	s.ErrorWebhook = webhook
}

// Set the webhookurl for the Panic level.
func (s *Slack) SetPanicWebhook(webhook string) {
	s.PanicWebhook = webhook
}

// Set the webhookurl for the Fatal level.
func (s *Slack) SetFatalWebhook(webhook string) {
	s.FatalWebhook = webhook
}
