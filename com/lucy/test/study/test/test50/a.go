package test50

//go:generate mockgen -destination=./mock/mock_human.go -package=mock -source=a.go
type MockHuman interface {
	Speak(cur string) string
	Walk() string
}

type MockHuman2 interface {
	Speak(cur string) string
	Walk() string
}
