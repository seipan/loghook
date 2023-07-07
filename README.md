## loghook
<div align="center">
<img src="img/logo.png" alt="属性" title="タイトル">
</div>

 logger to notify logs to slack,discord using webhook 
## Installation
```
go get github.com/seipan/loghook
```

## Usage
```go
package salck

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

	logger.NoSendDebug()
	logger.Debug("test")
	logger.NoSendInfo()
	logger.Infof("test %s", "info")

    logger.SetErrorWebhook(DiscordErrorWebhookURL)
	logger.Error("test")
}
```
