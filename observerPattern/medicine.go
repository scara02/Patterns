package main

import "fmt"

type Medicine struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newMedicine(name string) *Medicine {
	return &Medicine{
		name: name,
	}
}

func (m *Medicine) register(o Observer) {
	m.observerList = append(m.observerList, o)
}

func (m *Medicine) deregister(o Observer) {
	m.observerList = removeFromList(m.observerList, o)
}

func (m *Medicine) updateAvailability() {
	m.inStock = true
	fmt.Printf("Medicine %s is now in stock\n", m.name)
	m.notifyAll()
}

func (m *Medicine) notifyAll() {
	for _, observer := range m.observerList {
		observer.update(m.name)
	}
}

func removeFromList(observerList []Observer, o Observer) []Observer {
	length := len(observerList)
	for i, observer := range observerList {
		if o.getID() == observer.getID() {
			observerList[length-1], observerList[i] = observerList[i], observerList[length-1]
			return observerList[:length-1]
		}
	}
	return observerList
}
