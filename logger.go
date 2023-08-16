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
	"context"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/seipan/loghook/discord"
	"github.com/seipan/loghook/slack"
)

// This structure defines what is needed to output logs to any channel on discord or slack.
type Logger struct {
	level Level
	mutex sync.Mutex

	Types string

	Slack *Slack

	Discord *Discord

	// This is the url of the icon image of the bot that sends notifications to the discord
	// ex) https://cdn-ak.f.st-hatena.com/images/fotolife/h/hikiniku0115/20190806/20190806000644.png
	Img string

	// This is the name of the bot that will send notifications to the discord
	// ex) hogehoge
	Name string
}

func NewLogger(img string, name string, types string, webhook string) *Logger {
	return &Logger{
		level:   InfoLevel,
		Img:     img,
		Name:    name,
		Types:   types,
		Slack:   NewSlack(webhook),
		Discord: NewDiscord(webhook),
	}
}

func (l *Logger) check(level Level) bool {
	return (l.checkTypes() && l.Level() <= level)
}

func (l *Logger) checkTypes() bool {
	return (l.Types == "slack" || l.Types == "discord")
}

func (l *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&l.level), uint32(level))
}

// Sets the specified url in the webhook for each level

func (l *Logger) SetWebhook(webhook string) {
	if l.Types == "slack" {
		l.Slack.SetWebhook(webhook)
	} else {
		l.Discord.SetWebhook(webhook)
	}
}

func (l *Logger) SetDebugWebhook(webhook string) {
	if l.Types == "slack" {
		l.Slack.SetDebugWebhook(webhook)
	} else {
		l.Discord.SetDebugWebhook(webhook)
	}
}

func (l *Logger) SetInfoWebhook(webhook string) {
	if l.Types == "slack" {
		l.Slack.SetInfoWebhook(webhook)
	} else {
		l.Discord.SetInfoWebhook(webhook)
	}
}

func (l *Logger) SetWarnWebhook(webhook string) {
	if l.Types == "slack" {
		l.Slack.SetWarnWebhook(webhook)
	} else {
		l.Discord.SetWarnWebhook(webhook)
	}
}

func (l *Logger) SetErrorWebhook(webhook string) {
	if l.Types == "slack" {
		l.Slack.SetErrorWebhook(webhook)
	} else {
		l.Discord.SetErrorWebhook(webhook)
	}
}

func (l *Logger) SetPanicWebhook(webhook string) {
	if l.Types == "slack" {
		l.Slack.SetPanicWebhook(webhook)
	} else {
		l.Discord.SetPanicWebhook(webhook)
	}
}

func (l *Logger) SetFatalWebhook(webhook string) {
	if l.Types == "slack" {
		l.Slack.SetFatalWebhook(webhook)
	} else {
		l.Discord.SetFatalWebhook(webhook)
	}
}

func (l *Logger) Level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.level)))
}

func (l *Logger) resWebhookURLbyLevel(level Level) string {
	if l.Types == "slack" {
		switch level {
		case DebugLevel:
			return l.Slack.DebugWebhook()
		case InfoLevel:
			return l.Slack.InfoWebhook()
		case WarnLevel:
			return l.Slack.WarnWebhook()
		case ErrorLevel:
			return l.Slack.ErrorWebhook()
		case PanicLevel:
			return l.Slack.PanicWebhook()
		case FatalLevel:
			return l.Slack.FatalWebhook()
		default:
			return "unknown"
		}
	} else {
		switch level {
		case DebugLevel:
			return l.Discord.DebugWebhook()
		case InfoLevel:
			return l.Discord.InfoWebhook()
		case WarnLevel:
			return l.Discord.WarnWebhook()
		case ErrorLevel:
			return l.Discord.ErrorWebhook()
		case PanicLevel:
			return l.Discord.PanicWebhook()
		case FatalLevel:
			return l.Discord.FatalWebhook()
		default:
			return "unknown"
		}
	}
}

func (l *Logger) Webhook() string {
	if l.Types == "slack" {
		return l.Slack.Webhook()
	} else {
		return l.Discord.Webhook()
	}
}

// nosend webhook method.
func (l *Logger) NoSendWebhook() {
	if l.Types == "slack" {
		l.Slack.SetWebhook("nosend")
	} else {
		l.Discord.SetWebhook("nosend")
	}
}

func (l *Logger) NoSendInfo() {
	if l.Types == "slack" {
		l.Slack.SetInfoWebhook("nosend")
	} else {
		l.Discord.SetInfoWebhook("nosend")
	}
}

func (l *Logger) NoSendDebug() {
	if l.Types == "slack" {
		l.Slack.SetDebugWebhook("nosend")
	} else {
		l.Discord.SetDebugWebhook("nosend")
	}
}

func (l *Logger) NoSendWarn() {
	if l.Types == "slack" {
		l.Slack.SetWarnWebhook("nosend")
	} else {
		l.Discord.SetWarnWebhook("nosend")
	}
}

func (l *Logger) NoSendError() {
	if l.Types == "slack" {
		l.Slack.SetErrorWebhook("nosend")
	} else {
		l.Discord.SetErrorWebhook("nosend")
	}
}

