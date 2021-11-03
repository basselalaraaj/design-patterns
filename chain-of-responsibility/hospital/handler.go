package hospital

type handler interface {
	handle(*patient)
	setNext(handler)
}

type nextHandler struct {
	next handler
}

func (h *nextHandler) setNext(handler handler) {
	h.next = handler
}

func (h *nextHandler) handleNext(p *patient) {
	if h.next == nil {
		return
	}

	h.next.handle(p)
}
