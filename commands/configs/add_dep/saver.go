package add_dep

type depSaver interface {
	add(depConfig) error
	read() ([]depConfig, error)
}
type defaultDepSaver struct {
}

func (d defaultDepSaver) add(config depConfig) error {
	panic("implement me")
}

func (d defaultDepSaver) read() ([]depConfig, error) {
	panic("implement me")
}
