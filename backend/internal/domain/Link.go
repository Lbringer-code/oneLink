package domain

import "time"


type LinkDB struct{
	BundleSlug string `db:"bundle_slug"`
	Url string `db:"url"`
	Note string `db:"note"`
	DisplayText string `db:"display_text"`
	CreatedAt time.Time `db:"created_at"`
} 