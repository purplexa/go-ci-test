package main

import (
	"testing"
)

func populateData() Data {
	users := []User{
		{
			name:     "Lynn Conway",
			gender:   "female",
			pronouns: "she/her",
			age:      84,
		},
		{
			name:     "Alan Turing",
			gender:   "male",
			pronouns: "he/him",
			age:      41,
		},
		{
			name:     "Chanda Prescod-Weinstein",
			gender:   "agender",
			pronouns: "she/her",
			age:      40,
		},
		{
			name:     "Katherine Johnson",
			gender:   "female",
			pronouns: "she/her",
			age:      101,
		},
		{
			name:     "Mark Dean",
			gender:   "male",
			pronouns: "he/him",
			age:      65,
		},
	}
	data := Data{
		users: users,
	}
	return data
}

func compareUser(u1 User, u2 User) bool {
	return u1.name == u2.name &&
		u1.gender == u2.gender &&
		u1.pronouns == u2.pronouns &&
		u1.age == u2.age
}

func findUserInArray(u1 User, list []User) bool {
	for u2 := range list {
		if compareUser(u1, list[u2]) {
			return true
		}
	}
	return false
}

func TestFindUserByName(t *testing.T) {
	data := populateData()
	users := data.findUserByName("Lynn Conway")
	if len(users) != 1 {
		t.Errorf("Found %d results, expected 1", len(users))
		return
	}
	got := users[0]
	expected := User{
		name:     "Lynn Conway",
		gender:   "female",
		pronouns: "she/her",
		age:      84,
	}
	if !compareUser(got, expected) {
		t.Errorf("got %v wanted %v", got, expected)
	}
}

func TestFindUserByGender(t *testing.T) {
	data := populateData()
	users := data.findUserByGender("female")
	if len(users) != 2 {
		t.Errorf("Found %d results, expected 2", len(users))
		return
	}
	expected1 := User{
		name:     "Lynn Conway",
		gender:   "female",
		pronouns: "she/her",
		age:      84,
	}
	expected2 := User{
		name:     "Katherine Johnson",
		gender:   "female",
		pronouns: "she/her",
		age:      101,
	}
	if !findUserInArray(expected1, users) || !findUserInArray(expected2, users) {
		t.Error("Missing expected user from array")
	}
}

func TestFindUserYoungerThan(t *testing.T) {
	data := populateData()
	users := data.findUserYoungerThan(30)
	if len(users) != 0 {
		t.Errorf("Found %d results under 30, expected 0", len(users))
		return
	}
	users = data.findUserYoungerThan(50)
	if len(users) != 2 {
		t.Errorf("Found %d results under 50, expected 2", len(users))
		return
	}
	users = data.findUserYoungerThan(80)
	if len(users) != 3 {
		t.Errorf("Found %d results under 80, expected 3", len(users))
		return
	}
}

func TestFindUserOlderThan(t *testing.T) {
	data := populateData()
	users := data.findUserOlderThan(30)
	if len(users) != 5 {
		t.Errorf("Found %d results over 30, expected 5", len(users))
		return
	}
	users = data.findUserOlderThan(50)
	if len(users) != 3 {
		t.Errorf("Found %d results over 50, expected 3", len(users))
		return
	}
	users = data.findUserOlderThan(80)
	if len(users) != 2 {
		t.Errorf("Found %d results over 80, expected 2", len(users))
		return
	}
}
