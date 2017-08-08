## Golang Initial Setup

[Read](https://golang.org/doc/code.html) or [watch](https://www.youtube.com/watch?v=XCsL89YtqCs) Go basics to understand Go workspace and find out how to import other packages. Then continue with [Effective Go](https://golang.org/doc/effective_go.html).

Install Go using the official installation package or on MacOS using Homebrew `brew install go`.

`GOPATH` is set by default to `$HOME/go`.

VS Code works great with Go, just make sure you have the [Go extension](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go) installed.

### Explore

- [Documentation Index](http://godoc.org)
- [Projects by category](https://github.com/golang/go/wiki/Projects)
- [Online Playground](https://play.golang.org)

### Data allocation

- [Effective go](https://golang.org/doc/effective_go.html#data)
- Use `make` for slices, maps, and channels. It initialize the memory.
- Use `new` or composite literals for everything else. It zeores the memory.

### Routers

- [net/http](https://golang.org/pkg/net/http/) ([wiki](https://golang.org/doc/articles/wiki/))
- [gorilla/mux](http://www.gorillatoolkit.org)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)

### Live reload

- [codegangsta/gin](https://github.com/codegangsta/gin)

### SQL

- [database/sql](https://golang.org/pkg/database/sql/) ([wiki](https://github.com/golang/go/wiki/SQLInterface))

### Mongo

- [mgo driver](https://labix.org/mgo)