func (l *Logger) NoSendPanic() {
	if l.Types == "slack" {
		l.Slack.SetPanicWebhook("nosend")
	} else {
		l.Discord.SetPanicWebhook("nosend")
	}
}

func (l *Logger) NoSendFatal() {
	if l.Types == "slack" {
		l.Slack.SetFatalWebhook("nosend")
	} else {
		l.Discord.SetFatalWebhook("nosend")
	}
}

func (l *Logger) Log(level Level, args ...interface{}) {
	if l.check(level) {
		message := ""
		message = fmt.Sprint(args...)

		text := fmt.Sprintf("{\"time\": %s , \"level\": %s , \"message\" : %s}", time.Now().Format("2006-01-02 15:04:05"), level.String(), message)

		l.mutex.Lock()
		defer l.mutex.Unlock()
		fmt.Println(text)

		webhook := l.resWebhookURLbyLevel(level)
		if webhook == "nosend" {
			return
		} else if webhook == "unknown" || webhook == "" {
			webhook = l.Webhook()
		}

		// send log to discord or slack
		if l.Types == "discord" {
			dis := discord.SetWebhookStruct(l.Name, l.Img)
			dis = discord.SetWebfookMessage(dis, text, level.String())
			err := discord.SendLogToDiscord(webhook, dis)
			if err != nil {
				fmt.Printf("failed to send log to discord: %v\n", err)
			}
		} else if l.Types == "slack" {
			sl := slack.SetWebfookMessage(text, level.String(), l.Name, l.Img)
			err := slack.SendLogToSlack(webhook, sl)
			if err != nil {
				fmt.Printf("failed to send log to slack: %v\n", err)
			}
		}

	}
}

func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	if l.check(level) {
		message := ""
		message = fmt.Sprintf(format, args...)

		text := fmt.Sprintf("{\"time\": %s , \"level\": %s , \"message\" : %s}", time.Now().Format("2006-01-02 15:04:05"), level.String(), message)

		l.mutex.Lock()
		defer l.mutex.Unlock()
		fmt.Println(text)

		webhook := l.resWebhookURLbyLevel(level)
		if webhook == "nosend" {
			return
		} else if webhook == "unknown" || webhook == "" {
			webhook = l.Webhook()
		}

		// send log to discord or slack
		if l.Types == "discord" {
			dis := discord.SetWebhookStruct(l.Name, l.Img)
			dis = discord.SetWebfookMessage(dis, text, level.String())
			err := discord.SendLogToDiscord(webhook, dis)
			if err != nil {
				fmt.Printf("failed to send log to discord: %v\n", err)
			}
		} else if l.Types == "slack" {
			sl := slack.SetWebfookMessage(text, level.String(), l.Name, l.Img)
			err := slack.SendLogToSlack(webhook, sl)
			if err != nil {
				fmt.Printf("failed to send log to slack: %v\n", err)
			}
		}
	}
}

func (l *Logger) Info(i ...interface{}) {
	l.Log(InfoLevel, i...)

}

func (l *Logger) Infof(format string, i ...interface{}) {
	l.Logf(InfoLevel, format, i...)
}

func (l *Logger) InfoContext(ctx context.Context, format string, i ...interface{}) {
	l.Logf(InfoLevel, format, i...)
}

func (l *Logger) Debug(i ...interface{}) {
	l.Log(DebugLevel, i...)
}

func (l *Logger) Debugf(format string, i ...interface{}) {
	l.Logf(DebugLevel, format, i...)
}

func (l *Logger) DebugContext(ctx context.Context, i ...interface{}) {
	l.Log(DebugLevel, i...)
}

func (l *Logger) Error(i ...interface{}) {
	l.Log(ErrorLevel, i...)
}

func (l *Logger) Errorf(format string, i ...interface{}) {
	l.Logf(ErrorLevel, format, i...)
}

func (l *Logger) ErrorContext(ctx context.Context, i ...interface{}) {
	l.Log(ErrorLevel, i...)
}

func (l *Logger) Warn(i ...interface{}) {
	l.Log(WarnLevel, i...)
}

func (l *Logger) Warnf(format string, i ...interface{}) {
	l.Logf(WarnLevel, format, i...)
}

func (l *Logger) WarnContext(ctx context.Context, i ...interface{}) {
	l.Log(WarnLevel, i...)
}

func (l *Logger) Fatal(i ...interface{}) {
	l.Log(FatalLevel, i...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, i ...interface{}) {
	l.Logf(FatalLevel, format, i...)
	os.Exit(1)
}

func (l *Logger) FatalContext(ctx context.Context, i ...interface{}) {
	l.Log(FatalLevel, i...)
	os.Exit(1)
}

func (l *Logger) Panic(i ...interface{}) {
	l.Log(PanicLevel, i...)
	panic(fmt.Sprint(i...))
}

func (l *Logger) Panicf(format string, i ...interface{}) {
	l.Logf(PanicLevel, format, i...)
	panic(fmt.Sprintf(format, i...))
}

func (l *Logger) PanicContext(ctx context.Context, i ...interface{}) {
	l.Log(PanicLevel, i...)
	panic(fmt.Sprint(i...))
}
