package core

type EasyCarHttpHandler struct {
	rm TMInterface
}

func NewEasyCarHttpHandler(rm TMInterface) EasyCarHttpHandler {
	return EasyCarHttpHandler{rm: rm}
}

func (http *EasyCarHttpHandler) run() {
	http.run()
}
