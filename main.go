package main

import (
	"fmt"
	"time"
)

func main() {

	timeReport := getHarvestTeamTimeReport("2021-01-01", "2021-12-31")
	activeUsers := getHarvestActiveUsers()

	fmt.Println("len:TimeReports", len(timeReport.Results))
	fmt.Println("len:Users", len(activeUsers.Users))

	sum := 0.0

	activeUsersIDs := activeUsers.getActiveUsersIDs(2021, 1, 1)

	for _, user := range activeUsers.Users {
		if user.CreatedAt.Before(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)) {
			activeUsersIDs = append(activeUsersIDs, user.ID)
		}
	}

	count := 0
	for _, user := range timeReport.Results {
		if contains(activeUsersIDs, user.UserID) {
			if user.BillableHours > 700 {
				sum += user.BillableHours
				count++
			}
		}
	}
	average := sum / float64(count)
	fmt.Println("average:", average)
	fmt.Println("sum:", sum)
	fmt.Println("count:", count)
}

func (activeUsers ArrayUser) getActiveUsersIDs(year int, month int, day int) []int {
	activeUsersIDs := make([]int, 0)
	for _, user := range activeUsers.Users {
		if user.CreatedAt.Before(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)) {
			activeUsersIDs = append(activeUsersIDs, user.ID)
		}
	}
	return activeUsersIDs
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
