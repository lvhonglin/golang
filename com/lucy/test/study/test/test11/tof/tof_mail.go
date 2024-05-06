package tof

import (
	"encoding/json"
)

const mailAPI = "/ebus/tof4_msg/api/v1/Message/SendMailInfo"

type mail struct {
	To         string `json:"To"`         // 邮件接受者
	CC         string `json:"CC"`         // cc抄送
	Bcc        string `json:"Bcc"`        // bcc
	From       string `json:"From"`       // 发送者
	Title      string `json:"Title"`      // 标题
	Content    string `json:"Content"`    // 发送内容
	Priority   int    `json:"Priority"`   // 发送优先级 0: 普通 1: 高
	EmailType  int    `json:"EmailType"`  // 邮件类型 0: 外部  1: 内部
	BodyFormat int    `json:"BodyFormat"` // 内容格式 0: 文本  1: html
	AppendPath string `json:"AppendPath"` // 附件路径
}

// Send 发送内网邮件, 实现 tof.Sender 接口
func (m *mail) Send(passID, token, env string) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return send(mailAPI, passID, token, env, data)
}

// NewMail 新建一个邮件发送器
func NewMail(opts ...mailOption) Sender {
	m := &mail{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// mailOption 邮件发送器选项
type mailOption func(m *mail)

// WithSender 设置发件人
func WithSender(addr string) mailOption {
	return func(m *mail) {
		m.From = addr
	}
}

// WithAddressee 设置收件人
func WithAddressee(addr string) mailOption {
	return func(m *mail) {
		m.To = addr
	}
}

// WithAddressee 设置收件人
func WithCC(addr string) mailOption {
	return func(m *mail) {
		m.CC = addr
	}
}

// WithTitle 设置邮件标题
func WithTitle(tittle string) mailOption {
	return func(m *mail) {
		m.Title = tittle
	}
}

// WithContent 设置邮件正文
func WithContent(content string) mailOption {
	return func(m *mail) {
		m.Content = content
	}
}

// WithBodyFormat 设置邮件正文格式(0: 文本  1: html)
func WithBodyFormat(format int) mailOption {
	return func(m *mail) {
		m.BodyFormat = format
	}
}
