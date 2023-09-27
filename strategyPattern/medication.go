package main

type Medication struct {
	name  string
	route Route
}

func initMedication(name string, r Route) *Medication {
	return &Medication{
		name:  name,
		route: r,
	}
}

func (m *Medication) setRoute(r Route) {
	m.route = r
}

func (m *Medication) administerRoute() {
	m.route.routeAdministration(m)
}
