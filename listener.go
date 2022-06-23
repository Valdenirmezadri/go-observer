package observer

type listen[D any] struct {
	_ID string
	f   func(D)
}

func NewListener[D any](ID string, f func(D)) Listenner[D] {
	return &listen[D]{
		_ID: ID,
		f:   f,
	}
}

func (l *listen[D]) Listen(data D) {
	l.f(data)
}

func (l *listen[D]) ID() string {
	return l._ID
}
