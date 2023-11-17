package main

import (
	"encoding/base64"
	"strings"
)

func main() {
	decodeString, err := base64.StdEncoding.DecodeString("CTAAAAAAAIIAXAAAAAAAABjK8rq5DTpUL3RycGMuZV9jb21tZXJjZV9zYWxlLndlY2hhdF9jaGFubmVsc19hZGFwdGVyLkFkYXB0ZXJTZXJ2aWNlL1NldFByb2R1Y3REZWxpdmVyTWV0aG9kCgoxMjMxMjMxMjMxEggyMTIzMTIzMQ==")
	if err != nil {
		return
	}
	println(string(decodeString))
	a := string(decodeString)
	all := strings.ReplaceAll(a, "AdapterService", "WechatChannelsAdapter")
	toString := base64.StdEncoding.EncodeToString([]byte(all))
	println(toString)
}
