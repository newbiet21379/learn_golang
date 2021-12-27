package entity

type Message struct {
	Str  string
	Wait chan bool
}
