# Hello World

### Go environment check
Let's test if our go environment is correctly setup

    $> go version
    go version go1.7 linux/amd64

    $> echo $GOPATH
    /home/vds/go

    $> echo $GOROOT
    /home/vds/local/go


### Setup our first package
Let's set up our go package

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



