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
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"

	"github.com/seipan/loghook/discord"
)

// This structure defines what is needed to output logs to any channel on discord.
type Logger struct {
	level Level
	mutex sync.Mutex

	// This is the webhook url of the channel for which you want to send notifications to the discord.
	// ex) discord.com/api/webhooks/xxxxxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
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

	// This is the url of the icon image of the bot that sends notifications to the discord
	// ex) https://cdn-ak.f.st-hatena.com/images/fotolife/h/hikiniku0115/20190806/20190806000644.png
	Img string

	// This is the name of the bot that will send notifications to the discord
	// ex) hogehoge
	Name string
}

func NewLogger(webhook string, img string, name string) *Logger {
	return &Logger{
		level:   InfoLevel,
		Webhook: webhook,
		Img:     img,
		Name:    name,
	}
}

func (l *Logger) check(level Level) bool {
	return l.Level() <= level
}

func (l *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&l.level), uint32(level))
}

// Sets the specified url in the webhook for each level
func (l *Logger) SetDebugWebhook(webhook string) {
	l.DebugWebhook = webhook
}

func (l *Logger) SetInfoWebhook(webhook string) {
	l.InfoWebhook = webhook
}

func (l *Logger) SetWarnWebhook(webhook string) {
	l.WarnWebhook = webhook
}

func (l *Logger) SetErrorWebhook(webhook string) {
	l.ErrorWebhook = webhook
}

func (l *Logger) SetPanicWebhook(webhook string) {
	l.PanicWebhook = webhook
}

func (l *Logger) SetFatalWebhook(webhook string) {
	l.FatalWebhook = webhook
}

func (l *Logger) Level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.level)))
}

func (l *Logger) resWebhookURLbyLevel(level Level) string {
	switch level {
	case DebugLevel:
		return l.DebugWebhook
	case InfoLevel:
		return l.InfoWebhook
	case WarnLevel:
		return l.WarnWebhook
	case ErrorLevel:
		return l.ErrorWebhook
	case PanicLevel:
		return l.PanicWebhook
	case FatalLevel:
		return l.FatalWebhook
	default:
		return "unknown"
	}
}

func (l *Logger) Log(level Level, user string, args ...interface{}) {
	if l.check(level) {
		message := ""
		message = fmt.Sprint(args...)

		l.mutex.Lock()
		defer l.mutex.Unlock()
		log.Println(message)

		webhook := l.resWebhookURLbyLevel(level)
		if webhook == "unknown" || webhook == "" {
			webhook = l.Webhook
		}

		// send log to discord
		dis := discord.SetWebhookStruct(l.Name, l.Img)
		dis = discord.SetWebfookMessage(dis, message, user, level.String())
		discord.SendLogToDiscord(webhook, dis)
	}
}

func (l *Logger) Logf(level Level, user string, format string, args ...interface{}) {
	if l.check(level) {
		message := ""
		message = fmt.Sprintf(format, args...)

		l.mutex.Lock()
		defer l.mutex.Unlock()
		log.Println(message)

		webhook := l.resWebhookURLbyLevel(level)
		if webhook == "unknown" || webhook == "" {
			webhook = l.Webhook
		}

		// send log to discord
		dis := discord.SetWebhookStruct(l.Name, l.Img)
		dis = discord.SetWebfookMessage(dis, message, user, level.String())
		discord.SendLogToDiscord(webhook, dis)
	}
}

func (l *Logger) Info(user string, i ...interface{}) {
	l.Log(InfoLevel, user, i...)

}

func (l *Logger) Infof(user string, format string, i ...interface{}) {
	l.Logf(InfoLevel, user, format, i...)
}

func (l *Logger) Debug(user string, i ...interface{}) {
	l.Log(DebugLevel, user, i...)
}

func (l *Logger) Debugf(user string, format string, i ...interface{}) {
	l.Logf(DebugLevel, user, format, i...)
}

func (l *Logger) Error(user string, i ...interface{}) {
	l.Log(ErrorLevel, user, i...)
}

func (l *Logger) Errorf(user string, format string, i ...interface{}) {
	l.Logf(ErrorLevel, user, format, i...)
}

func (l *Logger) Warn(user string, i ...interface{}) {
	l.Log(WarnLevel, user, i...)
}

func (l *Logger) Warnf(user string, format string, i ...interface{}) {
	l.Logf(WarnLevel, user, format, i...)
}

func (l *Logger) Fatal(user string, i ...interface{}) {
	l.Log(FatalLevel, user, i...)
	os.Exit(1)
}

func (l *Logger) Fatalf(user string, format string, i ...interface{}) {
	l.Logf(FatalLevel, user, format, i...)
	os.Exit(1)
}

func (l *Logger) Panic(user string, i ...interface{}) {
	l.Log(PanicLevel, user, i...)
	panic(fmt.Sprint(i...))
}

func (l *Logger) Panicf(user string, format string, i ...interface{}) {
	l.Logf(PanicLevel, user, format, i...)
	panic(fmt.Sprintf(format, i...))
}
