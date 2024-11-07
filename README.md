# Reddit-Api Local Development
To get started, run the following:

```go build cmd/main/main.go```

```go ./main -id=<id here> -secret=<secret-here> -subreddit=/r/programming```

Note that subreddit strings must contain the leading /r/

There exist flags for both http host and port via `-host` and `-port`

The console will log posts and users, arbitrarily with a top 5 list.
There is an http listener, but not completely tested. Instead of the top 5 of each list each corresponding endpoint will return its full list.
# WIP
This project is very far from finished. The following items are still in need of completion:
- Unit testing
- Properly enriched logging
- Containerization (currently very easy with just the single binary needed to go)
- Proper make file
- Better error handling across the board
- Make proper use of the response headers for rate limiting
- Parameterize the subreddit so it can be passed in as a flag