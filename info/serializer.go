package info

type Storage interface {
	Save(Container) error
	Load() (Container,error)
}

