package main

func main() {
	inhalation := &Inhalation{}

	med1 := initMedication("Morphine", inhalation)
	med1.administerRoute()

	ingestion := &Ingestion{}

	med1.setRoute(ingestion)
	med1.administerRoute()

	injection := &Injection{}

	med1.setRoute(injection)
	med1.administerRoute()
}
