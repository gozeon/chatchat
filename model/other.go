package model

import "chatchat/database"

type Other struct {
	Placeholder string `json:"placeholder"`

	db *database.BoltDB
}

func NewOther(db *database.BoltDB) Other {
	return Other{
		db: db,
	}
}

func (b Other) Init() {
	other, _ := b.db.GetBucket("other")

	if other == nil {
		b.db.CreateBucketIfNotExists("other")
		b.db.Put("other", []byte("placeholder"), []byte("输入您想咨询的问题"))
	}
}

func (b Other) Get() Other {
	placeholder, _ := b.db.Get("other", []byte("placeholder"))

	return Other{
		Placeholder: string(placeholder),
	}
}
