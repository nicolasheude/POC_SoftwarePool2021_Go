package humanity

type Preparer interface {
	Prepare() error
}

func (h *Human) Prepare() error {
	h.Ready = true
	return nil
}
