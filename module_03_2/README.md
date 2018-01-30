
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

