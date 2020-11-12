package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DiscordWebhook : Holds data sent to discord to display a webhook
type DiscordWebhook struct {
	Content   string                 `json:"content,omitempty"`
	Username  string                 `json:"username,omitempty"`
	AvatarURL string                 `json:"avatar_url,omitempty"`
	Embeds    []*DiscordWebhookEmbed `json:"embeds"`
}

// DiscordWebhookEmbed : Holds data about a DiscordWebhook embed
type DiscordWebhookEmbed struct {
	Title       string    `json:"title,omitempty"`
	URL         string    `json:"url,omitempty"`
	Description string    `json:"description,omitempty"`
	Color       int       `json:"color,omitempty"`
	Timestamp   string    `json:"timestamp,omitempty"`
	Footer      footer    `json:"footer"`
	Thumbnail   thumbnail `json:"thumbnail"`
	Image       image     `json:"image,omitempty"`
	Author      author    `json:"author"`
	Fields      []field   `json:"fields,omitempty"`
}

type footer struct {
	IconURL string `json:"icon_url,omitempty"`
	Text    string `json:"text,omitempty"`
}

type thumbnail struct {
	URL string `json:"url,omitempty"`
}

type image struct {
	URL string `json:"url,omitempty"`
}

type author struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

// SetContent : Allows you to set the webhook content
func (d *DiscordWebhook) SetContent(content string) {
	d.Content = content
}

// SetUsername : Allows you to set the webhook username
func (d *DiscordWebhook) SetUsername(username string) {
	d.Username = username
}

// SetAvatarURL : Allows you to set the webhook avatar URL
func (d *DiscordWebhook) SetAvatarURL(avatarURL string) {
	d.AvatarURL = avatarURL
}

// NewEmbed : Adds an DiscordWebhookEmbed to the discord webhook and returns it
func (d *DiscordWebhook) NewEmbed() *DiscordWebhookEmbed {
	DiscordWebhookEmbed := &DiscordWebhookEmbed{}

	d.Embeds = append(d.Embeds, DiscordWebhookEmbed)
	return DiscordWebhookEmbed
}

// SetDescription : Allows you to set the DiscordWebhookEmbed description
func (e *DiscordWebhookEmbed) SetDescription(description string) {
	e.Description = description
}

// SetTitle : Allows you to set the DiscordWebhookEmbed title
func (e *DiscordWebhookEmbed) SetTitle(title string) {
	e.Title = title
}

// SetURL : Allows you to set the DiscordWebhookEmbed URL
func (e *DiscordWebhookEmbed) SetURL(URL string) {
	e.URL = URL
}

// SetColour : Allows you to set the DiscordWebhookEmbed colour
func (e *DiscordWebhookEmbed) SetColour(colourCode int) {
	e.Color = colourCode
}

// SetTimestamp : Sets the discord embed time to the current time
func (e *DiscordWebhookEmbed) SetTimestamp() {
	e.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05-0700")
}

// SetFooter : Allows you to set the DiscordWebhookEmbed footer
func (e *DiscordWebhookEmbed) SetFooter(text string, icon string) {
	e.Footer.IconURL = icon
	e.Footer.Text = text
}

// SetImage : Allows you to set the DiscordWebhookEmbed image
func (e *DiscordWebhookEmbed) SetImage(icon string) {
	e.Image.URL = icon
}

// SetThumbnail : Allows you to set the DiscordWebhookEmbed thumbnail
func (e *DiscordWebhookEmbed) SetThumbnail(icon string) {
	e.Thumbnail.URL = icon
}

// SetAuthor : Allows you to set the DiscordWebhookEmbed author
func (e *DiscordWebhookEmbed) SetAuthor(name string, URL string, iconURL string) {
	e.Author.Name = name
	e.Author.URL = URL
	e.Author.IconURL = iconURL
}

// AddField : Allows you to add a field to the discord DiscordWebhookEmbed
func (e *DiscordWebhookEmbed) AddField(name string, value string, inline bool) {
	field := field{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
	e.Fields = append(e.Fields, field)
}

// Send : Sends the discord webhook
func (d *DiscordWebhook) Send(webhookURL string) (bool, error) {
	jsonData, err := json.Marshal(d)

	if err != nil {
		return false, err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return true, nil
	}

	return false, fmt.Errorf("Bad status code - %v", resp.StatusCode)
}
