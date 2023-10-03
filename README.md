<div align="center">

![Last commit](https://img.shields.io/github/last-commit/seipan/loghook?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/seipan/loghook?style=flat-square)
![Issues](https://img.shields.io/github/issues/seipan/loghook?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/seipan/loghook?style=flat-square)
[![go](https://github.com/seipan/loghook/actions/workflows/go.yml/badge.svg)](https://github.com/seipan/loghook/actions/workflows/go.yml)

<img src="img/logo.png" alt="eyecatch" height="250">

# loghook

⚡  logger to notify logs to slack,discord using webhook   ⚡

<br>
<br>


</div>

## Installation
```
go get github.com/seipan/loghook
```

## Usage
When using it, you need to obtain the default webhook for discord and the incoming webhook for slack in advance.
```go
package discord

import "github.com/seipan/loghook"

var (
	// DiscordWebhookURL is a webhook url for discord.
	DiscordWebhookURL = "https://discord.com/api/webhooks/xxxxxxxx/xxxxxxxx"
)

func main() {
	logger := loghook.NewLogger("", "test", "discord", DiscordWebhookURL)
	logger.SetLevel(loghook.DebugLevel)
	logger.SetWebhook(DiscordWebhookURL)

	logger.Debug("test")
	logger.Infof("test %s", "info")
}
```

If you do not want to be notified of a particular log level, you can set
```go
func main(){
	logger.NoSendDebug()
	logger.Debug("test")
	logger.NoSendInfo()
	logger.Infof("test %s", "info")
}
```
You can also change the webhook to be notified for each log level
```go
func main(){
	logger.SetErrorWebhook(DiscordErrorWebhookURL)
	logger.Error("test")
}
```
If you want a more detailed example, please see the [examples](https://github.com/seipan/loghook/blob/main/example).
## License
Code licensed under 
[the MIT License](https://github.com/seipan/loghook/blob/main/LICENSE).
