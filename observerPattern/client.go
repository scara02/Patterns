package main

import "fmt"

type Client struct {
	id string
}

func (c *Client) getID() string {
	return c.id
}

func (c *Client) update(medicine string) {
	fmt.Printf("Sending notification to client %s for medicine restock %s\n", c.id, medicine)
}
