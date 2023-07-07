## Example

```go
package slack

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

## Result 
![image](https://github.com/seipan/loghook/assets/88176012/958fdb38-a61d-4668-85c9-3166595db1a8)
