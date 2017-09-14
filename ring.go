package logger

type Ring struct {
	ic chan string
	oc chan string
}

func NewRing(ic chan string, oc chan string) *Ring {
	return &Ring{ic, oc}
}

func (r *Ring) Run() {
	for v := range r.ic {
		select {
		case r.oc <- v:
		default:
			<-r.oc
			r.oc <- v
		}
	}
}

func (r *Ring) Produce(msg string) {
	r.ic <- msg
}
