package main

import "fmt"

var currentID int
var users Users

func init() {
	RepoCreateUser(User{
		Name:   "jfray",
		Mobile: "925-111-2222",
		Email:  "jfray@example.tld",
	})
	RepoCreateUser(User{
		Name:   "bff",
		Mobile: "510-111-2222",
		Email:  "bff@example.tld",
	})
}

func RepoFindUser(id int) User {
	for _, u := range users {
		if u.ID == id {
			return u
		}
	}
	// return blank if not found
	return User{}
}

func RepoCreateUser(u User) User {
	currentID += 1
	u.ID = currentID
	users = append(users, u)
	return u
}

func RepoDestroyUser(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}
