package tof

import (
	"encoding/json"
)

const botAPI = "/ebus/tof4_msg/api/v1/Message/SendWeComRobotInfo"

type bot struct {
	Key      string   `json:"Sender"`   // 发件人, 即 bot key
	Receiver []string `json:"Receiver"` // 接收人的英文名
	Type     string   `json:"Type"`     // 消息类型
	Markdown markdown `json:"Markdown"` // markdown 消息
}

type markdown struct {
	Content string `json:"content"`
}

// Send 发送企业微信机器人消息, 实现 tof.Sender 接口
func (b *bot) Send(passID, token, env string) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return send(botAPI, passID, token, env, data)
}

// NewBot 新建一个企业微信机器人发送
func NewBot(opts ...botOption) Sender {
	b := &bot{}
	for _, o := range opts {
		o(b)
	}
	return b
}

// botOption 企业微信发送选项
type botOption func(b *bot)

// WithKey 设置发件人, 即 bot key
func WithKey(botKey string) botOption {
	return func(b *bot) {
		b.Key = botKey
	}
}

// WithReceiver 设置收件人
func WithReceiver(reciver []string) botOption {
	return func(b *bot) {
		b.Receiver = reciver
	}
}

// WithType 设置消息类型
func WithType(t string) botOption {
	return func(b *bot) {
		b.Type = t
	}
}

// WithMarkdownContent 设置 markdown 内容
func WithMarkdownContent(content string) botOption {
	return func(b *bot) {
		b.Markdown.Content = content
	}
}
