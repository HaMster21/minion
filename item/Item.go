package item

type Item interface {
	Name() string
	Position() int
	Description() string
}
