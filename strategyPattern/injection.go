package main

import "fmt"

type Injection struct {
}

func (i *Ingestion) routeAdministration(m *Medication) {
	fmt.Printf("Route administration of %s is injection!!!\n", m.name)
}
