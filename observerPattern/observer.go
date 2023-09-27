package main

type Observer interface {
	getID() string
	update(string)
}
