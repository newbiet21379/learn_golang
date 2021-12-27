
To declare an “async” function in Go:

- The return type is `<-chan ReturnType`.

- Within the function, create a channel by `make(chan ReturnType)` and return the created channel at the end of the function.

- Start an anonymous goroutine by `go func() {...}` and implement the function’s logic inside that anonymous function.

- Return the result by sending the value to the channel.

- At the beginning of the anonymous function, add defer close(r) to close the channel once done.

To “await” the result, simply read the value from the channel by `v := <- fn()`.