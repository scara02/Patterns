package main

type Route interface {
	routeAdministration(m *Medication)
}
