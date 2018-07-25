package esa

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hiroakis/esa-go/request"
)

// All of the dummy data are from https://docs.esa.io/posts/102

var teamsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	teams := `
{
  "teams": [
    {
      "name": "docs",
      "privacy": "open",
      "description": "esa.io official documents",
      "icon": "https://img.esa.io/uploads/production/teams/105/icon/thumb_m_0537ab827c4b0c18b60af6cdd94f239c.png",
      "url": "https://docs.esa.io/"
    }
  ],
  "prev_page": null,
  "next_page": 1,
  "total_count": 1
}
`

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(teams))
})

var teamHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	team := `
{
  "name": "docs",
  "privacy": "open",
  "description": "esa.io official documents",
  "icon": "https://img.esa.io/uploads/production/teams/105/icon/thumb_m_0537ab827c4b0c18b60af6cdd94f239c.png",
  "url": "https://docs.esa.io/"
}
`

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(team))
})

var statsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	stats := `
{
  "members": 20,
  "posts": 1959,
  "comments": 2695,
  "stars": 3115,
  "daily_active_users": 8,
  "weekly_active_users": 14,
  "monthly_active_users": 15
}
`

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(stats))
})

var membersHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	members := `
{
  "members": [
    {
      "name": "Hiroaki Sano",
      "screen_name": "hiroakis",
      "icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png",
      "email": "hiroakis@example.com"
    },
    {
      "name": "Sano Hiroaki",
      "screen_name": "sano",
      "icon": "https://img.esa.io/uploads/production/users/2/icon/thumb_m_2690997f07b7de3014a36d90827603d6.jpg",
      "email": "sano@example.com"
    }
  ],
  "prev_page": null,
  "next_page": 1,
  "total_count": 2
}
`

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(members))
})

var postsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	posts := `
{
  "posts": [
    {
      "number": 1,
      "name": "hi!",
      "full_name": "日報/2015/05/09/hi! #api #dev",
      "wip": true,
      "body_md": "# Getting Started",
      "body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
      "created_at": "2015-05-09T11:54:50+09:00",
      "message": "Add Getting Started section",
      "url": "https://docs.esa.io/posts/1",
      "updated_at": "2015-05-09T11:54:51+09:00",
      "tags": [
        "api",
        "dev"
      ],
      "category": "日報/2015/05/09",
      "revision_number": 1,
      "created_by": {
        "name": "Hiroaki Sano",
        "screen_name": "hiroakis",
        "icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
      },
      "updated_by": {
        "name": "Hiroaki Sano",
        "screen_name": "hiroakis",
        "icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
      }
    }
  ],
  "prev_page": null,
  "next_page": 1,
  "total_count": 1
}
`

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(posts))
})

var postHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := `
{
  "number": 1,
  "name": "hi!",
  "full_name": "日報/2015/05/09/hi! #api #dev",
  "wip": true,
  "body_md": "# Getting Started",
  "body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
  "created_at": "2015-05-09T11:54:50+09:00",
  "message": "Add Getting Started section",
  "url": "https://docs.esa.io/posts/1",
  "updated_at": "2015-05-09T11:54:51+09:00",
  "tags": [
    "api",
    "dev"
  ],
  "category": "日報/2015/05/09",
  "revision_number": 1,
  "created_by": {
    "name": "Hiroaki Sano",
    "screen_name": "hiroakis",
    "icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
  },
  "updated_by": {
    "name": "Hiroaki Sano",
    "screen_name": "hiroakis",
    "icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
  },
  "kind": "flow",
  "comments_count": 1,
  "tasks_count": 1,
  "done_tasks_count": 1,
  "stargazers_count": 1,
  "watchers_count": 1,
  "star": true,
  "watch": true
}
`

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(post))
})

var createPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := `
{
  "number": 1,
  "name": "hi!",
  "full_name": "日報/2015/05/09/hi! #api #dev",
  "wip": true,
  "body_md": "# Getting Started",
  "body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
  "created_at": "2015-05-09T11:54:50+09:00",
  "message": "Add Getting Started section",
  "url": "https://docs.esa.io/posts/1",
  "updated_at": "2015-05-09T11:54:51+09:00",
  "tags": [
    "api",
    "dev"
  ],
  "category": "日報/2015/05/09",
  "revision_number": 1,
  "created_by": {
    "name": "Hiroaki Sano",
    "screen_name": "hiroakis",
    "icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
  },
  "updated_by": {
    "name": "Hiroaki Sano",
    "screen_name": "hiroakis",
    "icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
  },
  "kind": "flow",
  "comments_count": 1,
  "tasks_count": 1,
  "done_tasks_count": 1,
  "stargazers_count": 1,
  "watchers_count": 1,
  "star": true,
  "watch": true
}
`
	var postData request.PostData
	bufbody := &bytes.Buffer{}
	bufbody.ReadFrom(r.Body)
	json.Unmarshal(bufbody.Bytes(), &postData)

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Name != "hi!" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.BodyMd != "# Getting Started\n" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Tags[0] != "api" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Tags[1] != "dev" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Category != "dev/2015/05/10" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Wip != false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Message != "Add Getting Started section" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(post))
})

var updatePostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := `
{
  "number": 1,
  "name": "hi!",
  "full_name": "日報/2015/05/09/hi! #api #dev",
  "wip": true,
  "body_md": "# Getting Started",
  "body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
  "created_at": "2015-05-09T11:54:50+09:00",
  "message": "Add Getting Started section",
  "url": "https://docs.esa.io/posts/1",
  "updated_at": "2015-05-09T11:54:51+09:00",
  "tags": [
    "api",
    "dev"
  ],
  "category": "日報/2015/05/09",
  "revision_number": 1,
  "created_by": {
    "name": "Hiroaki Sano",
    "screen_name": "hiroakis",
    "icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
  },
  "updated_by": {
    "name": "Hiroaki Sano",
    "screen_name": "hiroakis",
    "icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
  },
  "overlapped": false,
  "kind": "flow",
  "comments_count": 1,
  "tasks_count": 1,
  "done_tasks_count": 1,
  "stargazers_count": 1,
  "watchers_count": 1,
  "star": true,
  "watch": true
}
`
	var postData request.PostData
	bufbody := &bytes.Buffer{}
	bufbody.ReadFrom(r.Body)
	json.Unmarshal(bufbody.Bytes(), &postData)

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Name != "hi!" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.BodyMd != "# Getting Started\n" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Tags[0] != "api" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Tags[1] != "dev" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Category != "dev/2015/05/10" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Wip != false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.Message != "Add Getting Started section" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.OriginalRevision.BodyMd != "# Getting ..." {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.OriginalRevision.Number != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if postData.Post.OriginalRevision.User != "hiroakis" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(post))
})

var deletePostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	// w.Write([]byte(post))
})

var commentsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	comments := `
{
  "comments": [
    {
      "id": 1,
      "body_md": "(大事)",
      "body_html": "<p>(大事)</p>",
      "created_at": "2014-05-10T12:45:42+09:00",
      "updated_at": "2014-05-18T23:02:29+09:00",
      "url": "https://docs.esa.io/posts/2#comment-1",
      "created_by": {
        "name": "Hiroaki Sano",
        "screen_name": "hiroakis",
        "icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
      }
    }
  ],
  "prev_page": null,
  "next_page": 1,
  "total_count": 1
}
`

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(comments))
})

var commentHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	comment := `
{
  "id": 13,
  "body_md": "読みたい",
  "body_html": "<p>読みたい</p>",
  "created_at": "2014-05-13T16:17:42+09:00",
  "updated_at": "2014-05-18T23:02:29+09:00",
  "url": "https://docs.esa.io/posts/13#comment-13",
  "created_by": {
    "name": "Sano Hiroaki",
    "screen_name": "sano",
    "icon": "https://img.esa.io/uploads/production/users/2/icon/thumb_m_2690997f07b7de3014a36d90827603d6.jpg"
  }
}
`

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(comment))
})

var createCommentHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	comment := `
{
  "id": 22767,
  "body_md": "LGTM!",
  "body_html": "<p>LGTM!</p>\n",
  "created_at": "2015-06-21T19:36:20+09:00",
  "updated_at": "2015-06-21T19:36:20+09:00",
  "url": "https://docs.esa.io/posts/2#comment-22767",
  "created_by": {
    "name": "Hiroaki Sano",
    "screen_name": "hiroakis",
    "icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
  }
}
`
	var commentData request.CommentData
	bufbody := &bytes.Buffer{}
	bufbody.ReadFrom(r.Body)
	json.Unmarshal(bufbody.Bytes(), &commentData)

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if commentData.Comment.BodyMd != "LGTM!" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(comment))
})

var updateCommentHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	comment := `
{
  "id": 22767,
  "body_md": "LGTM!!!",
  "body_html": "<p>LGTM!!!</p>\n",
  "created_at": "2015-06-21T19:36:20+09:00",
  "updated_at": "2015-06-21T19:40:33+09:00",
  "url": "https://docs.esa.io/posts/2#comment-22767",
  "created_by": {
    "name": "Hiroaki Sano",
    "screen_name": "hiroakis",
    "icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
  }
}
`
	var commentData request.CommentData
	bufbody := &bytes.Buffer{}
	bufbody.ReadFrom(r.Body)
	json.Unmarshal(bufbody.Bytes(), &commentData)

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if commentData.Comment.BodyMd != "LGTM!!!" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(comment))
})

var deleteCommentHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	// w.Write([]byte(post))
})

func fakeClient(testURL string) *EsaClient {
	esaClient := NewEsaClient("accessToken", "team")
	esaClient.SetApi(testURL)
	return esaClient
}

func TestGetTeams(t *testing.T) {

	testServer := httptest.NewServer(teamsHandler)
	defer testServer.Close()
	teams, err := fakeClient(testServer.URL).GetTeams()

	if err != nil {
		t.Error("Error occurred")
	}
	if teams.Teams[0].Name != "docs" {
		t.Error("Name does not match")
	}
	if teams.Teams[0].Privacy != "open" {
		t.Error("Privacy does not match")
	}
	if teams.Teams[0].Description != "esa.io official documents" {
		t.Error("Description does not match")
	}
	if teams.Teams[0].Icon != "https://img.esa.io/uploads/production/teams/105/icon/thumb_m_0537ab827c4b0c18b60af6cdd94f239c.png" {
		t.Error("Icon does not match")
	}
	if teams.Teams[0].Url != "https://docs.esa.io/" {
		t.Error("Url does not match")
	}
	if teams.PrevPage.String() != "" {
		t.Error("PrevPage does not match")
	}
	if teams.NextPage.String() != "1" {
		t.Error("NextPage does not match")
	}
	if teams.TotalCount != 1 {
		t.Error("TotalCount does not match")
	}
}

func TestGetTeam(t *testing.T) {

	testServer := httptest.NewServer(teamHandler)
	defer testServer.Close()
	team, err := fakeClient(testServer.URL).GetTeam()

	if err != nil {
		t.Error("Error occurred")
	}
	if team.Name != "docs" {
		t.Error("Name does not match")
	}
	if team.Privacy != "open" {
		t.Error("Privacy does not match")
	}
	if team.Description != "esa.io official documents" {
		t.Error("Description does not match")
	}
	if team.Icon != "https://img.esa.io/uploads/production/teams/105/icon/thumb_m_0537ab827c4b0c18b60af6cdd94f239c.png" {
		t.Error("Icon does not match")
	}
	if team.Url != "https://docs.esa.io/" {
		t.Error("Url does not match")
	}
}

