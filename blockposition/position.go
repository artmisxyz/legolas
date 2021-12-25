package blockposition

type BlockPositionHolder interface {
	Update(uint64) error
	Create() error
	Exists() bool
	Read() (uint64, error)
	Incr() error
}
