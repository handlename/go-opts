# go-opts

Shortcut for package "flag".

## Usage

code

```go
package main

import (
    "fmt"

    "github.com/handlename/go-opts"
)

type myOpts struct {
    Name string `flag:"name" default:"alice" description:"your name"`
    Age  int    `flag:"age" default:"25" description:"your age"`
}

func main() {
    o := myOpts{}
    opts.Parse(&o)
    fmt.Println(o)
}
```

run

```
$ go run sample.go --name=bob
{bob 25}

$ go run sample/app.go --help
Usage of /var/folders/7r/p77s6hs55s1c7ch482_2m75w0000gn/T/go-build994360613/command-line-arguments/_obj/exe/app:
  -age=25: your age
  -name="alice": your name
exit status 2
```

## Todo

Test...

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[handlename](https://github.com/handlename)
