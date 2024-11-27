package model

import "chatchat/database"

type Band struct {
	Logo string
	Url  string
	Name string

	db *database.BoltDB
}

func NewBand(db *database.BoltDB) Band {
	return Band{
		db: db,
	}
}

func (b Band) Init() {
	b.db.CreateBucketIfNotExists("band")
	b.db.Put("band", []byte("logo"), []byte("/static/logo.png"))
	b.db.Put("band", []byte("url"), []byte("/chat"))
	b.db.Put("band", []byte("name"), []byte("客服"))
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
