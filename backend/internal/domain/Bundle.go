package domain

import "time"

type BundleDB struct{
	Slug string `db:"slug"`
	Title string `db:"title"`
	CreatedAt time.Time `db:"created_at"`
	LastAccessed time.Time `db:"last_accessed"`
}