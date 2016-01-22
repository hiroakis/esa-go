# esa-go

esa API v1 client written in go.

## Installation

```
go get github.com/hiroakis/esa-go
```

## Basic Usage

```
import (
    "encoding/json"
    "fmt"
    esa "github.com/hiroakis/esa-go"
)

func main() {
    // Initializing client
    c := esa.NewEsaClient("API_KEY", "TEAM_NAME")

    // get all posts
    posts, err := c.GetPosts()
    if err != nil {
            fmt.Println(err)
    }
    fmt.Println(posts)

    // print specified field
    fmt.Println(posts.Posts[0].Name)

    // print with json string
    postsJson, _ := json.Marshal(posts)
    fmt.Println(string(postsJson))

    // Pagenation
    c.SetPage(1)
    posts, err = c.GetPosts()
    if err != nil {
            fmt.Println(err)
    }
    fmt.Println(posts)

    // Query
    c.SetQuery("category:memo")
    posts, err = c.GetPosts()
    if err != nil {
            fmt.Println(err)
    }
    fmt.Println(posts)
}
```

## Examples

```
    // teams
    teams, err := c.GetTeams()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(teams)

    // team
    c.SetTeam("TEAM_NAME")
    team, err := c.GetTeam()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(team)

    // stats
    stats, err := c.GetStats()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(stats)

    // members
    members, err := c.GetMembers()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(members)

    // post
    post, err := c.GetPost(1)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(post)

    // posts
    posts, err := c.GetPosts()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(posts)

    // posts
    c.SetQuery("category:memo")
    posts, err := c.GetPosts()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(posts)

    // create new post
    postContent :=
        PostContent{
            Name:     "hi!",
            BodyMd:   "hello",
            Category: "Users/hiroakis/memo",
        }

    createdPost, err := c.CreatePost(postContent)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(createdPost)

    // update post
    postContent :=
        PostContent{
            Name:     "hi!",
            BodyMd:   "おは",
            Category: "Users/hiroakis/memo",
        }

    updatedPost, err := c.UpdatePost(549, postContent)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(updatedPost)

    // delete post
    deletedPost, err := c.DeletePost(549)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(deletedPost)

    // comments
    comments, err := c.GetComments(543)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(comments)

    // comment
    comment, err := c.GetComment(80737)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(comment)

    // create comment
    commentContent := CommentContent{
        BodyMd: "comment!",
    }
    createdComment, err := c.CreateComment(543, commentContent)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(createdComment)

    // update comment
    commentContent = CommentContent{
        BodyMd: "comment!!!!",
    }
    updatedComment, err := c.UpdateComment(80737, commentContent)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(updatedComment)

    deletedComment, err := c.DeleteComment(80737)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(deletedComment)
```

## TODO

testing

## License

MIT