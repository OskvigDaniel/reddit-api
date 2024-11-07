package store

import (
	"maps"
	"slices"
	"sort"
	"strconv"
)

type PostStore struct {
	Posts    map[string]Post
	TopPosts []Post
}
type Post struct {
	Title   string
	Url     string
	Upvotes int
}

func (ps PostStore) GetPostsByUpvoteCount() {
	posts := slices.Collect(maps.Values(ps.Posts))
	sort.Slice(posts, func(i, j int) bool {
		// Sort posts by upvotes descending
		return posts[i].Upvotes > posts[j].Upvotes
	})

	ps.TopPosts = posts
}

func (ps PostStore) FormatTopPosts() string {
	output := "Top posts:\n"
	count := len(ps.TopPosts)
	if count > 5 {
		count = 5
	}
	// Arbitrarily display the top 5 posts
	for i := range count {
		post := ps.TopPosts[i]
		output += post.Title + " with " + strconv.Itoa(post.Upvotes) + " posts\n"
	}

	return output
}
