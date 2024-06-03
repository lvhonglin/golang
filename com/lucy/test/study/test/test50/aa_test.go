package test50

import (
	"github.com/golang/mock/gomock"
	"golang/com/lucy/test/study/test/test50/mock"
	"testing"
)

func TestA(t *testing.T) {
	controller := gomock.NewController(t)
	human := mock.NewMockMockHuman(controller)
	human.EXPECT().Speak(gomock.Any()).DoAndReturn(func(input string) string {
		println("1123123")
		println(input)

		return "aaaa"
	})
	t.Logf(human.Speak("1223"))
	t.Run("单测1", func(t *testing.T) {
		println("顶顶顶顶")
		t.Errorf("错了")
	})
	controller.Finish()
}
