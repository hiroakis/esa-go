package esa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hiroakis/esa-go/request"
	"github.com/hiroakis/esa-go/response"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	EsaAPIv1 = "https://api.esa.io/v1"
	// EsaAPIv1 = "http://localhost:5000"
)

type EsaClient struct {
	Team        string
	AccessToken string
	Api         string
	Page        int
	Query       string
	Client      *http.Client
}

func NewEsaClient(accessToken, team string) *EsaClient {

	esaClient := &EsaClient{
		AccessToken: accessToken,
		Team:        team,
		Api:         EsaAPIv1,
		Page:        -1,
		Query:       "",
		Client:      &http.Client{Timeout: time.Duration(10 * time.Second)},
	}

	return esaClient
}

func (c *EsaClient) SetTeam(team string) {
	c.Team = team
}

func (c *EsaClient) SetPage(page int) {
	c.Page = page
}

func (c *EsaClient) SetQuery(query string) {
	c.Query = query
}

func (c *EsaClient) SetClient(client *http.Client) {
	c.Client = client
}

func (c *EsaClient) SetApi(api string) {
	c.Api = api
}

func (c *EsaClient) GetTeams() (response.Teams, error) {
	teams := &response.Teams{}
	endpoint := fmt.Sprintf("%s/teams", c.Api)

	resp := c.sendGetRequest(endpoint)
	body, err := c.chackResponse(resp)
	if err != nil {
		return *teams, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &teams)
	return *teams, err
}

func (c *EsaClient) GetTeam() (response.Team, error) {
	team := &response.Team{}
	endpoint := fmt.Sprintf("%s/teams/%s", c.Api, c.Team)

	resp := c.sendGetRequest(endpoint)
	body, err := c.chackResponse(resp)
	if err != nil {
		return *team, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &team)
	return *team, nil
}

func (c *EsaClient) GetStats() (response.Stats, error) {
	stats := &response.Stats{}
	endpoint := fmt.Sprintf("%s/teams/%s/stats", c.Api, c.Team)

	resp := c.sendGetRequest(endpoint)
	body, err := c.chackResponse(resp)
	if err != nil {
		return *stats, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &stats)
	return *stats, err
}

func (c *EsaClient) GetMembers() (response.Members, error) {
	members := &response.Members{}
	endpoint := fmt.Sprintf("%s/teams/%s/members", c.Api, c.Team)

	resp := c.sendGetRequest(endpoint)
	body, err := c.chackResponse(resp)
	if err != nil {
		return *members, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &members)
	return *members, err
}

func (c *EsaClient) GetPost(postNumber int) (response.Post, error) {
	post := &response.Post{}
	endpoint := fmt.Sprintf("%s/teams/%s/posts/%d", c.Api, c.Team, postNumber)

	resp := c.sendGetRequest(endpoint)
	body, err := c.chackResponse(resp)
	if err != nil {
		return *post, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &post)
	return *post, err
}

func (c *EsaClient) GetPosts() (response.Posts, error) {
	posts := &response.Posts{}
	endpoint := fmt.Sprintf("%s/teams/%s/posts", c.Api, c.Team)

	resp := c.sendGetRequest(endpoint)
	body, err := c.chackResponse(resp)
	if err != nil {
		return *posts, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &posts)
	return *posts, err
}

func (c *EsaClient) CreatePost(reqPost request.Post) (response.Post, error) {
	post := &response.Post{}
	endpoint := fmt.Sprintf("%s/teams/%s/posts", c.Api, c.Team)

	pc, err := json.Marshal(request.PostData{reqPost})
	if err != nil {
		return *post, err
	}

	resp := c.sendPostRequest(endpoint, bytes.NewBuffer(pc))
	body, err := c.chackResponse(resp)
	if err != nil {
		return *post, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &post)
	return *post, err
}

func (c *EsaClient) UpdatePost(postNumber int, reqPost request.Post) (response.Post, error) {
	post := &response.Post{}
	endpoint := fmt.Sprintf("%s/teams/%s/posts/%d", c.Api, c.Team, postNumber)

	pc, err := json.Marshal(request.PostData{reqPost})
	if err != nil {
		return *post, err
	}

	resp := c.sendPatchRequest(endpoint, bytes.NewBuffer(pc))
	body, err := c.chackResponse(resp)
	if err != nil {
		return *post, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &post)
	return *post, err
}

