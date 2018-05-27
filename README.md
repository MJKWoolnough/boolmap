# boolmap
--
    import "github.com/MJKWoolnough/boolmap"

Package boolmap creates a map of bools using bytes for efficiency (needs
benchmarking for memory)

## Usage

#### type CrumbMap

```go
type CrumbMap map[uint64]byte
```

CrumbMap is a map of Crumbs (2-bits, values 0, 1, 2, 3)

#### func  NewCrumbMap

```go
func NewCrumbMap() CrumbMap
```
NewCrumbMap returns a new, initialised, CrumbMap

#### func (CrumbMap) Get

```go
func (c CrumbMap) Get(p uint64) byte
```
Get returns a crumb from the given position

#### func (CrumbMap) Set

```go
func (c CrumbMap) Set(p uint64, d byte)
```
Set sets the crumb at the given position

#### type CrumbSlice

```go
type CrumbSlice []byte
```

CrumbSlice is a slice of bytes, representing crumbs (2-bits)

#### func  NewCrumbSlice

```go
func NewCrumbSlice() *CrumbSlice
```
NewCrumbSlice returns a new, initialised, CrumbSlice

#### func  NewCrumbSliceSize

```go
func NewCrumbSliceSize(size uint) *CrumbSlice
```
NewCrumbSliceSize returns a new Crumbslice, initialised to the given size

#### func (CrumbSlice) Get

```go
func (c CrumbSlice) Get(p uint) byte
```
Get returns a crumb from the given position

#### func (*CrumbSlice) Set

```go
func (c *CrumbSlice) Set(p uint, d byte)
```
Set sets the crumb at the given position

#### type Map

```go
type Map map[uint64]byte
```

Map is the default boolmap

#### func  NewMap

```go
func NewMap() Map
```
NewMap returns a new, initialised Map

#### func (Map) Get

```go
func (m Map) Get(p uint64) byte
```
Get returns a bool, represented by a byte, for the specified position

#### func (Map) GetBool

```go
func (m Map) GetBool(p uint64) bool
```
GetBool returns a bool for the specified position

#### func (Map) Set

```go
func (m Map) Set(p uint64, d byte)
```
Set sets a bool, represented by a byte, at the specified position

#### func (Map) SetBool

```go
func (m Map) SetBool(p uint64, d bool)
```
SetBool sets a bool at the specified position

#### type NibbleMap

```go
type NibbleMap map[uint64]byte
```

NibbleMap is a map of Nibbles (4-bits, values 0-15)

#### func  NewNibbleMap

```go
func NewNibbleMap() NibbleMap
```
NewNibbleMap return a new, initialised, NibbleMap

#### func (NibbleMap) Get

```go
func (n NibbleMap) Get(p uint64) byte
```
Get returns a crumb from the given position

#### func (NibbleMap) Set

```go
func (n NibbleMap) Set(p uint64, d byte)
```
Set sets the crumb at the given position

#### type NibbleSlice

```go
type NibbleSlice []byte
```

NibbleSlice is a slice of bytes representing nibbles (4-bits)

#### func  NewNibbleSlice

```go
func NewNibbleSlice() NibbleSlice
```
NewNibbleSlice returns a new, initialised, CrumbSlice

#### func  NewNibbleSliceSize

```go
func NewNibbleSliceSize(size uint) NibbleSlice
```
NewNibbleSliceSize returns a new NibbleSlice, initialised to the given size

#### func (NibbleSlice) Get

```go
func (n NibbleSlice) Get(p uint) byte
```
Get returns a crumb from the given position

#### func (*NibbleSlice) Set

```go
func (n *NibbleSlice) Set(p uint, d byte)
```
Set sets the crumb at the given position

#### type Slice

```go
type Slice []byte
```

Slice is a slice of bytes representing bools

#### func  NewSlice

```go
func NewSlice() *Slice
```
NewSlice returnns a new, initialised Slice

#### func  NewSliceSize

```go
func NewSliceSize(size uint) *Slice
```
NewSliceSize returns a new Slice, intitialised to the size given

#### func (Slice) Get

```go
func (s Slice) Get(p uint) byte
```
Get returns a byte, representing a bool, at the specified position

#### func (Slice) GetBool

```go
func (s Slice) GetBool(p uint) bool
```
GetBool returns a bool for the specified position

#### func (*Slice) Set

```go
func (s *Slice) Set(p uint, d byte)
```
Set sets a bool, given as a byte, at the specified position

#### func (*Slice) SetBool

```go
func (s *Slice) SetBool(p uint, d bool)
```
SetBool sets a bool at the specified position
