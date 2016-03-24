package item

type Interactor interface {
	Get(string) interface{}
	Set(string, interface{}) error
}