func TestGetStats(t *testing.T) {

	testServer := httptest.NewServer(statsHandler)
	defer testServer.Close()
	stats, err := fakeClient(testServer.URL).GetStats()

	if err != nil {
		t.Error("Error occurred")
	}
	if stats.Members != 20 {
		t.Error("Name does not match")
	}
	if stats.Posts != 1959 {
		t.Error("Posts does not match")
	}
	if stats.Comments != 2695 {
		t.Error("Comments does not match")
	}
	if stats.Stars != 3115 {
		t.Error("Stars does not match")
	}
	if stats.DailyActiveUsers != 8 {
		t.Error("DailyActiveUsers does not match")
	}
	if stats.WeeklyActiveUsers != 14 {
		t.Error("WeeklyActiveUsers does not match")
	}
	if stats.MonthlyActiveUsers != 15 {
		t.Error("MonthlyActiveUsers does not match")
	}
}

func TestGetMembers(t *testing.T) {

	testServer := httptest.NewServer(membersHandler)
	defer testServer.Close()
	members, err := fakeClient(testServer.URL).GetMembers()

	if err != nil {
		t.Error("Error occurred")
	}
	if members.Members[0].Name != "Hiroaki Sano" {
		t.Error("Name does not match")
	}
	if members.Members[1].Name != "Sano Hiroaki" {
		t.Error("Name does not match")
	}
	if members.Members[0].ScreenName != "hiroakis" {
		t.Error("ScreenName does not match")
	}
	if members.Members[1].ScreenName != "sano" {
		t.Error("ScreenName does not match")
	}
	if members.Members[0].Icon != "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("Icon does not match")
	}
	if members.Members[1].Icon != "https://img.esa.io/uploads/production/users/2/icon/thumb_m_2690997f07b7de3014a36d90827603d6.jpg" {
		t.Error("Icon does not match")
	}
	if members.Members[0].Email != "hiroakis@example.com" {
		t.Error("Email does not match")
	}
	if members.Members[1].Email != "sano@example.com" {
		t.Error("Email does not match")
	}
	if members.PrevPage.String() != "" {
		t.Error("PrevPage does not match")
	}
	if members.NextPage.String() != "1" {
		t.Error("NextPage does not match")
	}
	if members.TotalCount != 2 {
		t.Error("TotalCount does not match")
	}
}

func TestGetPosts(t *testing.T) {

	testServer := httptest.NewServer(postsHandler)
	defer testServer.Close()
	posts, err := fakeClient(testServer.URL).GetPosts()

	if err != nil {
		t.Error("Error occurred")
	}
	if posts.Posts[0].Number != 1 {
		t.Error("Number does not match")
	}
	if posts.Posts[0].Name != "hi!" {
		t.Error("Name does not match")
	}
	if posts.Posts[0].FullName != "日報/2015/05/09/hi! #api #dev" {
		t.Error("FullName does not match")
	}
	if posts.Posts[0].Wip != true {
		t.Error("Wip does not match")
	}
	if posts.Posts[0].BodyMd != "# Getting Started" {
		t.Error("BodyMd does not match")
	}
	if posts.Posts[0].BodyHtml != "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n" {
		t.Error("BodyHtml does not match")
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-09T11:54:50+09:00")
	if posts.Posts[0].CreatedAt != createdAt {
		t.Error("CreatedAt does not match")
	}
	if posts.Posts[0].Message != "Add Getting Started section" {
		t.Error("Message does not match")
	}
	if posts.Posts[0].Url != "https://docs.esa.io/posts/1" {
		t.Error("Url does not match")
	}
	// "2015-05-09T11:54:51+09:00"
	updatedAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-09T11:54:51+09:00")
	if posts.Posts[0].UpdatedAt != updatedAt {
		t.Error("UpdatedAt does not match")
	}
	if posts.Posts[0].Tags[0] != "api" {
		t.Error("Tag[0] does not match")
	}
	if posts.Posts[0].Tags[1] != "dev" {
		t.Error("Tag[1] does not match")
	}
	if posts.Posts[0].Category != "日報/2015/05/09" {
		t.Error("Category does not match")
	}
	if posts.Posts[0].RevisionNumber != 1 {
		t.Error("RevisionNumber does not match")
	}
	if posts.Posts[0].CreatedBy.Name != "Hiroaki Sano" {
		t.Error("CreatedBy.Name does not match")
	}
	if posts.Posts[0].CreatedBy.ScreenName != "hiroakis" {
		t.Error("CreatedBy.ScreenName does not match")
	}
	if posts.Posts[0].CreatedBy.Icon != "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("CreatedBy.Icon does not match")
	}
	if posts.Posts[0].UpdatedBy.Name != "Hiroaki Sano" {
		t.Error("UpdatedBy.Name does not match")
	}
	if posts.Posts[0].UpdatedBy.ScreenName != "hiroakis" {
		t.Error("UpdatedBy.ScreenName does not match")
	}
	if posts.Posts[0].UpdatedBy.Icon != "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("UpdatedBy.Icon does not match")
	}
	if posts.PrevPage.String() != "" {
		t.Error("PrevPage does not match")
	}
	if posts.NextPage.String() != "1" {
		t.Error("NextPage does not match")
	}
	if posts.TotalCount != 1 {
		t.Error("TotalCount does not match")
	}
}

