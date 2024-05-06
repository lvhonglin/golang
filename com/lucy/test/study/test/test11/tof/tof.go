// Package tof 内网消息服务实现
package tof

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/xid"
)

const (
	tofOA  = "http://rio.tencent.com"        // 智能网关在 OA 区的接入点域名
	tofIdc = "http://idc.rio.tencent.com"    // 智能网关在 IDC 区的接入点域名
	tofDev = "http://devnet.rio.tencent.com" // 智能网关在 DevCloud 区的接入点域名
)

// Sender tof 消息发送接口
type Sender interface {
	// Send 根据 passid, token, env 发送 tof 消息
	Send(passID, token, env string) error
}

// Error 消息服务错误信息
type Error struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
	MsgID   string `json:"MsgId"`
}

// Error tof 错误信息 error 实现
func (t *Error) Error() string {
	return fmt.Sprintf("send email or message error, info: %v", t.ErrMsg)
}

// 生成 tof 签名
func sign(timestamp, nonce, token string) string {
	signStr := fmt.Sprintf("%s%s%s%s", timestamp, token, nonce, timestamp)
	sign := fmt.Sprintf("%X", sha256.Sum256([]byte(signStr)))
	return sign
}

// 生成 tof 请求头
func header(passid, token string) http.Header {
	h := http.Header{}
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	nonce := xid.New().String()
	sig := sign(timestamp, nonce, token)

	h.Add("x-rio-paasid", passid)
	h.Add("x-rio-nonce", nonce)
	h.Add("x-rio-timestamp", timestamp)
	h.Add("x-rio-signature", sig)
	h.Set("Content-Type", "application/json;charset=UTF-8")

	return h
}

// tof 标准发送逻辑
func send(api, passid, token, env string, data []byte) error {
	url := fmt.Sprintf("%v%v", tofDomain(env), api)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header = header(passid, token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	errInfo := &Error{}
	if err := json.Unmarshal(body, errInfo); err != nil {
		return err
	}
	if errInfo.ErrCode != 0 {
		return errInfo
	}
	return nil
}

func tofDomain(env string) string {
	tofDomain := map[string]string{
		"oa":  tofOA,
		"idc": tofIdc,
		"dev": tofDev,
	}
	return tofDomain[env]
}
