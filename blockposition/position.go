package blockposition

type BlockPositionHolder interface {
	Update(uint64) error
	Create() error
	Exists() bool
	Read(point Point) (uint64, error)
	Incr() error
}

type Point int

const (
	Start = iota
	Current
	Finish
)
