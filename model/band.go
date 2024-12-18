package model

import (
	"chatchat/database"
)

type Band struct {
	Logo string `json:"logo"`
	Url  string `json:"url"`
	Name string `json:"name"`

	db *database.BoltDB
}

func NewBand(db *database.BoltDB) Band {
	return Band{
		db: db,
	}
}

func (b Band) Init() {
	band, _ := b.db.GetBucket("band")

	if band == nil {
		b.db.CreateBucketIfNotExists("band")
		b.db.Put("band", []byte("logo"), []byte("/static/logo.png"))
		b.db.Put("band", []byte("url"), []byte("/chat"))
		b.db.Put("band", []byte("name"), []byte("chatchat"))
	}
}

func (b Band) Get() Band {
	logo, _ := b.db.Get("band", []byte("logo"))
	url, _ := b.db.Get("band", []byte("url"))
	name, _ := b.db.Get("band", []byte("name"))
	return Band{
		Logo: string(logo),
		Url:  string(url),
		Name: string(name),
	}
}
