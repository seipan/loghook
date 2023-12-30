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

	Types Option

	SendLevel Level

	// This is the url of the icon image of the bot that sends notifications to the discord
	// ex) https://cdn-ak.f.st-hatena.com/images/fotolife/h/hikiniku0115/20190806/20190806000644.png
	Img string

	// This is the name of the bot that will send notifications to the discord
	// ex) hogehoge
	Name string
}

func NewLogger(img string, name string, types string, webhook string) *Logger {
	return &Logger{
		level: InfoLevel,
		Img:   img,
		Name:  name,
		Types: *NewOption(types, ""),
	}
}

func (l *Logger) check(ctx context.Context, level Level) bool {
	return (l.checkTypes() && l.Level() <= level)
}

func (l *Logger) checkTypes() bool {
	return (l.Types.Types() == "slack" || l.Types.Types() == "discord")
}

func (l *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&l.level), uint32(level))
}

func (l *Logger) SetSendLevel(level Level) {
	atomic.StoreUint32((*uint32)(&l.SendLevel), uint32(level))
}

func (l *Logger) setNosend() {
	switch l.SendLevel {
	case DebugLevel:
		l.NoSendDebug()
	case InfoLevel:
		l.NoSendDebug()
		l.NoSendInfo()
	case WarnLevel:
		l.NoSendDebug()
		l.NoSendInfo()
		l.NoSendWarn()
	case ErrorLevel:
		l.NoSendDebug()
		l.NoSendInfo()
		l.NoSendWarn()
		l.NoSendError()
	case PanicLevel:
		l.NoSendDebug()
		l.NoSendInfo()
		l.NoSendWarn()
		l.NoSendError()
		l.NoSendPanic()
	case FatalLevel:
		l.NoSendDebug()
		l.NoSendInfo()
		l.NoSendWarn()
		l.NoSendError()
		l.NoSendPanic()
		l.NoSendFatal()
	}
}

// Sets the specified url in the webhook for each level
func (l *Logger) SetWebhook(webhook string) {
	l.Types.SetWebhookURL(webhook)
}

func (l *Logger) SetDebugWebhook(webhook string) {
	l.Types.SetDebugWebhookURL(webhook)
}

func (l *Logger) SetInfoWebhook(webhook string) {
	l.Types.SetInfoWebhookURL(webhook)
}

func (l *Logger) SetWarnWebhook(webhook string) {
	l.Types.SetWarnWebhookURL(webhook)
}

func (l *Logger) SetErrorWebhook(webhook string) {
	l.Types.SetErrorWebhookURL(webhook)
}

func (l *Logger) SetPanicWebhook(webhook string) {
	l.Types.SetPanicWebhookURL(webhook)
}

func (l *Logger) SetFatalWebhook(webhook string) {
	l.Types.SetFatalWebhookURL(webhook)
}

func (l *Logger) Level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.level)))
}

func (l *Logger) resWebhookURLbyLevel(level Level) string {
	switch level {
	case DebugLevel:
		return l.Types.DebugWebhookURL()
	case InfoLevel:
		return l.Types.InfoWebhookURL()
	case WarnLevel:
		return l.Types.WarnWebhookURL()
	case ErrorLevel:
		return l.Types.ErrorWebhookURL()
	case PanicLevel:
		return l.Types.PanicWebhookURL()
	case FatalLevel:
		return l.Types.FatalWebhookURL()
	default:
		return "unknown"
	}
}

func (l *Logger) Webhook() string {
	return l.Types.WebhookURL()
}

// nosend webhook method.
func (l *Logger) NoSendWebhook() {
	l.Types.SetWebhookURL("nosend")
}

func (l *Logger) NoSendInfo() {
	l.Types.SetInfoWebhookURL("nosend")
}

func (l *Logger) NoSendDebug() {
	l.Types.SetDebugWebhookURL("nosend")
}

func (l *Logger) NoSendWarn() {
	l.Types.SetWarnWebhookURL("nosend")
}

func (l *Logger) NoSendError() {
	l.Types.SetErrorWebhookURL("nosend")
}

func (l *Logger) NoSendPanic() {
	l.Types.SetPanicWebhookURL("nosend")
}

