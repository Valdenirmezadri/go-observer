package observer

import (
	"fmt"
)

type publisher[D any] struct {
	async       bool
	subscribers []Listenner[D]
}

func NewPublisher[D any](async bool) Publisher[D] {
	return &publisher[D]{
		async: async,
	}
}

func (p *publisher[D]) Subscribe(o Listenner[D]) error {
	if o == nil {
		return fmt.Errorf("listenner is nil")
	}

	p.subscribers = append(p.subscribers, o)

	return nil
}

func (p *publisher[D]) Next(data D) {
	for _, s := range p.subscribers {
		if p.async {
			go s.Listen(data)
		} else {
			s.Listen(data)
		}
	}
}

func (p *publisher[D]) UnSubscribe(o Listenner[D]) {
	p.subscribers = p.removeFromslice(p.subscribers, o)
}

func (p *publisher[D]) removeFromslice(observerList []Listenner[D], observerToRemove Listenner[D]) []Listenner[D] {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.ID() == observer.ID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
