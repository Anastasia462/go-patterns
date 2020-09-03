package builder

type Human struct {
	Parts string
}

func (h *Human) Born() string {
	return h.Parts
}
