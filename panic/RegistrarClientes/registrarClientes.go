package main

import (
	"errors"
	"fmt"
)

type Client struct {
	File  int
	Name  string
	ID    int
	Phone string
	Home  string
}

var clients = []Client{
	{1, "Juan", 123, "123456", "Home 1"},
	{2, "Ana", 456, "234567", "Home 2"},
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Corrige los errores antes de continuar")
		}
		fmt.Println("End of execution")
		fmt.Println("Several errors were detected at runtime")
	}()

	newClient := Client{3, "Pedro", 789, "345678", "Home 3"}

	if err := registerClient(newClient); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func registerClient(newClient Client) error {
	for _, client := range clients {
		if client.ID == newClient.ID {
			return errors.New("Error: client already exists")
		}
	}

	_, err := validateNewClient(newClient)
	if err != nil {
		return err
	}

	clients = append(clients, newClient)

	return nil
}

func validateNewClient(newClient Client) (bool, error) {
	if newClient.File == 0 || newClient.Name == "" || newClient.ID == 0 || newClient.Phone == "" || newClient.Home == "" {
		return false, errors.New("Error: all client fields must have a non zero value")
	} else {
		return true, nil
	}
}
