package esa

import (
	"time"
)

type Team struct {
	Name        string `json:"name"`
	Privacy     string `json:"privacy"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Url         string `json:"url"`
}

type Teams struct {
	Teams      []Team `json:"teams"`
	PrevPage   string `json:"prev_page"`
	NextPage   string `json:"next_page"`
	TotalCount int    `json:"total_count"`
}

type Stats struct {
	Members            int `json:"members"`
	Posts              int `json:"posts"`
	Comments           int `json:"comments"`
	Stars              int `json:"stars"`
	DailyActiveUsers   int `json:"daily_active_users"`
	WeeklyActiveUsers  int `json:"weekly_active_users"`
	MonthlyActiveUsers int `json:"monthly_active_users"`
}

type Member struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
	Email      string `json:"email"`
}

type Members struct {
	Members    []Member `json:"members"`
	PrevPage   string   `json:"prev_page"`
	NextPage   string   `json:"next_page"`
	TotalCount int      `json:"total_count"`
}

type ByUser struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
}

type Posts struct {
	Posts      []Post `json:"posts"`
	PrevPage   string `json:"prev_page"`
	NextPage   string `json:"next_page"`
	TotalCount int    `json:"total_count"`
}

type Post struct {
	Number          int       `json:"number"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Wip             bool      `json:"wip"`
	BodyMd          string    `json:"body_md"`
	BodyHtml        string    `json:"body_html"`
	CreatedAt       time.Time `json:"created_at"`
	Message         string    `json:"message"`
	Url             string    `json:"url"`
	UpdatedAt       time.Time `json:"updated_at"`
	Tags            []string  `json:"tags"`
	Category        string    `json:"category"`
	RevisionNumber  int       `json:"revision_number"`
	CreatedBy       ByUser    `json:"created_by"`
	UpdatedBy       ByUser    `json:"updated_by"`
	Kind            string    `json:"kind"`
	CommentsCount   int       `json:"comments_count"`
	TasksCount      int       `json:"tasks_count"`
	DoneTasksCount  int       `json:"done_tasks_count"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Star            bool      `json:"star"`
	Watch           bool      `json:"watch"`
}

type Comment struct {
	Id        int       `json:"id"`
	BodyMd    string    `json:"body_md"`
	BodyHtml  string    `json:"body_html"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	CreatedBy ByUser    `json:"created_by"`
}

type Comments struct {
	Comments   []Comment `json:"comments"`
	PrevPage   string    `json:"prev_page"`
	NextPage   string    `json:"next_page"`
	TotalCount int       `json:"total_count"`
}
