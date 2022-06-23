package observer

type Publisher[data any] interface {
	Subscribe(Listenner[data]) error
	Next(data)
}

type Listenner[D any] interface {
	ID() string
	Listen(D)
}
