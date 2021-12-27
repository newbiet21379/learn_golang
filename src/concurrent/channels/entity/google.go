package entity

type Search func(query string) Result
type Result struct {
	Data string
}
