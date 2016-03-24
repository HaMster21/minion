package item

import "github.com/hamster21/minion/storage"

type Serializor interface {
	Save(storage.Backend) error
	Load(storage.Backend) error
}

