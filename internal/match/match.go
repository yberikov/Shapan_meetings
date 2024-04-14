package match

import (
	"fmt"
	"hacknu/internal/models"
	"sync"
	"time"
)

type UserPool struct {
	Mutex sync.Mutex
	Users []models.User
}

func (up *UserPool) AddUser(user models.User) {
	up.Mutex.Lock()
	defer up.Mutex.Unlock()
	dateTime := user.Date
	dateWithoutTime := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, dateTime.Location())
	user.Date = dateWithoutTime
	for _, u := range up.Users {
		if u.Email == user.Email && u.Date == user.Date {
			fmt.Println("Already exist", user)
			return
		}
	}

	up.Users = append(up.Users, user)
	fmt.Println("User add:", user)
}

func (up *UserPool) GetAllUsers() []models.User {
	up.Mutex.Lock()
	defer up.Mutex.Unlock()

	return up.Users
}

func (up *UserPool) FindMatch() (models.User, models.User, error) {
	users := up.GetAllUsers()
	fmt.Println(users)
	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			fmt.Println("Try first user:", users[i])
			fmt.Println("Try first second:", users[j])
			if usersMatch(users[i], users[j]) && timeIntervalsIntersect(users[i], users[j]) {
				user1 := users[i]
				user2 := users[j]
				// Remove user1 and user2 from the slice
				users = append(users[:i], users[i+1:]...)
				users = append(users[:j-1], users[j:]...)
				fmt.Printf("Match found: %s and %s\n", user1.Email, user2.Email)
				return user1, user2, nil
			}
			// If users don't match or their time intervals don't intersect, keep them

		}
	}
	fmt.Println("Match not found")
	return models.User{}, models.User{}, nil
}

// Check if two users match based on their data
func usersMatch(user1, user2 models.User) bool {
	return user1.Date == user2.Date && user1.Email != user2.Email // Adjust conditions as needed
}

// Check if the time intervals of two users intersect
func timeIntervalsIntersect(user1, user2 models.User) bool {
	from1, _ := time.Parse(time.RFC3339, user1.From)
	to1, _ := time.Parse(time.RFC3339, user1.To)
	from2, _ := time.Parse(time.RFC3339, user2.From)
	to2, _ := time.Parse(time.RFC3339, user2.To)

	return !(to1.Before(from2) || to2.Before(from1))
}

// a1 a2 b1  b2
