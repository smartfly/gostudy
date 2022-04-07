package enum

//go:generate stringer -type=FishType
type FishType int

const (
	A FishType = iota
	B
	C
	D
)

//func (f FishType) String() string {
//	return [...]string{"A", "B", "C", "D"}[f]
//}
