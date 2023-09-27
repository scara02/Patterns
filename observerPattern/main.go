package main

func main() {
	medicine := newMedicine("Aspirin")

	client1 := &Client{id: "chuuya@bsd.jp"}
	medicine.register(client1)

	client2 := &Client{id: "dazai@bsd.jp"}
	medicine.register(client2)

	medicine.updateAvailability()

	medicine.deregister(client2)

	medicine.updateAvailability()
}
