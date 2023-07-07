## Example
Create a WebhookuURL using Incoming Webhook
```go
package slack

import "github.com/seipan/loghook"

var (
	// DiscordWebhookURL is a webhook url for discord.
	SlackWebhookURL = "https://hooks.slack.com/services/xxxxx/xxxx/xxxxxxxxx"
)

func main() {
	logger := loghook.NewLogger("", "test", "slack", SlackWebhookURL)
	logger.SetLevel(loghook.DebugLevel)
	logger.SetWebhook(SlackWebhookURL)

	logger.Debug("test")
	logger.Infof("test %s", "info")

	logger.NoSendDebug()
	logger.Debug("test")
	logger.NoSendInfo()
	logger.Infof("test %s", "info")

	logger.SetErrorWebhook(SlackErrorWebhookURL)
	logger.Error("test")
}

```

## Result
![image](https://github.com/seipan/loghook/assets/88176012/ce577a95-4dc1-4e38-a1dc-96bfe643fc7a)
