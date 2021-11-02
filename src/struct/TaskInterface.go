package _struct
import (
	"Learn_Golang/src/learn"
)
type TaskInterface interface {
	FindShortestCombination( task learn.Task ) int
	ShortestSubString( input string ) int
}