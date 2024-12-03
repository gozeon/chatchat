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

		b.db.Put("keyword", []byte("图片"), []byte(`{"code":"image","data":{"picUrl":"//gw.alicdn.com/tfs/TB1j2Y3xUY1gK0jSZFMXXaWcVXa-602-337.png"}}`))
		b.db.Put("keyword", []byte("视频"), []byte(`{
    "code": "video",
    "data": {
        "url": "//alime-channel.oss-cn-zhangjiakou.aliyuncs.com/customer-upload/1715943072999_d6adce42366a4eabb773d1251ca224ff.mov?Expires=1716375078&OSSAccessKeyId=LTAI5tApwLpSAsD6RjcwxwKJ&Signature=jdsuqZvN44J9x9KUzU6CXDCVUuI%3D"
    }
}`))
		b.db.Put("keyword", []byte("轮播"), []byte(`{
    "code": "Banners",
    "data": {
        "list": [
            {
                "image": "//alime-base-pic.oss-cn-zhangjiakou.aliyuncs.com/channel/1218901/1716194632902_447426bc77f3466f97faccc0849038a0.png",
                "link": "//www.taobao.com"
            }
        ]
    },
    "createdAt": 1716204911810,
    "position": "left",
    "hasTime": true
}`))
		b.db.Put("keyword", []byte("所有组件"), []byte(`[
    {
        "code": "system",
        "data": {
            "text": "智能助理进入对话，为您服务"
        }
    },
    {
        "code": "text",
        "data": {
            "text": "智能助理为您服务，请问有什么可以帮您？"
        }
    },
    {
        "code": "image",
        "data": {
            "picUrl": "//gw.alicdn.com/tfs/TB1j2Y3xUY1gK0jSZFMXXaWcVXa-602-337.png"
        }
    },
    {
        "code": "mix",
        "data": {
            "warpCard": false,
            "list": [
                {
                    "code": "text",
                    "data": {
                        "text": "文本内容"
                    }
                },
                {
                    "code": "image",
                    "data": {
                        "picUrl": "//gw.alicdn.com/tfs/TB1j2Y3xUY1gK0jSZFMXXaWcVXa-602-337.png"
                    }
                }
            ],
            "suffixList": [
                {
                    "code": "text",
                    "data": {
                        "text": "追加内容"
                    }
                }
            ]
        }
    },
    {
        "code": "promotion",
        "data": {
            "array": [
                {
                    "image": "//alime-base.oss-cn-beijing.aliyuncs.com/avatar/alime-base.oss-cn-beijing-internal.aliyuncs.com1569811067816-首页推荐卡底图（售前）.jpg",
                    "toggle": "//gw.alicdn.com/tfs/TB1D79ZXAL0gK0jSZFtXXXQCXXa-100-100.png",
                    "type": "recommend",
                    "list": [
                        {
                            "title": "收到商品不新鲜怎么办？",
                            "hot": true,
                            "content": "收到商品不新鲜怎么办？"
                        },
                        {
                            "title": "怎么改配送时间/地址/电话？",
                            "hot": true,
                            "content": "配送时间/地址/电话错了，怎么修改"
                        },
                        {
                            "title": "我的订单什么时间配送",
                            "content": "我的订单什么时间配送"
                        },
                        {
                            "title": "已下单，还能临时加/减商品吗？",
                            "content": "已下单，还能临时加/减商品吗？"
                        }
                    ]
                },
                {
                    "image": "//gw.alicdn.com/tfs/TB114P3XHY1gK0jSZTEXXXDQVXa-400-372.jpg",
                    "action": "send",
                    "text": "点此学习美食做法",
                    "type": "default",
                    "title": "热门菜谱",
                    "params": {
                        "content": "热门菜谱"
                    }
                },
                {
                    "image": "//gw.alicdn.com/tfs/TB1rsT0Xxv1gK0jSZFFXXb0sXXa-400-358.jpg",
                    "action": "send",
                    "text": "看看你家的天气吧",
                    "type": "default",
                    "title": "天气查询",
                    "params": {
                        "content": "天气查询"
                    }
                }
            ]
        }
    },
    {
        "code": "recommend",
        "data": {
            "list": [
                {
                    "title": "如何办理准生证？"
                },
                {
                    "title": "生育登记成功后在哪里看"
                },
                {
                    "title": "链接形式",
                    "url": "//www.baidu.com/"
                }
            ]
        }
    },
    {
        "code": "foundation-model",
        "data": {
            "content": "对话机器人 POC 环境地址是： http://poc-console.alixiaomi.com:8081/",
            "sentenceList": [
                {
                    "content": "如需进一步详情或具体价格，请告知。"
                }
            ],
            "relatedQuestionList": [
                {
                    "content": "测试关联问1"
                },
                {
                    "content": "测试关联问2"
                }
            ],
            "streamEnd": true
        },
        "meta": {
            "evaluable": true
        }
    },
    {
        "code": "slot",
        "data": {
            "hideShortcuts": true,
            "list": [
                {
                    "title": "社保卡"
                },
                {
                    "title": "参保证明打印"
                },
                {
                    "title": "失业保险"
                },
                {
                    "title": "工伤保险"
                },
                {
                    "title": "养老保险"
                }
            ]
        }
    },
    {
        "position": "right",
        "code": "text",
        "data": {
            "text": "我要办理公积金贷款"
        }
    },
    {
        "hasTime": true,
        "code": "knowledge",
        "data": {
            "text": "办理公积金贷款的流程主要包括以下步骤：<br>\n              提交申请材料。需要提供包括住房公积金联名卡、有效身份证件、婚姻状况证明、购房合同、首付款凭证、房产证等。<br>\n              申请条件确认。确保满足公积金正常缴存12个月以上，公积金账户状态正常，无未结清的公积金贷款，有稳定的收入和信用记录。<br>\n              贷款审批。银行或公积金中心会对提交的材料进行审核，包括贷款额度的评估。<br>\n              签订合同。审批通过后，与银行签订借款合同。<br>\n              等待放款。银行会根据合同条款放款。<br>\n              此外，应确保所有资料齐全、准确，以满足公积金贷款的条件和要求。具体流程和所需材料可能因地区而有所不同，建议咨询当地的住房公积金管理中心或银行以获取详细信息",
            "childKnowledges": [
                {
                    "showTitle": "公积金贷款准备材料"
                },
                {
                    "showTitle": "银行贷款办理步骤"
                },
                {
                    "showTitle": "银行贷款准备材料"
                }
            ]
        },
        "meta": {
            "evaluable": true
        }
    },
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
    },
    {
        "code": "sub-business-list",
        "data": {
            "array": [
                {
                    "icon": "//alime-base-pic.oss-cn-zhangjiakou.aliyuncs.com/channel/1159612/1709100588192_a8075393151b4fabbdc4a53154c76522.png",
                    "title": "换货申请",
                    "list": [
                        {
                            "title": "点击查看教程title",
                            "type": "sendText",
                            "content": "点击查看教程content"
                        }
                    ]
                },
                {
                    "icon": "//alime-base-pic.oss-cn-zhangjiakou.aliyuncs.com/channel/1159612/1709100640240_a7e3559eb3d547d88317590162986b06.png",
                    "title": "快捷查单",
                    "list": [
                        {
                            "link": "//docs.qq.com/doc/DRVBUR2xMR2NmYmNz",
                            "title": "点击查看教程",
                            "type": "url",
                            "content": ""
                        }
                    ]
                }
            ]
        }
    },
    {
        "code": "image",
        "data": {
            "picUrl": "//alime-base-pic.oss-cn-zhangjiakou.aliyuncs.com/channel/1218901/1716194632902_447426bc77f3466f97faccc0849038a0.png"
        }
    },
    {
        "code": "video",
        "data": {
            "url": "//alime-channel.oss-cn-zhangjiakou.aliyuncs.com/customer-upload/1715943072999_d6adce42366a4eabb773d1251ca224ff.mov?Expires=1716375078&OSSAccessKeyId=LTAI5tApwLpSAsD6RjcwxwKJ&Signature=jdsuqZvN44J9x9KUzU6CXDCVUuI%3D"
        },
        "position": "left",
        "createdAt": 1715943078843,
        "hasTime": false
    },
    {
        "code": "firstmsg",
        "data": {
            "operations": [],
            "briefs": [
                {
                    "title": "如何做好冬季疫情防范（示例）",
                    "subContent": [
                        {
                            "title": "疫苗接种点查询（示例）title",
                            "content": "疫苗接种点查询（示例）content"
                        }
                    ]
                },
                {
                    "title": "本地卫健委电话查询（示例）",
                    "subContent": []
                },
                {
                    "title": "新冠症状表现查询（示例）",
                    "subContent": []
                }
            ]
        },
        "createdAt": 1716204911810,
        "position": "left",
        "hasTime": false
    },
    {
        "code": "sub-skill-list",
        "data": {
            "array": [
                {
                    "image": "",
                    "action": "send",
                    "text": "示例文本text",
                    "params": {
                        "content": "优惠示例content"
                    },
                    "title": "优惠示例title"
                },
                {
                    "image": "",
                    "action": "send",
                    "text": "示例文本",
                    "params": {
                        "content": "评价示例"
                    },
                    "title": "评价示例"
                },
                {
                    "image": "",
                    "action": "send",
                    "text": "示例文本",
                    "params": {
                        "content": "极速示例"
                    },
                    "title": "极速示例"
                }
            ]
        },
        "createdAt": 1716204911810,
        "position": "left",
        "hasTime": false
    },
    {
        "code": "Banners",
        "data": {
            "list": [
                {
                    "image": "//alime-base-pic.oss-cn-zhangjiakou.aliyuncs.com/channel/1218901/1716194632902_447426bc77f3466f97faccc0849038a0.png",
                    "link": "//www.taobao.com",
                    "startTime": "",
                    "endTime": ""
                }
            ]
        },
        "createdAt": 1716204911810,
        "position": "left",
        "hasTime": true
    },
    {
        "code": "qa-collapse-list",
        "data": {
            "list": [
                {
                    "title": "《知识管理规定》（示例）",
                    "content": "知识管理的暂行规定详细介绍，400字以内（示例文案）"
                },
                {
                    "title": "知识运营最佳实践（示例）",
                    "content": "知识运营最佳实践详细介绍，400字以内（示例文案）"
                },
                {
                    "title": "智能对话机器人在企业服务场景的实践与应用（示例）",
                    "content": "智能对话机器人在企业服务场景的实践与应用详细介绍，400字以内（示例文案）"
                }
            ]
        }
    },
    {
        "code": "audio",
        "data": {
            "src": "//beebot-knowledgecloud-aliyun-public-cn-shanghai.oss-cn-shanghai.aliyuncs.com/template/6.4.0/chat-sdk/audio.mp3"
        }
    },
    {
        "code": "video",
        "data": {
            "url": "//cloud.video.taobao.com/play/u/3544775890/p/1/e/6/t/1/277895493609.mp4"
        }
    },
    {
        "code": "mix",
        "data": {
            "wrapCard": true,
            "list": [
                {
                    "code": "databot-card",
                    "data": "{\"cellList\":[[\"总花费\",\"产品名称\"],[\"40947.5\",\"云安全中心\"],[\"16665\",\"Web应用防火墙\"],[\"9036.97003540955489874\",\"文件存储NAS\"],[\"7449.669921875\",\"检索分析服务 Elasticsearch版\"],[\"7436.1597900390625\",\"云数据库 ClickHouse\"],[\"5965.540030241012575\",\"云监控\"],[\"5796.1499633789063\",\"微服务引擎\"],[\"1720\",\"数据库自治服务\"],[\"890.36000714823604027\",\"共享带宽\"],[\"405.369999991729855544\",\"数据库备份\"]],\"clickable\":true,\"drawType\":\"bar\",\"nextPage\":false,\"title\":\"(元数据：北控水务;)\"}"
                },
                {
                    "code": "answer-reference",
                    "data": {
                        "answerReference": {
                            "itemList": [
                                {
                                    "content": "{\"code\":\"llm-databot-card\",\"data\":{\"cellList\":[[\"产品名称\",\"总花费\"],[\"云安全中心\",\"40947.5\"],[\"Web应用防火墙\",\"16665\"],[\"文件存储NAS\",\"9036.97003540955489874\"],[\"检索分析服务 Elasticsearch版\",\"7449.669921875\"],[\"云数据库 ClickHouse\",\"7436.1597900390625\"],[\"云监控\",\"5965.540030241012575\"],[\"微服务引擎\",\"5796.1499633789063\"],[\"数据库自治服务\",\"1720\"],[\"共享带宽\",\"890.36000714823604027\"],[\"数据库备份\",\"405.369999991729855544\"]],\"clickable\":false,\"drawType\":\"table\",\"nextPage\":false}}",
                                    "contentType": "CARD_TEMPLATE",
                                    "dataSource": "TABLE",
                                    "number": 1,
                                    "referenceExt": {
                                        "docName": "北控水务"
                                    },
                                    "title": "北控水务"
                                }
                            ]
                        },
                        "cardTemplateContent": "{\"code\":\"llm-databot-card\",\"data\":\"{\\\"cellList\\\":[[\\\"总花费\\\",\\\"产品名称\\\"],[\\\"40947.5\\\",\\\"云安全中心\\\"],[\\\"16665\\\",\\\"Web应用防火墙\\\"],[\\\"9036.97003540955489874\\\",\\\"文件存储NAS\\\"],[\\\"7449.669921875\\\",\\\"检索分析服务 Elasticsearch版\\\"],[\\\"7436.1597900390625\\\",\\\"云数据库 ClickHouse\\\"],[\\\"5965.540030241012575\\\",\\\"云监控\\\"],[\\\"5796.1499633789063\\\",\\\"微服务引擎\\\"],[\\\"1720\\\",\\\"数据库自治服务\\\"],[\\\"890.36000714823604027\\\",\\\"共享带宽\\\"],[\\\"405.369999991729855544\\\",\\\"数据库备份\\\"]],\\\"clickable\\\":true,\\\"drawType\\\":\\\"bar\\\",\\\"nextPage\\\":false,\\\"title\\\":\\\"(元数据：北控水务;)\\\"}\"}",
                        "contentType": "CARD_TEMPLATE"
                    }
                }
            ]
        },
        "_id": "82633feb-6662-40c2-ba75-c24563b3eeff",
        "meta": {
            "evaluable": true,
            "chatUuid": "82633feb-6662-40c2-ba75-c24563b3eeff",
            "sid": "6de1e0d0198111ef8224b93b5ad938db"
        }
    }
]`))
		b.db.Put("keyword", []byte("不回复"), []byte(`null`))
		b.db.Put("keyword", []byte("文字"), []byte(`{
    "code": "text",
    "data": {
        "text": "智能助理为您服务，请问有什么可以帮您？"
    }
}`))
		b.db.Put("keyword", []byte("系统消息"), []byte(`{
    "code": "system",
    "data": {
        "text": "智能助理进入对话，为您服务"
    }
}`))
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
