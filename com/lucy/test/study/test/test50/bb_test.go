package test50

import (
	"github.com/agiledragon/gomonkey"
	"testing"
)

func Laugh() string {
	return "asdasdasda"
}
func Test_gomonkey(t *testing.T) {
	patch := gomonkey.ApplyFunc(Laugh, func() string {
		return "11111"
	})
	patch.Reset()
	t.Errorf("laguh:%s", Laugh())
}
