package embed_with_duplicate_name_test

import (
	"fmt"
	"testing"
)

type user struct {
	ID string
}

type admin struct {
	*user
	ID string
}

type superUser struct {
	*user
	ID string
}

func (u *user) GetID() string {
	return u.ID
}

func (a *admin) GetID() string {
	return a.ID
}

func TestUser(t *testing.T) {
	a := admin{
		user: &user{ID: "user1"},
		ID:   "admin1",
	}
	fmt.Println("a.ID()", a.GetID()) // "admin1"
	fmt.Println("a.ID", a.ID)        // "admin1"

	s := superUser{
		user: &user{ID: "user1"},
		ID:   "superuser1",
	}
	fmt.Println("s.GetID()", s.GetID()) // "user1"
	fmt.Println("s.ID", s.ID)           // "superuser1"
}
