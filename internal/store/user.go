package store

import (
	"maps"
	"slices"
	"sort"
	"strconv"
)

type UserStore struct {
	Users    map[string]User
	TopUsers []User
}

type User struct {
	Username  string
	PostCount int
}

func (us UserStore) GetUsersByPostCount() {
	users := slices.Collect(maps.Values(us.Users))
	sort.Slice(users, func(i, j int) bool {
		// Sort users by post count descending
		return users[i].PostCount > users[j].PostCount
	})
	us.TopUsers = users
}

func (us UserStore) FormatTopUsers() string {
	output := "Top users:\n"
	count := len(us.TopUsers)
	if count > 5 {
		count = 5
	}
	// Arbitrarily display the top 5 users
	for i := range count {
		user := us.TopUsers[i]
		output += user.Username + " with " + strconv.Itoa(user.PostCount) + " posts\n"
	}

	return output
}