func TestGetPost(t *testing.T) {

	testServer := httptest.NewServer(postHandler)
	defer testServer.Close()
	post, err := fakeClient(testServer.URL).GetPost(1)

	if err != nil {
		t.Error("Error occurred")
	}
	if post.Number != 1 {
		t.Error("Number does not match")
	}
	if post.Name != "hi!" {
		t.Error("Name does not match")
	}
	if post.FullName != "日報/2015/05/09/hi! #api #dev" {
		t.Error("FullName does not match")
	}
	if post.Wip != true {
		t.Error("Wip does not match")
	}
	if post.BodyMd != "# Getting Started" {
		t.Error("BodyMd does not match")
	}
	if post.BodyHtml != "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n" {
		t.Error("BodyHtml does not match")
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-09T11:54:50+09:00")
	if post.CreatedAt != createdAt {
		t.Error("CreatedAt does not match")
	}
	if post.Message != "Add Getting Started section" {
		t.Error("Message does not match")
	}
	if post.Url != "https://docs.esa.io/posts/1" {
		t.Error("Url does not match")
	}
	// "2015-05-09T11:54:51+09:00"
	updatedAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-09T11:54:51+09:00")
	if post.UpdatedAt != updatedAt {
		t.Error("UpdatedAt does not match")
	}
	if post.Tags[0] != "api" {
		t.Error("Tag[0] does not match")
	}
	if post.Tags[1] != "dev" {
		t.Error("Tag[1] does not match")
	}
	if post.Category != "日報/2015/05/09" {
		t.Error("Category does not match")
	}
	if post.RevisionNumber != 1 {
		t.Error("RevisionNumber does not match")
	}
	if post.CreatedBy.Name != "Hiroaki Sano" {
		t.Error("CreatedBy.Name does not match")
	}
	if post.CreatedBy.ScreenName != "hiroakis" {
		t.Error("CreatedBy.ScreenName does not match")
	}
	if post.CreatedBy.Icon != "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("CreatedBy.Icon does not match")
	}
	if post.UpdatedBy.Name != "Hiroaki Sano" {
		t.Error("UpdatedBy.Name does not match")
	}
	if post.UpdatedBy.ScreenName != "hiroakis" {
		t.Error("UpdatedBy.ScreenName does not match")
	}
	if post.UpdatedBy.Icon != "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("UpdatedBy.Icon does not match")
	}
	if post.Kind != "flow" {
		t.Error("Kind does not match")
	}
	if post.CommentsCount != 1 {
		t.Error("CommentsCount does not match")
	}
	if post.TasksCount != 1 {
		t.Error("TasksCount does not match")
	}
	if post.DoneTasksCount != 1 {
		t.Error("DoneTasksCount does not match")
	}
	if post.StargazersCount != 1 {
		t.Error("StargazersCount does not match")
	}
	if post.WatchersCount != 1 {
		t.Error("WatchersCount does not match")
	}
	if post.Star != true {
		t.Error("Star does not match")
	}
	if post.Watch != true {
		t.Error("Watch does not match")
	}
}

