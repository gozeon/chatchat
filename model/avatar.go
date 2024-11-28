package model

import (
	"chatchat/database"
	"encoding/json"
)

type Avatar struct {
	WhiteList []string `json:"avatarWhiteList"`
	User      struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"user"`
	Robot struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"robot"`

	db *database.BoltDB
}

func NewAvatar(db *database.BoltDB) Avatar {
	return Avatar{
		db: db,
	}
}

func (b Avatar) Init() {
	b.db.CreateBucketIfNotExists("avatar")

	avatar := Avatar{
		WhiteList: []string{"all"},
		User: struct {
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
		}{
			Name:   "",
			Avatar: "//gw.alicdn.com/tfs/TB1DYHLwMHqK1RjSZFEXXcGMXXa-56-62.svg",
		},
		Robot: struct {
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
		}{
			Name:   "",
			Avatar: "//gw.alicdn.com/tfs/TB1U7FBiAT2gK0jSZPcXXcKkpXa-108-108.jpg",
		},
	}

	v, _ := json.Marshal(avatar)
	b.db.Put("avatar", []byte("config"), v)
}

func (b Avatar) Get() Avatar {
	config, _ := b.db.Get("avatar", []byte("config"))
	var a Avatar
	json.Unmarshal(config, &a)
	return a
}
