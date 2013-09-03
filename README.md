# boolmap
--
    import "github.com/MJKWoolnough/boolmap"

boolmap creates a map of bools using uint64s for efficiency (needs benchmarking for memory)

## Usage

#### type Map

```go
type Map struct {
}
```


#### func  NewMap

```go
func NewMap() Map
```

#### func (Map) Get

```go
func (m Map) Get(p uint64) bool
```

#### func (Map) Set

```go
func (m Map) Set(p uint64, d bool)
```
