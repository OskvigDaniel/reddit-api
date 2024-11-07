package monitor

import (
	"fmt"
	"reddit-api/internal/integration/reddit"
	"reddit-api/internal/store"
	"time"
)

func StartMonitor(us *store.UserStore, ps *store.PostStore, clientId string, clientSecret string, subreddit string) {
	// Reddit api allows 100 queries per minute before throttling,
	// the query time has been padded to allow for breathing room
	fmt.Print("Starting monitor loop\n")
	ticker := time.NewTicker(2000 * time.Millisecond)

	go monitor(ticker, us, ps, clientId, clientSecret, subreddit)
}

func monitor(ticker *time.Ticker, us *store.UserStore, ps *store.PostStore, clientId string, clientSecret string, subreddit string) {
	token := reddit.GetAuthToken(clientId, clientSecret)
	fmt.Println("Auth token received: " + token)
	var lastSeenPost string
	for {
		<-ticker.C
		fmt.Print("Searching for posts...\n")
		listing, err := reddit.GetPostsForSubreddit(token, lastSeenPost, subreddit)

		if err != nil {
			// if get posts returns an error without crashing, it is
			// due to rate limiting.
			// TODO: properly handle rate limits
			ticker.Reset(10000 * time.Millisecond)
		}

		if len(listing.Data.Children) == 0 {
			//TODO: remove the need to modify control flow
			continue
		}

		// If this is our first seen post, mark it as last seen and wait for next loop
		if lastSeenPost == "" && len(listing.Data.Children) > 0 {
			lastSeenPost = listing.Data.Children[0].Data.Name
			fmt.Println("first post found " + lastSeenPost)
		} else {
			fmt.Println("Posts found")
			for i := range listing.Data.Children {
				post := listing.Data.Children[i]
				ps.Posts[post.Data.ID] = store.Post{Title: post.Data.Title, Url: post.Data.URL, Upvotes: post.Data.Ups}
				user := us.Users[post.Data.Author]
				if user.Username == "" {
					us.Users[post.Data.Author] = store.User{Username: post.Data.Author, PostCount: 1}
				} else {
					us.Users[post.Data.Author] = store.User{Username: post.Data.Author, PostCount: user.PostCount + 1}
				}
			}

			ps.GetPostsByUpvoteCount()
			us.GetUsersByPostCount()

			fmt.Print(ps.FormatTopPosts())
			fmt.Print(us.FormatTopUsers())
		}
	}
}