func (l *Logger) NoSendFatal() {
	l.Types.SetFatalWebhookURL("nosend")
}

func (l *Logger) Log(ctx context.Context, level Level, args ...interface{}) {
	if l.check(ctx, level) {
		l.setNosend()
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
		if l.Types.Types() == "discord" {
			dis := discord.SetWebhookStruct(l.Name, l.Img)
			dis = discord.SetWebfookMessage(dis, text, level.String())
			err := discord.SendLogToDiscord(webhook, dis)
			if err != nil {
				fmt.Printf("failed to send log to discord: %v\n", err)
			}
		} else if l.Types.Types() == "slack" {
			sl := slack.SetWebfookMessage(text, level.String(), l.Name, l.Img)
			err := slack.SendLogToSlack(webhook, sl)
			if err != nil {
				fmt.Printf("failed to send log to slack: %v\n", err)
			}
		}

	}
}

func (l *Logger) Logf(ctx context.Context, level Level, format string, args ...interface{}) {
	if l.check(ctx, level) {
		l.setNosend()
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
		if l.Types.Types() == "discord" {
			dis := discord.SetWebhookStruct(l.Name, l.Img)
			dis = discord.SetWebfookMessage(dis, text, level.String())
			err := discord.SendLogToDiscord(webhook, dis)
			if err != nil {
				fmt.Printf("failed to send log to discord: %v\n", err)
			}
		} else if l.Types.Types() == "slack" {
			sl := slack.SetWebfookMessage(text, level.String(), l.Name, l.Img)
			err := slack.SendLogToSlack(webhook, sl)
			if err != nil {
				fmt.Printf("failed to send log to slack: %v\n", err)
			}
		}
	}
}

func (l *Logger) Info(i ...interface{}) {
	l.Log(context.Background(), InfoLevel, i...)

}

func (l *Logger) Infof(format string, i ...interface{}) {
	l.Logf(context.Background(), InfoLevel, format, i...)
}

func (l *Logger) InfoContext(ctx context.Context, format string, i ...interface{}) {
	l.Logf(ctx, InfoLevel, format, i...)
}

func (l *Logger) Debug(i ...interface{}) {
	l.Log(context.Background(), DebugLevel, i...)
}

func (l *Logger) Debugf(format string, i ...interface{}) {
	l.Logf(context.Background(), DebugLevel, format, i...)
}

func (l *Logger) DebugContext(ctx context.Context, i ...interface{}) {
	l.Log(ctx, DebugLevel, i...)
}

func (l *Logger) Error(i ...interface{}) {
	l.Log(context.Background(), ErrorLevel, i...)
}

func (l *Logger) Errorf(format string, i ...interface{}) {
	l.Logf(context.Background(), ErrorLevel, format, i...)
}

func (l *Logger) ErrorContext(ctx context.Context, i ...interface{}) {
	l.Log(ctx, ErrorLevel, i...)
}

func (l *Logger) Warn(i ...interface{}) {
	l.Log(context.Background(), WarnLevel, i...)
}

func (l *Logger) Warnf(format string, i ...interface{}) {
	l.Logf(context.Background(), WarnLevel, format, i...)
}

func (l *Logger) WarnContext(ctx context.Context, i ...interface{}) {
	l.Log(ctx, WarnLevel, i...)
}

func (l *Logger) Fatal(i ...interface{}) {
	l.Log(context.Background(), FatalLevel, i...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, i ...interface{}) {
	l.Logf(context.Background(), FatalLevel, format, i...)
	os.Exit(1)
}

func (l *Logger) FatalContext(ctx context.Context, i ...interface{}) {
	l.Log(ctx, FatalLevel, i...)
	os.Exit(1)
}

func (l *Logger) Panic(i ...interface{}) {
	l.Log(context.Background(), PanicLevel, i...)
	panic(fmt.Sprint(i...))
}

func (l *Logger) Panicf(format string, i ...interface{}) {
	l.Logf(context.Background(), PanicLevel, format, i...)
	panic(fmt.Sprintf(format, i...))
}

func (l *Logger) PanicContext(ctx context.Context, i ...interface{}) {
	l.Log(ctx, PanicLevel, i...)
	panic(fmt.Sprint(i...))
}
