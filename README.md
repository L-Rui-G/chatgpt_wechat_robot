<div align="center">
<h1>ChatGPT-DreamStudio WeChat Robot</h1>
<p>  🎨基于GO语言实现的微信聊天和图片生成机器人🎨 </p>
<img src="https://camo.githubusercontent.com/82291b0fe831bfc6781e07fc5090cbd0a8b912bb8b8d4fec0696c881834f81ac/68747470733a2f2f70726f626f742e6d656469612f394575424971676170492e676966" width="800"  height="3">
</div><div align="left"></div>
个人微信接入ChatGPT，实现和GPT机器人互动聊天，同时支持基于文本生成图像。支持私聊回复和群聊艾特回复。

### 实现功能

* GPT机器人模型热度可配置
* 支持gpt3.5,gpt4模型
* 提问增加上下文&指令清空上下文
* DreamStudio图像生成模型参数可配置
* 可设定图像生成触发指令
* 机器人私聊回复&机器人群聊@回复
* 好友添加自动通过可配置

### 实现机制
1. 利用微信A作为机器人扫码登录程序模拟的微信电脑端，程序后端调用API接口进行文本回复和图片生成。其他微信账号与微信A聊天实现微信个人机器人功能。基于[openwechat](https://github.com/eatmoreapple/openwechat)开源仓库实现

2. 基于openai官网提供的GPT API，实现文本交互功能，每个新账号前三个月有18美元免费额度。
3. 基于stability官网提供的DreamStudio API,实现图像生成功能，每个账号注册送500张图像生成免费额度，玩玩基本够用了，不够的话10$可购买5000张。

> GPT的[官方文档](https://beta.openai.com/docs/models/overview)和详细[参数示例](https://beta.openai.com/examples) 。
>
> DreamStudio的[官方文档](https://platform.stability.ai/docs/getting-started)和详细[参数示例](https://platform.stability.ai/rest-api#tag/v1generation/operation/textToImage) 。

### 使用前提

* 有openai账号，并且创建好api_key，注册事项可以参考[此文章](https://juejin.cn/post/7173447848292253704) 。
* 有dreamstudio.ai账号，并且创建好api_key
* 微信必须实名认证。最好用小号

### 注意事项

* 项目仅供娱乐，滥用可能有微信封禁的风险，请勿用于商业用途。
* 请注意收发敏感信息，本项目不做信息过滤。
* dreamstudio图像生成 仅对英文的支持比较好

## 结果展示

<img src="https://blog-1257904201.cos.ap-shanghai.myqcloud.com/imgScreenshot_20230324_181306_edit_70435407551751.jpg.jpg" alt="Screenshot_20230324_181306_edit_70435407551751.jpg" style="zoom:50%;" />

## docker运行

使用docker快速运行本项目。

#### 1.基于配置文件挂载运行(推荐)

```sh
# 1. 创建目录
$ mkdir -p /data/openai
$ cd /data/openai
# 2. 创建配置文件
$ touch config.json
# 3. 编辑配置文件 ...  配置内容粘贴下文 【配置说明】并按需修改
$ vim config.json
# 4. 拉取镜像
$ docker run -dti --name wechatbot -v /data/openai/config.json:/app/config.json  yinqishuo/wechatbot:latest
# 5. 进入容器内部，打开日志文件扫码登陆
$ docker exec -it wechatbot bash 
$ tail -f -n 50 /app/run.log 

# 操作出错后删除容器的操作
$ docker stop wechatbot
$ docker remove wechatbot

# 退出容器
$ exit
```

其中配置文件参考下边的配置文件说明。 

#### 2.基于环境变量运行

```sh
# 运行项目，环境变量参考下方配置说明
$ docker run -itd --name wechatbot --restart=always \
 -e GPTAPIKEY=换成你的GPT key \
 -e AUTO_PASS=false \
 -e SESSION_TIMEOUT=60s \
 -e PROXY_URL=false\
 -e PROXY_URL=http://localhost:1080\
 -e MODEL=gpt-3.5-turbo \
 -e MAX_TOKENS=512 \
 -e TEMPREATURE=0.9 \
 -e REPLY_PREFIX=我是来自机器人回复: \
 -e SESSION_CLEAR_TOKEN=下个问题 \
 -e DREAMSTDIO_APIKEY=换成你的dreamstudio key \
 -e ENGINE_ID=stable-diffusion-v1-5 \
 -e PICTURE_WIDTH=512 \
 -e PICTURE_HEIGHT=512 \
 -e STEPS=30 \
 -e CFG_SCALE= 7 \
 -e PICTURE_TOKEN=生成图片 \
 yinqishuo/wechatbot:latest

#进入容器内部，打开日志文件扫码登陆
$ docker exec -it wechatbot bash 
$ tail -f -n 50 /app/run.log 

# 退出容器
$ exit
```

运行命令中映射的配置文件参考下边的配置文件说明。

#### 3.配置说明

模板：

```json
{
  "gpt_api_key": "你的gpt api key",
  "auto_pass": true,
  "proxy": false,
  "proxy_url": "http://localhost:1080",
  "session_timeout": 60,
  "max_tokens": 1024,
  "model": "gpt-3.5-turbo",
  "temperature": 1,
  "reply_prefix": "来自机器人回复：",
  "session_clear_token": "我要问下一个问题了",

  "dreamstdio_api_key":"你的dreamstdio账号api_key",
  "engine_id":"stable-diffusion-v1-5",
  "picture_width":512,
  "picture_height":512,
  "steps":30,
  "cfg_scale":7,
  "picture_token":"生成图片"
}
```

参数说明：

```
"gpt_api_key":						# openai账号里设置的api_key
"auto_pass":# 是否自动通过好友添加
"proxy": false,#是否添加代理
  "proxy_url": "http://localhost:1080",#代理的地址和端口
  "session_timeout": 60,            # 会话超时时间，默认60秒，单位秒，在会话时间内所有发送给机器人的信息会作为上下文
  "max_tokens": 1024,               # GPT响应字符数，最大2048，默认值512。会影响接口响应速度，字符越大响应越慢
  "model": "gpt-3.5-turbo",      # GPT选用模型，默认gpt-3.5-turbo，具体选项参考官网训练场
  "temperature": 1,                 # GPT热度，0到1，默认0.9，数字越大创造力越强，但更偏离训练事实，越低越接近训练事实
  "reply_prefix": "来自机器人回复：", # 私聊回复前缀
  "session_clear_token": "清空会话"  # 会话清空口令，默认`下个问题`
  "dreamstdio_api_key":"你的dreamstdio账号api_key",     #dreamstdio账号的api_key
  "engine_id":"stable-diffusion-v1-5",     			  #dreamstdio模型的名称
  "picture_width":512,								  #生成图片的宽度，长度，默认512*512
  "picture_height":512,								#要求为64的倍数，且>=128,尺寸越大消耗的credits越多
  "steps":30,										#代表模型的渲染步数，越高图片越精细，所需的渲染时间也越长，默认为30，数值越大消耗的credits越多；
  "cfg_scale":7,									#表示生成图像与文本提示的相似度，越高越像
  "picture_token":"生成图片"						  #生成图像的触发口令
```

## 源码运行

适合了解go语言编程并想进行源码修改的同学，

````shell
# 获取项目
$ git clone https://github.com/yinqishuo/chatgpt-dreamstudio_wechat_robot

# 进入项目目录
$ cd chatgpt_wechat_robot

# 复制配置文件
$ cp config.dev.json config.json

# 添加依赖
$ go mod tidy

# 启动项目
$ go run main.go

# 若想编译为可执行文件
$ go run main.go

# 若想打包成docker镜像，需安装Docker ，建议在linux环境下打包，镜像名称为wechatbot
$ make docker

# 执行镜像，步骤如上
````

### 常见问题

> 如无法登录`login error: write storage.json: bad file descriptor`
> 删除掉storage.json文件重新登录。

> 如无法登录`login error: wechat network error: Get "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage": 301 response missing Location header`
> 一般是微信登录权限问题，先确保PC端能否正常登录。

> 其他无法登录问题
> 尝试删除掉storage.json文件，结束进程(linux一般是kill -9 进程id)之后重启程序，重新扫码登录。
> 如果为docket部署，Supervisord进程管理工具会自动重启程序。

> 机器人一直答非所问
> 可能因为上下文累积过多。切换不同问题时，发送指令：启动时配置的`session_clear_token`字段。会清空上下文

https://link.zhihu.com/?target=https%3A//github.com/Maks-s/sd-akashic

### 图像生成技巧

想让 AI 图像生成器创作出精准高质的图像，填写准确合适的文本提示词十分重要。

1. 不要只输入一个简单的词语（raw prompt），如 Panda（熊猫）、A warrior（战士）等，这样生成的图像会缺少美感和艺术性。
2. 使用风格提示词能让图像更具艺术性。在提示词中加入艺术风格的关键词，如 Realistic（写实）、Oil painting（油画）、Pencil drawing（铅笔画）、Concept art（概念艺术）等；此外写实风格的提示词有多重表达形式，如[ a photo of + raw prompt ]、[ a photograph of + raw prompt ]、[ raw prompt,hyperrealistic ]、[ raw prompt,realistic ]。
3. 使用艺术家名称让风格更具像或保持风格一致。比如想表现抽象艺术，可以使用[made by Pablo Picassoa]或者 [ raw prompt,Picassoa]。还可以同时输入多名艺术家，效果会更加有趣。
4. 最终修饰词。在文本末尾加上的一个修饰词，使图像更符合你想要的效果。比如想要逼真的灯光，可以加上“Unreal Engine”，展现精密细节加上“4K”或“8K”，想要更有艺术性可以加上“trending on artstation”等。

**① Stable Diffusion Artist Studies**

网址： [https://proximacentaurib.notion.site/e2537cbf42c34b7e9a9a4126f81dfd0d](https://link.zhihu.com/?target=https%3A//proximacentaurib.notion.site/e2537cbf42c34b7e9a9a4126f81dfd0d)

一个由国外网友收集建立的艺术家风格概览表，找到你喜欢的风格后在自己的提示词中加上对应艺术家的名字，就能生成类似风格的图片。

**② Stable Diffusion prompting cheatsheet**

网址： [https://moritz.pm/posts/parameters](https://link.zhihu.com/?target=https%3A//moritz.pm/posts/parameters)

一个简短的提示词列表，里面列举了如果你想要实现 3D、精致细节、光照、大环境等效果，应该使用哪些关键词。

**③ Stable Diffusion Akashic Records**

网址： [https://github.com/Maks-s/sd-akashic](https://link.zhihu.com/?target=https%3A//github.com/Maks-s/sd-akashic)

一个专业的研究资料库，收集了关于模型原理、艺术风格、提示词、使用技巧和其他有用的工具，适合想深入了解文本-图像扩散模型的人阅读。

> 本段内容复制于[AI绘画神器：DreamStudio - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/557665226)



### 友情提示

本项目是 fork 他人的项目来进行学习和使用，请勿商用，可以下载下来做自定义的功能。
项目基于[eatmoreapple](https://github.com/eatmoreapple)/**[openwechat](https://github.com/eatmoreapple/openwechat)** 、[ZYallers](https://github.com/ZYallers)/**[chatgpt_wechat_robot](https://github.com/ZYallers/chatgpt_wechat_robot)** 、[qingconglaixueit](https://github.com/qingconglaixueit)/**[wechatbot](https://github.com/qingconglaixueit/wechatbot)**开发。  