package main

import (
	"sort"
)

// Short link type
type ShortLink struct {
	Url     string
	Title   string
	Acesses int
}

// Database type
type Database struct {
	data   map[string]ShortLink
	length int
}

// Instance a new Database
func NewDatabase() *Database {
	return &Database{
		data:   map[string]ShortLink{},
		length: 0,
	}
}

// Get top 100 visited urls
func (d *Database) GetTop100() []string {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range d.data {
		ss = append(ss, kv{k, v.Acesses})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	ranking := []string{}

	for _, kv := range ss {
		ranking = append(ranking, kv.Key)
	}

	left := 100
	if len(ss) < 100 {
		left = len(ss)
	}

	return ranking[:left]
}

// Get length of database
func (d Database) Length() int {
	return d.length
}

// Push URL to database
func (d *Database) Push(url, token string) {
	d.data[token] = ShortLink{Url: url}
	d.length++
}

// Update url title
func (d *Database) UpdateTitle(token, title string) bool {
	shortLink, ok := d.data[token]
	if ok {
		shortLink.Title = title
		d.data[token] = shortLink
	}
	return ok
}

// Push URL to database
func (d *Database) Get(token string) (string, bool) {
	shortLink, ok := d.data[token]
	if ok {
		shortLink.Acesses++
		d.data[token] = shortLink
	}
	return shortLink.Url, ok
}
