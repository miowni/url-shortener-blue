# Blue Coding Golang Challange

## Setting up

Clone this repository by running the following command:

```
git clone git@github.com:miowni/url-shortener-blue.git
```

Build it using Go, by running:

```
cd url-shortener-blue
go build
```

And then run the binary file:

```
./url-shortener-blue
```

## Requirements

- We must be able to put a URL into the home page and get back a URL of the shortest possible length.
  - This requirement was implemented during the interview.
- We must be redirected to the full URL when we enter the short URL (ex: http://localhost:3000/a => https://google.com)
  - This requirement was implemented during the interview.
- There should be a page that shows the top 100 most frequently accessed URLs.
  - This requirement was implemented during the interview.
- There must be a concurrent task that crawls the URL being shortened, pulls the \<title> from the website and stores it.
  - How I would implement it: run a goroutine that fetches the title using Gin and update the title in the database with the already existing function `UpdateTitle()`.
- Display the title with the URL on the top 100 board.
  - How I would implement it: just fetch the title from the database inside the already existing function `GetTop100()`, since it would be already stored.
- There must be a README that explains how to setup the application and the algorithm used for generating the URL short code.
  - This task was completed after the interview.

## Additional information

This code was posted on my alternative GitHub account only to avoid professional issues.
