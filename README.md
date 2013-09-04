# boolmap
--
    import "github.com/MJKWoolnough/boolmap"

boolmap creates a map of bools using bytes for efficiency (needs benchmarking for memory)

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

#### func (*Map) Get

```go
func (m *Map) Get(p uint) bool
```

#### func (*Map) Set

```go
func (m *Map) Set(p uint, d bool)
```

#### type Slice

```go
type Slice struct {
}
```


#### func  NewSlice

```go
func NewSlice() Slice
```

#### func (*Slice) Get

```go
func (s *Slice) Get(p uint) bool
```

#### func (*Slice) Set

```go
func (s *Slice) Set(p uint, d bool)
```
