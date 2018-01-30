
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

## Create your own package

Create a workspace for your package

```
mkdir $GOPATH/src/module_003
cd $GOPATH/src/module_003
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
import "module_003/mytool"

func main() {
	fmt.Println(mytool.Double(5))
}
```

```
go run main.go
```

## Download package with `dep`

Create workspace
```
cd $GOPATH/src
mkdir module_003_2
```

```
dep init
ls .
```

Create main.go using an external package
```
vim main.go
```

```
package main

import "fmt"
import "github.com/gosimple/slug"

func main() {
	fmt.Println(slug.Make("I wish I can fly."))
}
```

Try running main.go (should fail)
```
go run main.go
```

`dep ensure` will install external packages used in source

```
dep ensure
ls -R vendor
```

Try running main.go (should run)
```
go run main.go
```

