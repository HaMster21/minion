package item

import "fmt"

type Item interface {
	fmt.Stringer
	Serializer
	Interactor
}
