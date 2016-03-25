package item

type Serializer interface {
	Save() error
	Load() error
}