func TestCreatePost(t *testing.T) {
	testServer := httptest.NewServer(createPostHandler)
	defer testServer.Close()

	reqPost :=
		request.Post{
			Name:   "hi!",
			BodyMd: "# Getting Started\n",
			Tags: []string{
				"api",
				"dev",
			},
			Category: "dev/2015/05/10",
			Wip:      false,
			Message:  "Add Getting Started section",
		}
	post, err := fakeClient(testServer.URL).CreatePost(reqPost)

	if err != nil {
		t.Error(err)
	}
	if post.Number != 1 {
		t.Error("Number does not match")
	}
	if post.Name != "hi!" {
		t.Error("Name does not match")
	}
	if post.FullName != "日報/2015/05/09/hi! #api #dev" {
		t.Error("FullName does not match")
	}
	if post.Wip != true {
		t.Error("Wip does not match")
	}
	if post.BodyMd != "# Getting Started" {
		t.Error("BodyMd does not match")
	}
	if post.BodyHtml != "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n" {
		t.Error("BodyHtml does not match")
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-09T11:54:50+09:00")
	if post.CreatedAt != createdAt {
		t.Error("CreatedAt does not match")
	}
	if post.Message != "Add Getting Started section" {
		t.Error("Message does not match")
	}
	if post.Url != "https://docs.esa.io/posts/1" {
		t.Error("Url does not match")
	}
	// "2015-05-09T11:54:51+09:00"
	updatedAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-09T11:54:51+09:00")
	if post.UpdatedAt != updatedAt {
		t.Error("UpdatedAt does not match")
	}
	if post.Tags[0] != "api" {
		t.Error("Tag[0] does not match")
	}
	if post.Tags[1] != "dev" {
		t.Error("Tag[1] does not match")
	}
	if post.Category != "日報/2015/05/09" {
		t.Error("Category does not match")
	}
	if post.RevisionNumber != 1 {
		t.Error("RevisionNumber does not match")
	}
	if post.CreatedBy.Name != "Hiroaki Sano" {
		t.Error("CreatedBy.Name does not match")
	}
	if post.CreatedBy.ScreenName != "hiroakis" {
		t.Error("CreatedBy.ScreenName does not match")
	}
	if post.CreatedBy.Icon != "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("CreatedBy.Icon does not match")
	}
	if post.UpdatedBy.Name != "Hiroaki Sano" {
		t.Error("UpdatedBy.Name does not match")
	}
	if post.UpdatedBy.ScreenName != "hiroakis" {
		t.Error("UpdatedBy.ScreenName does not match")
	}
}

func TestUpdatePost(t *testing.T) {
	testServer := httptest.NewServer(updatePostHandler)
	defer testServer.Close()

	reqPost :=
		request.Post{
			Name:   "hi!",
			BodyMd: "# Getting Started\n",
			Tags: []string{
				"api",
				"dev",
			},
			Category: "dev/2015/05/10",
			Wip:      false,
			Message:  "Add Getting Started section",
			OriginalRevision: request.OriginalRevision{
				BodyMd: "# Getting ...",
				Number: 1,
				User:   "hiroakis",
			},
		}
	post, err := fakeClient(testServer.URL).UpdatePost(1, reqPost)

	if err != nil {
		t.Error(err)
	}
	if post.Number != 1 {
		t.Error("Number does not match")
	}
	if post.Name != "hi!" {
		t.Error("Name does not match")
	}
	if post.FullName != "日報/2015/05/09/hi! #api #dev" {
		t.Error("FullName does not match")
	}
	if post.Wip != true {
		t.Error("Wip does not match")
	}
	if post.BodyMd != "# Getting Started" {
		t.Error("BodyMd does not match")
	}
	if post.BodyHtml != "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n" {
		t.Error("BodyHtml does not match")
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-09T11:54:50+09:00")
	if post.CreatedAt != createdAt {
		t.Error("CreatedAt does not match")
	}
	if post.Message != "Add Getting Started section" {
		t.Error("Message does not match")
	}
	if post.Url != "https://docs.esa.io/posts/1" {
		t.Error("Url does not match")
	}
	// "2015-05-09T11:54:51+09:00"
	updatedAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-09T11:54:51+09:00")
	if post.UpdatedAt != updatedAt {
		t.Error("UpdatedAt does not match")
	}
	if post.Tags[0] != "api" {
		t.Error("Tag[0] does not match")
	}
	if post.Tags[1] != "dev" {
		t.Error("Tag[1] does not match")
	}
	if post.Category != "日報/2015/05/09" {
		t.Error("Category does not match")
	}
	if post.RevisionNumber != 1 {
		t.Error("RevisionNumber does not match")
	}
	if post.CreatedBy.Name != "Hiroaki Sano" {
		t.Error("CreatedBy.Name does not match")
	}
	if post.CreatedBy.ScreenName != "hiroakis" {
		t.Error("CreatedBy.ScreenName does not match")
	}
	if post.CreatedBy.Icon != "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("CreatedBy.Icon does not match")
	}
	if post.UpdatedBy.Name != "Hiroaki Sano" {
		t.Error("UpdatedBy.Name does not match")
	}
	if post.UpdatedBy.ScreenName != "hiroakis" {
		t.Error("UpdatedBy.ScreenName does not match")
	}
}

