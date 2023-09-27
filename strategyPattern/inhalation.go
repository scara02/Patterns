package main

import "fmt"

type Inhalation struct {
}

func (i *Inhalation) routeAdministration(m *Medication) {
	fmt.Printf("Route administration of %s is inhalation!!!\n", m.name)
}
