package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gen2brain/beeep"
)

type Waring interface {
	Show(message string)
}

type ConsoleWarning struct{}

func (c ConsoleWarning) Show(message string) {
	fmt.Fprintf(os.Stderr, "[%s]: %s\n", os.Args[0], message)
}

type DesktopWarning struct{}

func (d DesktopWarning) Show(message string) {
	beeep.Alert(os.Args[0], message)
}

type SlackWarning struct {
	URL     string
	Channel string
}

type slackMessage struct {
	Text      string `json:text`
	Username  string `json:username`
	IconEmoji string `json:icon_emoji`
	Channel   string `json:channel`
}

func (s SlackWarning) Show(message string) {
	params, _ := json.Marshal(slackMessage{
		Text:      message,
		Username:  os.Args[0],
		IconEmoji: ":robot_face:",
		Channel:   s.Channel,
	})

	resp, err := http.PostForm(
		s.URL,
		url.Values{"payload": string(params)},
	)
	if err == nil {
		io.ReadAll(resp.Body)
		defer resp.Body.Close()
	}
}

func main() {
	var warn Warning

	warn = &ConsoleWarning{}
	warn.Show("Hello world to console")

	warn = &DesktopWarning{}
	warn.Show("Hello world to desktop")

	warn = &SlackWarning{
		URL:     os.Getenv("SLACK_URL"),
		Channel: "#general",
	}
	warn.Show("Hello world to slack")
}

// ダミー変数を使って、interface の実装がされているかなどチェックできる
var _ Warning = &DesktopWarning{}