func TestDeletePost(t *testing.T) {
	testServer := httptest.NewServer(deletePostHandler)
	defer testServer.Close()
	deleted, err := fakeClient(testServer.URL).DeletePost(1)

	if err != nil {
		t.Error(err)
	}
	if deleted != true {
		t.Error("error")
	}
}

func TestGetComments(t *testing.T) {

	testServer := httptest.NewServer(commentsHandler)
	defer testServer.Close()
	comments, err := fakeClient(testServer.URL).GetComments(1)

	if err != nil {
		t.Error("Error occurred")
	}
	if comments.Comments[0].Id != 1 {
		t.Error("Id does not match")
	}
	if comments.Comments[0].BodyMd != "(大事)" {
		t.Error("BodyMd does not match")
	}
	if comments.Comments[0].BodyHtml != "<p>(大事)</p>" {
		t.Error("BodyHtml does not match")
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2014-05-10T12:45:42+09:00")
	if comments.Comments[0].CreatedAt != createdAt {
		t.Error("CreatedAt does not match")
	}
	if comments.Comments[0].Url != "https://docs.esa.io/posts/2#comment-1" {
		t.Error("Url does not match")
	}
	updatedAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2014-05-18T23:02:29+09:00")
	if comments.Comments[0].UpdatedAt != updatedAt {
		t.Error("UpdatedAt does not match")
	}
	if comments.Comments[0].CreatedBy.Name != "Hiroaki Sano" {
		t.Error("CreatedBy.Name does not match")
	}
	if comments.Comments[0].CreatedBy.ScreenName != "hiroakis" {
		t.Error("CreatedBy.ScreenName does not match")
	}
	if comments.Comments[0].CreatedBy.Icon != "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("CreatedBy.Icon does not match")
	}
	if comments.PrevPage.String() != "" {
		t.Error("PrevPage does not match")
	}
	if comments.NextPage.String() != "1" {
		t.Error("NextPage does not match")
	}
	if comments.TotalCount != 1 {
		t.Error("TotalCount does not match")
	}
}

