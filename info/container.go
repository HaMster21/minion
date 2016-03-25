package info

type Container interface {
	Get(domPath string) interface{}
	Set(domPath string, value interface{}) error

	Raw() map[string]interface{}
}
