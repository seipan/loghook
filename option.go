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

// Option is a type of structure that holds types such as slack, discord, etc.
// You can freely customize this when you want to notify a different service in slack, discord, etc.
type Option struct {
	// slack discord, etc.
	types string

	Webhook Webhook

	Handler Handler
}

func (o *Option) Types() string {
	return o.types
}

func (o *Option) SetTypes(types string) {
	o.types = types
}

func (o *Option) SetWebhook(webhook Webhook) {
	o.Webhook = webhook
}

func (o *Option) SetHandler(handler Handler) {
	o.Handler = handler
}

func (o *Option) SetWebhookURL(url string) {
	o.Webhook.SetWebhook(url)
}

func (o *Option) SetDebugWebhookURL(url string) {
	o.Webhook.SetDebugWebhook(url)
}

func (o *Option) SetInfoWebhookURL(url string) {
	o.Webhook.SetInfoWebhook(url)
}

func (o *Option) SetWarnWebhookURL(url string) {
	o.Webhook.SetWarnWebhook(url)
}

func (o *Option) SetErrorWebhookURL(url string) {
	o.Webhook.SetErrorWebhook(url)
}

func (o *Option) SetPanicWebhookURL(url string) {
	o.Webhook.SetPanicWebhook(url)
}

func (o *Option) SetFatalWebhookURL(url string) {
	o.Webhook.SetFatalWebhook(url)
}

func (o *Option) WebhookURL() string {
	return o.Webhook.Webhook()
}

func (o *Option) DebugWebhookURL() string {
	return o.Webhook.DebugWebhook()
}

func (o *Option) InfoWebhookURL() string {
	return o.Webhook.InfoWebhook()
}

func (o *Option) WarnWebhookURL() string {
	return o.Webhook.WarnWebhook()
}

func (o *Option) ErrorWebhookURL() string {
	return o.Webhook.ErrorWebhook()
}

func (o *Option) PanicWebhookURL() string {
	return o.Webhook.PanicWebhook()
}

func (o *Option) FatalWebhookURL() string {
	return o.Webhook.FatalWebhook()
}

func NewOption(types string, webhook string) *Option {
	return &Option{
		types:   types,
		Webhook: *NewWebhook(webhook),
	}
}
