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
type Webhook struct {
	// This is the webhook url of the channel for which you want to send notifications to the slack.
	webhook string

	// This is the webhook url for debug level log
	debugWebhook string

	// This is the webhook url for info level log
	infoWebhook string

	// This is the webhook url for warn level log
	warnWebhook string

	// This is the webhook url for error level log
	errorWebhook string

	// This is the webhook url for panic level log
	panicWebhook string

	// This is the webhook url for fatal level log
	fatalWebhook string
}

func NewWebhook(webhook string) *Webhook {
	return &Webhook{webhook: webhook}
}

func (s *Webhook) Webhook() string {
	return s.webhook
}

func (s *Webhook) DebugWebhook() string {
	return s.debugWebhook
}

func (s *Webhook) InfoWebhook() string {
	return s.infoWebhook
}

func (s *Webhook) WarnWebhook() string {
	return s.warnWebhook
}

func (s *Webhook) ErrorWebhook() string {
	return s.errorWebhook
}

func (s *Webhook) PanicWebhook() string {
	return s.panicWebhook
}

func (s *Webhook) FatalWebhook() string {
	return s.fatalWebhook
}

// Sets the webhook url.
// This url will be used if the webhook url for any level is not set
func (s *Webhook) SetWebhook(webhook string) {
	s.webhook = webhook
}

// Set the webhookurl for the Debug level.
func (s *Webhook) SetDebugWebhook(webhook string) {
	s.debugWebhook = webhook
}

// Set the webhookurl for the Info level.
func (s *Webhook) SetInfoWebhook(webhook string) {
	s.infoWebhook = webhook
}

// Set the webhookurl for the Warn level.
func (s *Webhook) SetWarnWebhook(webhook string) {
	s.warnWebhook = webhook
}

// Set the webhookurl for the Error level.
func (s *Webhook) SetErrorWebhook(webhook string) {
	s.errorWebhook = webhook
}

// Set the webhookurl for the Panic level.
func (s *Webhook) SetPanicWebhook(webhook string) {
	s.panicWebhook = webhook
}

// Set the webhookurl for the Fatal level.
func (s *Webhook) SetFatalWebhook(webhook string) {
	s.fatalWebhook = webhook
}
