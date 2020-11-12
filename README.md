A simple Go library to send Discord webhooks

Sample usage: 

```go
webhook := webhook.DiscordWebhook{}

webhook.SetUsername("Username")
webhook.SetAvatarURL("https://upload.wikimedia.org/wikipedia/commons/2/20/Kielbasa.jpg")

embed := webhook.NewEmbed()

embed.SetAuthor("Sample author", "", "")
embed.SetDescription(t.ProductName)
embed.SetURL(t.ProductURL)

embed.AddField("This is a field", "This is the value", false)
embed.AddField("Another field", "Another value", false)

webhook.Send(WEBHOOK_URL_HERE)
```