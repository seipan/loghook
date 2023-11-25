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

// This structure Discord holds the webhook url for logging to discord.
type Discord struct {
	// This is the webhook url of the channel for which you want to send notifications to the discord.
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

// NewDiscord returns a structure Discord that manages the webhooks of Discord.
func NewDiscord(webhook string) *Discord {
	return &Discord{webhook: webhook}
}

// These return each webhook of Discord.
// If the webhook url for any level is not set, the webhook url set by SetWebhook() will be used.
func (d *Discord) Webhook() string {
	return d.webhook
}

// It returns the webhook for the debug level.
func (d *Discord) DebugWebhook() string {
	return d.debugWebhook
}

// It returns the webhook for the info level.
func (d *Discord) InfoWebhook() string {
	return d.infoWebhook
}

// It returns the webhook for the warn level.
func (d *Discord) WarnWebhook() string {
	return d.warnWebhook
}

// It returns the webhook for the error level.
func (d *Discord) ErrorWebhook() string {
	return d.errorWebhook
}

// It returns the webhook for the panic level.
func (d *Discord) PanicWebhook() string {
	return d.panicWebhook
}

// It returns the webhook for the fatal level.
func (d *Discord) FatalWebhook() string {
	return d.fatalWebhook
}

// Sets the webhook url.
// This url will be used if the webhook url for any level is not set
func (d *Discord) SetWebhook(webhook string) {
	d.webhook = webhook
}

// Set the webhookurl for the Debug level.
func (d *Discord) SetDebugWebhook(webhook string) {
	d.debugWebhook = webhook
}

// Set the webhookurl for the Info level.
func (d *Discord) SetInfoWebhook(webhook string) {
	d.infoWebhook = webhook
}

// Set the webhookurl for the Warn level.
func (d *Discord) SetWarnWebhook(webhook string) {
	d.warnWebhook = webhook
}

// Set the webhookurl for the Error level.
func (d *Discord) SetErrorWebhook(webhook string) {
	d.errorWebhook = webhook
}

// Set the webhookurl for the Panic level.
func (d *Discord) SetPanicWebhook(webhook string) {
	d.panicWebhook = webhook
}

// Set the webhookurl for the Fatal level.
func (d *Discord) SetFatalWebhook(webhook string) {
	d.fatalWebhook = webhook
}
