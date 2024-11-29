package model

import (
	"chatchat/database"
)

type Message struct {
	db *database.BoltDB
}

func NewMessage(db *database.BoltDB) Message {
	return Message{
		db: db,
	}
}

func (b Message) Init() {
	message, _ := b.db.GetBucket("message")

	if message == nil {
		b.db.CreateBucketIfNotExists("message")

		message := `[{"code":"system","data":{"text":"智能助理进入对话，为您服务"}},{"code":"text","data":{"text":"智能助理为您服务，请问有什么可以帮您？"}}]`
		b.db.Put("message", []byte("default"), []byte(message))

		message = `
[
    {
        "name": "发送文本",
        "isHighlight": true
    },
    {
        "name": "可见文本",
        "type": "text",
        "text": "实际发送的文本",
        "isNew": true
    },
    {
        "name": "点击跳转",
        "type": "url",
        "url": "https://github.com/gozeon/chatchat"
    },
    {
        "name": "唤起卡片",
        "type": "card",
        "card": {
            "code": "knowledge",
            "data": {
                "text": "<p><strong><span style='color: rgb(32, 33, 34); text-wrap-mode: wrap; background-color: rgb(255, 255, 255);'>莫聽穿林打葉聲</span></strong><span style='color: rgb(32, 33, 34); text-wrap-mode: wrap; background-color: rgb(255, 255, 255);'>，何妨吟嘯且徐行；</span><br/><span style='color: rgb(32, 33, 34); text-wrap-mode: wrap; background-color: rgb(255, 255, 255);'>竹杖芒鞋輕勝馬，誰怕？</span><span style='text-wrap-mode: wrap; background-color: rgb(255, 255, 255); color: rgb(255, 0, 0);'>一蓑烟雨任平生</span><span style='color: rgb(32, 33, 34); text-wrap-mode: wrap; background-color: rgb(255, 255, 255);'>。</span><br/><br/><span style='color: rgb(32, 33, 34); text-wrap-mode: wrap; background-color: rgb(255, 255, 255);'>料峭春風吹酒醒，微冷、山頭斜照卻相迎；</span><br/><span style='color: rgb(32, 33, 34); text-wrap-mode: wrap; background-color: rgb(255, 255, 255);'>回首向來蕭瑟處，歸去、也無風雨也無晴。</span></p>"
            }
        }
    }
]`
		b.db.Put("message", []byte("quick"), []byte(message))

	}
}

func (b Message) Get(key string) any {
	msg, _ := b.db.Get("message", []byte(key))

	return string(msg)
}
