
## Download package with `go get`

```
go get -u github.com/golang/dep
```

Check what is installed
```
ls $GOPATH/bin
ls $GOPATH/pkg
ls $GOPATH/src/github.com/golang/dep
```

Try using it
```
dep help
```

## Create your own package

Create a workspace for your package

```
mkdir $GOPATH/src/module_03_1
cd $GOPATH/src/module_03_1
mkdir mytool
cd mytool
vim main.go
```

Write your package `mytool`

```
package mytool

func Double(x int) int {
	return x * 2
}
```

Install `mytool`

```
cd ..
go install -v ./mytool
```

Check what is installed

```
ls $GOPATH/pkg
```

## Use your package

```
vim main.go
```

```
package main

import "fmt"
import "module_03_1/mytool"

func main() {
	fmt.Println(mytool.Double(5))
}
```

Try running main.go (should output `10` )
```
go run main.go
```