func (c *EsaClient) DeletePost(postNumber int) (bool, error) {
	endpoint := fmt.Sprintf("%s/teams/%s/posts/%d", c.Api, c.Team, postNumber)

	resp := c.sendDeleteRequest(endpoint)
	_, err := c.chackResponse(resp)
	if err != nil {
		return false, err
	}
	defer c.closeHttpResponse(resp)
	return true, err
}

func (c *EsaClient) GetComments(postNumber int) (response.Comments, error) {
	comments := &response.Comments{}
	endpoint := fmt.Sprintf("%s/teams/%s/posts/%d/comments", c.Api, c.Team, postNumber)

	resp := c.sendGetRequest(endpoint)
	body, err := c.chackResponse(resp)
	if err != nil {
		return *comments, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &comments)
	return *comments, err
}

func (c *EsaClient) GetComment(commentNumber int) (response.Comment, error) {
	comment := &response.Comment{}
	endpoint := fmt.Sprintf("%s/teams/%s/comments/%d", c.Api, c.Team, commentNumber)

	resp := c.sendGetRequest(endpoint)
	body, err := c.chackResponse(resp)
	if err != nil {
		return *comment, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &comment)
	return *comment, err
}

func (c *EsaClient) CreateComment(postNumber int, reqComment request.Comment) (response.Comment, error) {
	comment := &response.Comment{}
	endpoint := fmt.Sprintf("%s/teams/%s/posts/%d/comments", c.Api, c.Team, postNumber)

	cc, err := json.Marshal(request.CommentData{reqComment})
	if err != nil {
		return *comment, err
	}

	resp := c.sendPostRequest(endpoint, bytes.NewBuffer(cc))
	body, err := c.chackResponse(resp)
	if err != nil {
		return *comment, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &comment)
	return *comment, err
}

func (c *EsaClient) UpdateComment(commentId int, reqComment request.Comment) (response.Comment, error) {
	comment := &response.Comment{}
	endpoint := fmt.Sprintf("%s/teams/%s/comments/%d", c.Api, c.Team, commentId)

	cc, err := json.Marshal(request.CommentData{reqComment})
	if err != nil {
		return *comment, err
	}

	resp := c.sendPatchRequest(endpoint, bytes.NewBuffer(cc))
	body, err := c.chackResponse(resp)
	if err != nil {
		return *comment, err
	}
	defer c.closeHttpResponse(resp)

	json.Unmarshal(body, &comment)
	return *comment, err
}

func (c *EsaClient) DeleteComment(commentId int) (bool, error) {
	endpoint := fmt.Sprintf("%s/teams/%s/comments/%d", c.Api, c.Team, commentId)

	resp := c.sendDeleteRequest(endpoint)
	_, err := c.chackResponse(resp)
	if err != nil {
		return false, err
	}
	defer c.closeHttpResponse(resp)
	return true, err
}

func (c *EsaClient) sendHttpRequest(method, endpoint string, data io.Reader) *http.Response {

	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		fmt.Println(err)
	}
	req = c.buildRequest(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

func (c *EsaClient) sendGetRequest(endpoint string) *http.Response {
	resp := c.sendHttpRequest("GET", endpoint, nil)
	return resp
}

func (c *EsaClient) sendPostRequest(endpoint string, data io.Reader) *http.Response {
	resp := c.sendHttpRequest("POST", endpoint, data)
	return resp
}

func (c *EsaClient) sendPatchRequest(endpoint string, data io.Reader) *http.Response {
	resp := c.sendHttpRequest("PATCH", endpoint, data)
	return resp
}

func (c *EsaClient) sendDeleteRequest(endpoint string) *http.Response {
	resp := c.sendHttpRequest("DELETE", endpoint, nil)
	return resp
}

func (c *EsaClient) buildRequest(req *http.Request) *http.Request {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	values := url.Values{}
	if c.Page != -1 {
		values.Add("page", fmt.Sprintf("%d", c.Page))
	}
	if c.Query != "" {
		values.Add("q", c.Query)
	}
	req.URL.RawQuery = values.Encode()
	return req
}

func (c *EsaClient) chackResponse(resp *http.Response) ([]byte, error) {
	var body []byte
	var err error

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		err = fmt.Errorf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	body, _ = ioutil.ReadAll(resp.Body)
	return body, err
}

func (c *EsaClient) closeHttpResponse(resp *http.Response) {
	resp.Body.Close()
}
