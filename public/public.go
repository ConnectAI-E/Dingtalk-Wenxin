package public

import (
	"github.com/ConnectAI-E/Dingtalk-Wenxin/config"
	"github.com/ConnectAI-E/Dingtalk-Wenxin/pkg/cache"
	"github.com/ConnectAI-E/Dingtalk-Wenxin/pkg/db"
	"github.com/ConnectAI-E/Dingtalk-Wenxin/pkg/dingbot"
)

var UserService cache.UserServiceInterface
var Config *config.Configuration
var Prompt *[]config.Prompt
var DingTalkClientManager dingbot.DingTalkClientManagerInterface

const DingTalkClientIdKeyName = "DingTalkClientId"

func InitSvc() {
	// 加载配置
	Config = config.LoadConfig()
	// 加载prompt
	Prompt = config.LoadPrompt()
	// 初始化缓存
	UserService = cache.NewUserService()
	// 初始化钉钉开放平台的客户端，用于访问上传图片等能力
	DingTalkClientManager = dingbot.NewDingTalkClientManager(Config)
	// 初始化数据库
	db.InitDB()
}
