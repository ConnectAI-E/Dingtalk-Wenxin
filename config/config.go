package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ConnectAI-E/Dingtalk-Wenxin/pkg/logger"

	"gopkg.in/yaml.v2"
)

type Credential struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// Configuration 项目配置
type Configuration struct {
	// 日志级别，info或者debug
	LogLevel string `yaml:"log_level"`
	// 运行模式
	RunMode string `yaml:"run_mode"`
	// 百度应用的clientID
	BaiduClientID string `yaml:"baidu_client_id"`
	// 百度应用的clientSecret
	BaiduClientSecret string `yaml:"baidu_client_secret"`
	// 默认对话模式
	DefaultMode string `yaml:"default_mode"`
	// 会话超时时间
	SessionTimeout time.Duration `yaml:"session_timeout"`
	// 用户单日最大请求次数
	MaxRequest int `yaml:"max_request"`
	// 指定服务启动端口，默认为 8090
	Port string `yaml:"port"`
	// 限定对话类型 0：不限 1：单聊 2：群聊
	ChatType string `yaml:"chat_type"`
	// 哪些群组可以进行对话
	AllowGroups []string `yaml:"allow_groups"`
	// 哪些outgoing群组可以进行对话
	AllowOutgoingGroups []string `yaml:"allow_outgoing_groups"`
	// 哪些用户可以进行对话
	AllowUsers []string `yaml:"allow_users"`
	// 哪些用户不可以进行对话
	DenyUsers []string `yaml:"deny_users"`
	// 哪些Vip用户可以进行无限对话
	VipUsers []string `yaml:"vip_users"`
	// 指定哪些人为此系统的管理员，必须指定，否则所有人都是
	AdminUsers []string `yaml:"admin_users"`
	// 敏感词，提问时触发，则不允许提问，回答的内容中触发，则以 🚫 代替
	SensitiveWords []string `yaml:"sensitive_words"`
	// 自定义帮助信息
	Help string `yaml:"help"`
	// 钉钉应用鉴权凭据
	Credentials []Credential `yaml:"credentials"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		data, err := ioutil.ReadFile("config.yml")
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatal(err)
		}

		// 如果环境变量有配置，读取环境变量
		logLevel := os.Getenv("LOG_LEVEL")
		if logLevel != "" {
			config.LogLevel = logLevel
		}
		runMode := os.Getenv("RUN_MODE")
		if runMode != "" {
			config.RunMode = runMode
		}
		baiDuClientId := os.Getenv("BAIDU_CLIENT_ID")
		if baiDuClientId != "" {
			config.BaiduClientID = baiDuClientId
		}
		baiDuClientSecret := os.Getenv("BAIDU_CLIENT_SECRET")
		if baiDuClientSecret != "" {
			config.BaiduClientSecret = baiDuClientSecret
		}
		sessionTimeout := os.Getenv("SESSION_TIMEOUT")
		if sessionTimeout != "" {
			duration, err := strconv.ParseInt(sessionTimeout, 10, 64)
			if err != nil {
				logger.Fatal(fmt.Sprintf("config session timeout err: %v ,get is %v", err, sessionTimeout))
				return
			}
			config.SessionTimeout = time.Duration(duration) * time.Second
		} else {
			config.SessionTimeout = time.Duration(config.SessionTimeout) * time.Second
		}
		defaultMode := os.Getenv("DEFAULT_MODE")
		if defaultMode != "" {
			config.DefaultMode = defaultMode
		}
		maxRequest := os.Getenv("MAX_REQUEST")
		if maxRequest != "" {
			newMR, _ := strconv.Atoi(maxRequest)
			config.MaxRequest = newMR
		}
		port := os.Getenv("PORT")
		if port != "" {
			config.Port = port
		}
		chatType := os.Getenv("CHAT_TYPE")
		if chatType != "" {
			config.ChatType = chatType
		}
		allowGroups := os.Getenv("ALLOW_GROUPS")
		if allowGroups != "" {
			config.AllowGroups = strings.Split(allowGroups, ",")
		}
		allowOutgoingGroups := os.Getenv("ALLOW_OUTGOING_GROUPS")
		if allowOutgoingGroups != "" {
			config.AllowOutgoingGroups = strings.Split(allowOutgoingGroups, ",")
		}
		allowUsers := os.Getenv("ALLOW_USERS")
		if allowUsers != "" {
			config.AllowUsers = strings.Split(allowUsers, ",")
		}
		denyUsers := os.Getenv("DENY_USERS")
		if denyUsers != "" {
			config.DenyUsers = strings.Split(denyUsers, ",")
		}
		vipUsers := os.Getenv("VIP_USERS")
		if vipUsers != "" {
			config.VipUsers = strings.Split(vipUsers, ",")
		}
		adminUsers := os.Getenv("ADMIN_USERS")
		if adminUsers != "" {
			config.AdminUsers = strings.Split(adminUsers, ",")
		}
		sensitiveWords := os.Getenv("SENSITIVE_WORDS")
		if sensitiveWords != "" {
			config.SensitiveWords = strings.Split(sensitiveWords, ",")
		}
		help := os.Getenv("HELP")
		if help != "" {
			config.Help = help
		}
		credentials := os.Getenv("DINGTALK_CREDENTIALS")
		if credentials != "" {
			if config.Credentials == nil {
				config.Credentials = []Credential{}
			}
			for _, idSecret := range strings.Split(credentials, ",") {
				items := strings.SplitN(idSecret, ":", 2)
				if len(items) == 2 {
					config.Credentials = append(config.Credentials, Credential{ClientID: items[0], ClientSecret: items[1]})
				}
			}
		}

	})

	// 一些默认值
	if config.LogLevel == "" {
		config.LogLevel = "info"
	}
	if config.RunMode == "" {
		config.LogLevel = "http"
	}
	if config.DefaultMode == "" {
		config.DefaultMode = "单聊"
	}
	if config.Port == "" {
		config.Port = "8090"
	}
	if config.ChatType == "" {
		config.ChatType = "0"
	}
	return config
}
