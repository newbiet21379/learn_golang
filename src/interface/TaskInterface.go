package _interface
import (
	"Learn_Golang/src/struct"
)
type TaskInterface interface {
	FindShortestCombination( task _struct.Task ) int
	ShortestSubString( input string ) int
}