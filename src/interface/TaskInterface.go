package _interface
import (
	"github.com/newbiet21379/learn_golang/struct"
)
type TaskInterface interface {
	FindShortestCombination( task _struct.Task ) int
	ShortestSubString( input string ) int
}