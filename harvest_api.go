package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type User struct {
	ID                           int         `json:"id"`
	FirstName                    string      `json:"first_name"`
	LastName                     string      `json:"last_name"`
	Email                        string      `json:"email"`
	Telephone                    string      `json:"telephone"`
	Timezone                     string      `json:"timezone"`
	WeeklyCapacity               int         `json:"weekly_capacity"`
	HasAccessToAllFutureProjects bool        `json:"has_access_to_all_future_projects"`
	IsContractor                 bool        `json:"is_contractor"`
	IsAdmin                      bool        `json:"is_admin"`
	IsProjectManager             bool        `json:"is_project_manager"`
	CanSeeRates                  bool        `json:"can_see_rates"`
	CanCreateProjects            bool        `json:"can_create_projects"`
	CanCreateInvoices            bool        `json:"can_create_invoices"`
	CanCloseAccount              bool        `json:"can_close_account"`
	IsActive                     bool        `json:"is_active"`
	CalendarIntegrationEnabled   bool        `json:"calendar_integration_enabled"`
	CalendarIntegrationSource    interface{} `json:"calendar_integration_source"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	DefaultHourlyRate            interface{} `json:"default_hourly_rate"`
	CostRate                     float64     `json:"cost_rate"`
	Roles                        []string    `json:"roles"`
	PermissionsClaims            []string    `json:"permissions_claims"`
	AvatarURL                    string      `json:"avatar_url"`
}

type UserTimeReport struct {
	UserID         int     `json:"user_id"`
	UserName       string  `json:"user_name"`
	IsContractor   bool    `json:"is_contractor"`
	TotalHours     float64 `json:"total_hours"`
	BillableHours  float64 `json:"billable_hours"`
	Currency       string  `json:"currency"`
	BillableAmount float64 `json:"billable_amount"`
}

type ArrayUserTimeReport struct {
	Results []UserTimeReport `json:"results"`
}

type ArrayUser struct {
	Users []User `json:"users"`
}

func getHarvestAPIresponse(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Go Harvest API Sample")
	req.Header.Set("Harvest-Account-ID", os.Getenv("HarvestAccountID"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", resp.StatusCode)
	}

	defer resp.Body.Close()
	return body
}

func getHarvestTeamTimeReport(startDate string, endDate string) ArrayUserTimeReport {
	url := "https://api.harvestapp.com/v2/reports/time/team?from=" + startDate + "&to=" + endDate
	body := getHarvestAPIresponse(url)
	var jsonResponse ArrayUserTimeReport
	json.Unmarshal(body, &jsonResponse)
	return jsonResponse
}

func getHarvestActiveUsers() ArrayUser {
	rows := 100
	var listSlice ArrayUser
	for i := 1; rows == 100; i++ {
		results := getHarvestActiveUsersPage(i)
		rows = len(results.Users)
		listSlice.Users = append(listSlice.Users, results.Users...)
	}
	return listSlice
}

func getHarvestActiveUsersPage(page int) ArrayUser {
	url := "https://api.harvestapp.com/v2/users?is_active=true&page=" + strconv.Itoa(page)
	body := getHarvestAPIresponse(url)
	var jsonResponse ArrayUser
	err := json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsonResponse
}