func TestGetComment(t *testing.T) {

	testServer := httptest.NewServer(commentHandler)
	defer testServer.Close()
	comment, err := fakeClient(testServer.URL).GetComment(1)

	if err != nil {
		t.Error("Error occurred")
	}
	if comment.Id != 13 {
		t.Error("Id does not match")
	}
	if comment.BodyMd != "読みたい" {
		t.Error("BodyMd does not match")
	}
	if comment.BodyHtml != "<p>読みたい</p>" {
		t.Error("BodyHtml does not match")
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2014-05-13T16:17:42+09:00")
	if comment.CreatedAt != createdAt {
		t.Error("CreatedAt does not match")
	}
	updatedAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2014-05-18T23:02:29+09:00")
	if comment.UpdatedAt != updatedAt {
		t.Error("UpdatedAt does not match")
	}
	if comment.Url != "https://docs.esa.io/posts/13#comment-13" {
		t.Error("Url does not match")
	}
	if comment.CreatedBy.Name != "Sano Hiroaki" {
		t.Error("CreatedBy.Name does not match")
	}
	if comment.CreatedBy.ScreenName != "sano" {
		t.Error("CreatedBy.ScreenName does not match")
	}
	if comment.CreatedBy.Icon != "https://img.esa.io/uploads/production/users/2/icon/thumb_m_2690997f07b7de3014a36d90827603d6.jpg" {
		t.Error("CreatedBy.Icon does not match")
	}
}

func TestCreateComment(t *testing.T) {

	testServer := httptest.NewServer(createCommentHandler)
	defer testServer.Close()

	reqComment := request.Comment{
		BodyMd: "LGTM!",
	}
	comment, err := fakeClient(testServer.URL).CreateComment(2, reqComment)

	if err != nil {
		t.Error(err)
	}
	if comment.Id != 22767 {
		t.Error("Id does not match")
	}
	if comment.BodyMd != "LGTM!" {
		t.Error("BodyMd does not match")
	}
	if comment.BodyHtml != "<p>LGTM!</p>\n" {
		t.Error("BodyHtml does not match")
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-06-21T19:36:20+09:00")
	if comment.CreatedAt != createdAt {
		t.Error("CreatedAt does not match")
	}
	updatedAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-06-21T19:36:20+09:00")
	if comment.UpdatedAt != updatedAt {
		t.Error("UpdatedAt does not match")
	}
	if comment.Url != "https://docs.esa.io/posts/2#comment-22767" {
		t.Error("Url does not match")
	}
	if comment.CreatedBy.Name != "Hiroaki Sano" {
		t.Error("CreatedBy.Name does not match")
	}
	if comment.CreatedBy.ScreenName != "hiroakis" {
		t.Error("CreatedBy.ScreenName does not match")
	}
	if comment.CreatedBy.Icon != "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("CreatedBy.Icon does not match")
	}
}

func TestUpdateComment(t *testing.T) {

	testServer := httptest.NewServer(updateCommentHandler)
	defer testServer.Close()

	reqComment := request.Comment{
		BodyMd: "LGTM!!!",
	}
	comment, err := fakeClient(testServer.URL).UpdateComment(2, reqComment)

	if err != nil {
		t.Error(err)
	}
	if comment.Id != 22767 {
		t.Error("Id does not match")
	}
	if comment.BodyMd != "LGTM!!!" {
		t.Error("BodyMd does not match")
	}
	if comment.BodyHtml != "<p>LGTM!!!</p>\n" {
		t.Error("BodyHtml does not match")
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-06-21T19:36:20+09:00")
	if comment.CreatedAt != createdAt {
		t.Error("CreatedAt does not match")
	}
	updatedAt, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-06-21T19:40:33+09:00")
	if comment.UpdatedAt != updatedAt {
		t.Error("UpdatedAt does not match")
	}
	if comment.Url != "https://docs.esa.io/posts/2#comment-22767" {
		t.Error("Url does not match")
	}
	if comment.CreatedBy.Name != "Hiroaki Sano" {
		t.Error("CreatedBy.Name does not match")
	}
	if comment.CreatedBy.ScreenName != "hiroakis" {
		t.Error("CreatedBy.ScreenName does not match")
	}
	if comment.CreatedBy.Icon != "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png" {
		t.Error("CreatedBy.Icon does not match")
	}
}

func TestDeleteComment(t *testing.T) {
	testServer := httptest.NewServer(deleteCommentHandler)
	defer testServer.Close()
	deleted, err := fakeClient(testServer.URL).DeleteComment(1)

	if err != nil {
		t.Error(err)
	}
	if deleted != true {
		t.Error("error")
	}
}
