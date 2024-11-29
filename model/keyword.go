package model

import (
	"chatchat/database"
)

type Keyword struct {
	db *database.BoltDB
}

func NewKeyword(db *database.BoltDB) Keyword {
	return Keyword{
		db: db,
	}
}

func (b Keyword) Init() {
	keyword, _ := b.db.GetBucket("keyword")

	if keyword == nil {
		b.db.CreateBucketIfNotExists("keyword")
		dal := `
{
    "code": "list-v2",
    "data": {
        "autoSize": true,
        "type": "default",
        "title": "猜你想问",
        "list": [
            {
                "icon": "",
                "action": "sendText",
                "text": "购买后如何查看信息？text",
                "tag": "",
                "content": "购买后如何查看信息？content",
                "url": ""
            },
            {
                "icon": "",
                "action": "sendText",
                "text": "【夜间】自动发货",
                "tag": "",
                "content": "【夜间】自动发货",
                "url": ""
            },
            {
                "icon": "",
                "action": "openWindow",
                "text": "【安装包】问题汇总",
                "tag": "",
                "content": "",
                "url": "//docs.qq.com/doc/DRVZCeWNGRGVhclhJ"
            },
            {
                "icon": "",
                "action": "sendText",
                "text": "谷歌辅助邮箱怎么用？",
                "tag": "",
                "content": "谷歌辅助邮箱怎么用？",
                "url": ""
            },
            {
                "icon": "",
                "action": "openWindow",
                "text": "【脸书】双重验证教程",
                "tag": "",
                "content": "",
                "url": "//docs.qq.com/doc/DRVhySEtIT3hXQUVm"
            },
            {
                "icon": "",
                "action": "openWindow",
                "text": "【推特】确认码教程",
                "tag": "",
                "content": "",
                "url": "//docs.qq.com/doc/DRVhkSU9nZXR6Z1pz"
            },
            {
                "icon": "",
                "action": "openWindow",
                "text": "【推特】双重验证教程",
                "tag": "",
                "content": "",
                "url": "//docs.qq.com/doc/DRWR2b0NMZnhTZG9H"
            },
            {
                "icon": "",
                "action": "openWindow",
                "text": "【Ins】双重验证教程",
                "tag": "",
                "content": "",
                "url": "//docs.qq.com/doc/DRU5ocXFGb0pzRFR1"
            }
        ]
    }
}`
		b.db.Put("keyword", []byte("default"), []byte(dal))

		img := `{"code":"image","data":{"picUrl":"//gw.alicdn.com/tfs/TB1j2Y3xUY1gK0jSZFMXXaWcVXa-602-337.png"}}`

		b.db.Put("keyword", []byte("img"), []byte(img))
	}
}

func (b Keyword) Get(key string) any {
	msg, _ := b.db.Get("keyword", []byte(key))

	return string(msg)
}

func (b Keyword) GetList() map[string]string {
	result, err := b.db.GetAllKeysAndValues("keyword")
	if err != nil {
		return map[string]string{}
	}

	delete(result, "default")
	return result
}
