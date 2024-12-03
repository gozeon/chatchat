package model

import "chatchat/database"

type Other struct {
	Placeholder string `json:"placeholder"`
	Sidebar     string `json:"sidebar"`

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
		b.db.Put("other", []byte("sidebar"), []byte(`[
    {
        "code": "sidebar-suggestion",
        "data": [
            {
                "label": "tab1",
                "list": [
                    "图片",
                    "文字",
                    "视频"
                ]
            },
            {
                "label": "tab2",
                "list": [
                    "所有组件",
                    "不回复"
                ]
            }
        ]
    },
    {
        "code": "sidebar-tool",
        "title": "常用工具",
        "data": [
            {
                "name": "GitHub",
                "icon": "clipboard-list",
                "url": "https://github.com/gozeon/chatchat"
            },
            {
                "name": "Chatui",
                "icon": "exclamation-shield",
                "url": "https://chatui.io/"
            },
            {
                "name": "ChatSDK",
                "icon": "history",
                "url": "https://chatbot.console.aliyun.com/ChatSDK"
            }
        ]
    },
    {
        "code": "sidebar-phone",
        "title": "全省统一政务服务热线",
        "data": [
            "12345"
        ]
    }
]`))
	}
}

func (b Other) Get() Other {
	placeholder, _ := b.db.Get("other", []byte("placeholder"))
	sidebar, _ := b.db.Get("other", []byte("sidebar"))

	return Other{
		Placeholder: string(placeholder),
		Sidebar:     string(sidebar),
	}
}
