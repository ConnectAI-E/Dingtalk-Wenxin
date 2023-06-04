## 前言

本项目可以助你将百度文心一言集成到钉钉群聊当中。

## 功能介绍

- 🚀 帮助菜单：通过发送 `帮助` 将看到帮助列表，[🖼 查看示例](#%E5%B8%AE%E5%8A%A9%E5%88%97%E8%A1%A8)
- 🥷 私聊：支持与机器人单独私聊(无需艾特)，[🖼 查看示例](#%E4%B8%8E%E6%9C%BA%E5%99%A8%E4%BA%BA%E7%A7%81%E8%81%8A)
- 💬 群聊：支持在群里艾特机器人进行对话
- 🙋 单聊模式：每次对话都是一次新的对话，没有历史聊天上下文联系
- 🗣 串聊模式：带上下文理解的对话模式
- 🎨 图片生成：通过发送 `#图片`关键字开头的内容进行生成图片，[🖼 查看示例](#%E7%94%9F%E6%88%90%E5%9B%BE%E7%89%87)
- 🎭 角色扮演：支持场景模式，通过 `#周报` 的方式触发内置prompt模板 [🖼 查看示例](#%E9%80%9A%E8%BF%87%E5%86%85%E7%BD%AEprompt%E8%81%8A%E5%A4%A9)
- 🧑‍💻 频率限制：通过配置指定，自定义单个用户单日最大对话次数
- 💵 余额查询：通过发送 `余额` 关键字查询当前key所剩额度，[🖼 查看示例](#%E6%9F%A5%E8%AF%A2%E4%BD%99%E9%A2%9D)
- 🔗 自定义api域名：通过配置指定，解决国内服务器无法直接访问openai的问题
- 🪜 添加代理：通过配置指定，通过给应用注入代理解决国内服务器无法访问的问题
- 👐 默认模式：支持自定义默认的聊天模式，通过配置化指定
- 📝 查询对话：通过发送`#查对话 username:xxx`查询xxx的对话历史，可在线预览，可下载到本地
- 👹 白名单机制：通过配置指定，支持指定群组名称和用户名称作为白名单，从而实现可控范围与机器人对话
- 💂‍♀️ 管理员机制：通过配置指定管理员，部分敏感操作，以及一些应用配置，管理员有权限进行操作
- ㊙️ 敏感词过滤：通过配置指定敏感词，提问时触发，则不允许提问，回答的内容中触发，则以 🚫 代替
- 🚇 stream模式：指定钉钉的stream模式，目前钉钉已全量开放该功能，项目也默认以此模式启动

## 使用前提

* 通过智能云创建应用，参考[文档](https://cloud.baidu.com/doc/WENXINWORKSHOP/s/flfmc9do2)，创建应用之后，获取到应用秘钥对。
* 在钉钉开发者后台创建应用，在应用的消息推送中选择stream模式。

## 使用教程

### 第一步，部署应用

#### docker部署

推荐你使用docker快速运行本项目。

```
第一种：基于环境变量运行
# 运行项目
$ docker run -itd --name chatgpt -p 8090:8090 \
  -v ./data:/app/data \
  -e LOG_LEVEL="info" -e BAIDU_CLIENT_ID=换成你的ID -e BAIDU_CLIENT_SECRET=换成你的key \
  -e DEFAULT_MODE="单聊" -e MAX_REQUEST=0 -e PORT=8090 -e CHAT_TYPE="0" \
  -e ALLOW_GROUPS=a,b -e ALLOW_OUTGOING_GROUPS=a,b -e ALLOW_USERS=a,b -e DENY_USERS=a,b -e VIP_USERS=a,b -e ADMIN_USERS=a,b \
  -e SENSITIVE_WORDS="aa,bb" -e RUN_MODE="stream" \
  -e DINGTALK_CREDENTIALS="your_client_id1:secret1,your_client_id2:secret2" \
  -e HELP="欢迎使用本工具\n\n你可以查看：[用户指南](https://github.com/ConnectAI-E/Dingtalk-Wenxin/blob/main/docs/userGuide.md)\n\n这是一个[开源项目](https://github.com/ConnectAI-E/Dingtalk-Wenxin/)
  ，觉得不错你可以来波素质三连."  \
  --restart=always  dockerproxy.com/eryajf/dingtalk-wenxin:latest
```

> 运行命令中映射的配置文件参考下边的[配置文件说明](#%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6%E8%AF%B4%E6%98%8E)。

- `📢 注意：`ALLOW_GROUPS,ALLOW_USERS,DENY_USERS,VIP_USERS,ADMIN_USERS 参数为数组，如果需要指定多个，可用英文逗号分割。outgoing机器人模式下这些参数无效。


```
第二种：基于配置文件挂载运行
# 复制配置文件，根据自己实际情况，调整配置里的内容
$ cp config.example.yml config.yml  # 其中 config.example.yml 从项目的根目录获取

# 运行项目
$ docker run -itd --name chatgpt -v `pwd`/config.yml:/app/config.yml --restart=always  dockerproxy.com/eryajf/dingtalk-wenxin:latest
```

其中配置文件参考下边的配置文件说明。

```
第三种：使用 docker compose 运行
$ wget https://raw.githubusercontent.com/ConnectAI-E/Dingtalk-Wenxin/main/docker-compose.yml

$ vim docker-compose.yml

$ docker compose up -d
```

部署完成之后，就可以在群里艾特机器人进行体验了。

#### 二进制部署


如果你想通过命令行直接部署，可以直接下载release中的[压缩包](https://github.com/ConnectAI-E/Dingtalk-Wenxin/releases) ，请根据自己系统以及架构选择合适的压缩包，下载之后直接解压运行。

下载之后，在本地解压，即可看到可执行程序，与配置文件：

```sh
$ tar xf dingtalk-wenxin-v0.0.4-darwin-arm64.tar.gz
$ cd dingtalk-wenxin-v0.0.4-darwin-arm64
$ cp config.example.yml  config.yml # 然后根据情况调整配置文件内容,宿主机如遇端口冲突,可通过调整config.yml中的port参数自定义服务端口
$ ./dingtalk-wenxin  # 直接运行

# 如果要守护在后台运行
$ nohup ./dingtalk-wenxin &> run.log &
$ tail -f run.log
```


### 第二步，创建机器人

#### 企业内部应用

创建步骤参考文档：[企业内部开发机器人](https://open.dingtalk.com/document/robots/enterprise-created-chatbot)，或者根据如下步骤进行配置。

1. 创建机器人。
    <details>
      <summary>🖼 点我查看示例图</summary>
      <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20221209_163616.png">
    </details>

   > `📢 注意1：`可能现在创建机器人的时候名字为`chatgpt`会被钉钉限制，请用其他名字命名。
   >
   > `📢 注意2：`第四步骤点击创建应用的时候，务必选择使用旧版，从而创建旧版机器人，如果选择新版，则机器人的功能集成在了应用当中的消息推送模块儿。

   步骤比较简单，这里就不赘述了。

2. 配置机器人回调接口。
    <details>
      <summary>🖼 点我查看示例图</summary>
      <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20221209_163652.png">
    </details>

   创建完毕之后，点击机器人开发管理，然后配置将要部署的服务所在服务器的出口IP，以及将要给服务配置的域名。

  > `📢 注意：` 添加消息接收地址的时候，务必确保服务在正常运行且可通过回调地址访问，否则保存时将会失败。
  >
  > `📢 如果提示：` 消息接收地址校验失败（请确保公网可访问该地址，如无有效SSL证书，可选择禁用证书校验），那么可以先输入一个`https://`，然后就能看到`禁用https`的选项了，选择禁用，然后再把地址改成`http`就好了。

3. 发布机器人。
    <details>
      <summary>🖼 点我查看示例图</summary>
      <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20221209_163709.png">
    </details>

   点击版本管理与发布，然后点击上线，这个时候就能在钉钉的群里中添加这个机器人了。

4. 群聊添加机器人。
    <details>
      <summary>🖼 点我查看示例图</summary>
      <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20221209_163724.png">
    </details>

## 亮点特色

### 与机器人私聊

`2023-03-08`补充，我发现也可以不在群里艾特机器人聊天，还可点击机器人，然后点击发消息，通过与机器人直接对话进行聊天：

> 由 [@Raytow](https://github.com/Raytow) 同学发现，在机器人自动生成的测试群里无法直接私聊机器人，在其他群里单独添加这个机器人，然后再点击就可以跟它私聊了。

<details>
    <summary>🖼 点我查看示例图</summary>
    <img src="https://user-images.githubusercontent.com/33259379/223607306-2ac836a2-7ce5-4a12-a16e-bec40b22d8d6.png">
</details>


### 帮助列表

> 艾特机器人发送空内容或者帮助，会返回帮助列表。

<details>
    <summary>🖼 点我查看示例图</summary>
    <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20230216_221253.png">
</details>

### 切换模式

> 发送指定关键字，可以切换不同的模式。

<details>
    <summary>🖼 点我查看示例图</summary>
    <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20230215_184655.png">
</details>

> 📢 注意：串聊模式下，群里每个人的聊天上下文是独立的。
> 📢 注意：默认对话模式为单聊，因此不必发送单聊即可进入单聊模式，而要进入串聊，则需要发送串聊关键字进行切换，当串聊内容超过最大限制的时候，你可以发送重置，然后再次进入串聊模式。


### 日常问题

<details>
    <summary>🖼 点我查看示例图</summary>
    <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20221209_163739.png">
</details>

### 通过内置prompt聊天

> 发送模板两个字，会返回当前内置支持的prompt列表。

<details>
    <summary>🖼 点我查看示例图</summary>
    <img src="https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20230323_152703.jpg">
</details>

> 如果你发现有比较优秀的prompt，欢迎PR。注意：一些与钉钉使用场景不是很匹配的，就不要提交了。


##  本地开发

```sh
# 获取项目
$ git clone https://github.com/ConnectAI-E/Dingtalk-Wenxin.git

# 进入项目目录
$ cd dingtalk-wenxin

# 复制配置文件，根据个人实际情况进行配置
$ cp config.example.yml config.yml

# 启动项目
$ go run main.go
```

## 配置文件说明

```yaml
# 应用的日志级别，info or debug
log_level: "info"
# 运行模式，http 或者 stream ，强烈建议你使用stream模式，通过此链接了解：https://open.dingtalk.com/document/isvapp/stream
run_mode: "stream"
# 百度应用中的client_id
baidu_client_id: ""
# 百度应用中的 client_secret
baidu_client_secret: ""
# 指定默认的对话模式，可根据实际需求进行自定义，如果不设置，默认为单聊，即无上下文关联的对话模式
default_mode: "单聊"
# 单人单日请求次数上限，默认为0，即不限制
max_request: 0
# 指定服务启动端口，默认为 8090，一般在二进制宿主机部署时，遇到端口冲突时使用，如果run_mode为stream模式，则可以忽略该配置项
port: "8090"
# 限定对话类型 0：不限 1：只能单聊 2：只能群聊
chat_type: "0"
# 哪些群组可以进行对话（仅在chat_type为0、2时有效），如果留空，则表示允许所有群组，如果要限制，则列表中写群ID（ConversationID）
# 群ID，可在群组中 @机器人 群ID 来查看日志获取，例如日志会输出：[🙋 企业内部机器人 在『测试』群的ConversationID为: "cidrabcdefgh1234567890AAAAA"]，获取后可填写该参数并重启程序
allow_groups: []
# 哪些普通群（使用outgoing机器人）可以进行对话，如果留空，则表示允许所有群组，如果要限制，则列表中写群ID（ConversationID）
# 群ID，可在群组中 @机器人 群ID 来查看日志获取，例如日志会输出：[🙋 outgoing机器人 在『测试』群的ConversationID为: "cidrabcdefgh1234567890AAAAA"]，获取后可填写该参数并重启程序
# 如果不想支持outgoing机器人功能，这里可以随意设置一个内部群组，例如：cidrabcdefgh1234567890AAAAA；或随意一个字符串，例如：disabled
# 建议该功能默认关闭：除非你必须要用到outgoing机器人
allow_outgoing_groups: []
# 以下 allow_users、deny_users、vip_users、admin_users 配置中填写的是用户的userid，outgoing机器人模式下不适用这些配置
# 比如 ["1301691029702722","1301691029702733"]，这个信息需要在钉钉管理后台的通讯录当中获取：https://oa.dingtalk.com/contacts.htm#/contacts
# 哪些用户可以进行对话，如果留空，则表示允许所有用户，如果要限制，则列表中写用户的userid
allow_users: []
# 哪些用户不可以进行对话，如果留空，则表示允许所有用户（如allow_user有配置，需满足相应条件），如果要限制，则列表中写用户的userid，黑名单优先级高于白名单
deny_users: []
# 哪些用户可以进行无限对话，如果留空，则表示只允许管理员（如max_request配置为0，则允许所有人）
# 如果要针对指定VIP用户放开限制（如max_request配置不为0），则列表中写用户的userid
vip_users: []
# 指定哪些人为此系统的管理员，如果留空，则表示没有人是管理员，如果要限制，则列表中写用户的userid
admin_users: []
# 敏感词，提问时触发，则不允许提问，回答的内容中触发，则以 🚫 代替
sensitive_words: []
# 帮助信息，放在配置文件，可供自定义
help: "这是自定义的帮助信息。"

# 钉钉应用鉴权凭据信息，支持多个应用。通过请求时候鉴权来识别是来自哪个机器人应用的消息
# 设置credentials 之后，即具备了访问钉钉平台绝大部分 OpenAPI 的能力；例如上传图片到钉钉平台，提升图片体验，结合 Stream 模式简化服务部署
# client_id 对应钉钉平台 AppKey/SuiteKey；client_secret 对应 AppSecret/SuiteSecret
credentials:
  -
    client_id: "put-your-client-id-here"
    client_secret: "put-your-client-secret-here"
```


## 进群交流

我创建了一个钉钉的交流群，欢迎进群交流。

![](https://cdn.staticaly.com/gh/eryajf/tu/main/img/image_20230405_191425.jpg)


## 赞赏

如果觉得这个项目对你有帮助，你可以请作者[喝杯咖啡 ☕️](https://wiki.eryajf.net/reward/)