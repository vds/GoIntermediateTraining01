# Hello World
we'll learn the very basics of go, how to test the go environment, how to create the first package and to run our first applicatio.

### Go environment check
Let's test if our go environment is correctly setup:

    $> go version
    go version goXXXXX

    $> echo $GOPATH
    /home/XXXXX/go
    /Home/XXXXX (MSWindows)
    
    $> echo $GOROOT
    /home/XXXX/local/go (*nix)
    /Go (MSWindows)

### Setup our first package
Let's set up our go package:

    $> mkdir $GOPATH/src/helloworld

    $> cd $GOPATH/src/helloworld

Open $GOPATH/src/helloworld/helloword.go

### Hello World

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World")
}
```

### Test it

    $> go run helloworld.go
    Hello, World


### Useful Links
A Tour of Go: https://tour.golang.org/welcome/1