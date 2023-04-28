package handlers

import "time"

type Post struct {
	Title     string    `json:"title"`
	PostDate  time.Time `json:"postDate"`
	PostOwner string    `json:"postOwner"`
}

type Database struct {
	posts []Post
}

func (d *Database) addPost(post Post) {
	d.posts = append(d.posts, post)
}

func (d *Database) getPostByOwner(owner string) (Post, bool) {
	for _, post := range d.posts {
		if post.PostOwner == owner {
			return post, true
		}
	}
	return Post{}, false
}

func (d *Database) getPostsList() []Post {
	posts := []Post{
		Post{
			Title:     "New Goal Here.",
			PostOwner: "Ahmad Dar",
			PostDate:  time.Now(),
		},
		Post{
			Title:     "Trying to get a port bar.",
			PostOwner: "Maher tom",
			PostDate:  time.Now().Add(time.Duration(5 * 10000)),
		},
		Post{
			Title:     "Start for the sky when u are the manual.",
			PostOwner: "Tom Jerry",
			PostDate:  time.Now().Truncate(time.Duration(12 * 10000)),
		},
	}

	return posts

}
