package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/qingconglaixueit/wechatbot/pkg/logger"
)

// Configuration 项目配置
type Configuration struct {
	// gpt apikey
	GPTApiKey string `json:"gpt_api_key"`
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`
	// 是否通过代理
	Proxy bool `json:"proxy"`
	// 代理地址
	ProxyUrl string `json:"proxy_url"`
	// 会话超时时间
	SessionTimeout time.Duration `json:"session_timeout"`
	// GPT请求最大字符数
	MaxTokens int `json:"max_tokens"`
	// GPT模型
	Model string `json:"model"`
	// 热度
	Temperature float32 `json:"temperature"`
	// 回复前缀
	ReplyPrefix string `json:"reply_prefix"`
	// 清空会话口令
	SessionClearToken string `json:"session_clear_token"`

	// dreamstdio apikey
	DreamStdioApiKey string `json:"dreamstdio_api_key"`
	// dreamstdio模型名称
	EngineId string `json:"engine_id"`
	// 图像生成的高度
	PicWidth uint `json:"picture_width"`
	// 图像生成的高度
	PicHeight uint `json:"picture_height"`
	// 图像生成迭代次数
	Steps uint `json:"steps"`
	// 图像生成系数
	CfgScale uint `json:"cfg_scale"`
	// 图像生成识别指令
	PictureToken string `json:"picture_token"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 给配置赋默认值
		config = &Configuration{
			AutoPass:          false,
			SessionTimeout:    60,
			Proxy:             false,
			ProxyUrl:          "http://localhost:1080",
			MaxTokens:         512,
			Model:             "text-davinci-003",
			Temperature:       0.9,
			SessionClearToken: "下个问题",
			EngineId:          "stable-diffusion-v1-5",
			PicWidth:          512,
			PicHeight:         512,
			Steps:             30,
			CfgScale:          7,
			PictureToken:      "生成图片",
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				logger.Danger(fmt.Sprintf("open config error: %v", err))
				return
			}
			defer f.Close()
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				logger.Danger(fmt.Sprintf("decode config error: %v", err))
				return
			}
		}
		// 有环境变量使用环境变量
		GPTApiKey := os.Getenv("GPTAPIKEY")
		AutoPass := os.Getenv("AUTO_PASS")
		Proxy := os.Getenv("PROXY")
		ProxyUrl := os.Getenv("PROXY_URL")
		SessionTimeout := os.Getenv("SESSION_TIMEOUT")
		Model := os.Getenv("MODEL")
		MaxTokens := os.Getenv("MAX_TOKENS")
		Temperature := os.Getenv("TEMPREATURE")
		ReplyPrefix := os.Getenv("REPLY_PREFIX")
		SessionClearToken := os.Getenv("SESSION_CLEAR_TOKEN")

		DreamStdioApiKey := os.Getenv("DREAMSTDIO_APIKEY")
		EngineId := os.Getenv("ENGINE_ID")
		PicWidth := os.Getenv("PICTURE_WIDTH")
		PicHeight := os.Getenv("PICTURE_HEIGHT")
		Steps := os.Getenv("STEPS")
		CfgScale := os.Getenv("CFG_SCALE")
		PictureToken := os.Getenv("PICTURE_TOKEN")
		if GPTApiKey != "" {
			config.GPTApiKey = GPTApiKey
		}
		if AutoPass == "true" {
			config.AutoPass = true
		}
		if Proxy == "true" {
			config.Proxy = true
		}
		if ProxyUrl != "" {
			config.ProxyUrl = ProxyUrl
		}
		if SessionTimeout != "" {
			duration, err := time.ParseDuration(SessionTimeout)
			if err != nil {
				logger.Danger(fmt.Sprintf("config session timeout error: %v, get is %v", err, SessionTimeout))
				return
			}
			config.SessionTimeout = duration
		}
		if Model != "" {
			config.Model = Model
		}
		if MaxTokens != "" {
			max, err := strconv.Atoi(MaxTokens)
			if err != nil {
				logger.Danger(fmt.Sprintf("config max tokens error: %v ,get is %v", err, MaxTokens))
				return
			}
			config.MaxTokens = int(max)
		}
		if Temperature != "" {
			temp, err := strconv.ParseFloat(Temperature, 32)
			if err != nil {
				logger.Danger(fmt.Sprintf("config temperature error: %v, get is %v", err, Temperature))
				return
			}
			config.Temperature = float32(temp)
		}
		if ReplyPrefix != "" {
			config.ReplyPrefix = ReplyPrefix
		}
		if SessionClearToken != "" {
			config.SessionClearToken = SessionClearToken
		}
		if DreamStdioApiKey != "" {
			config.DreamStdioApiKey = DreamStdioApiKey
		}
		if EngineId != "" {
			config.EngineId = EngineId
		}
		if PicWidth != "" {
			width, err := strconv.Atoi(PicWidth)
			if err != nil {
				logger.Danger(fmt.Sprintf("config width  error: %v ,get is %v", err, PicWidth))
				return
			}
			config.PicWidth = uint(width)
		}
		if PicHeight != "" {
			height, err := strconv.Atoi(PicHeight)
			if err != nil {
				logger.Danger(fmt.Sprintf("config height error: %v ,get is %v", err, PicHeight))
				return
			}
			config.PicHeight = uint(height)
		}
		if Steps != "" {
			steps, err := strconv.Atoi(Steps)
			if err != nil {
				logger.Danger(fmt.Sprintf("config steps  error: %v ,get is %v", err, Steps))
				return
			}
			config.Steps = uint(steps)
		}
		if CfgScale != "" {
			cfg_scale, err := strconv.Atoi(CfgScale)
			if err != nil {
				logger.Danger(fmt.Sprintf("config cfg_scale  error: %v ,get is %v", err, CfgScale))
				return
			}
			config.CfgScale = uint(cfg_scale)
		}
		if PictureToken != "" {
			config.PictureToken = PictureToken
		}

	})
	if config.GPTApiKey == "" {
		logger.Danger("config error: GPTapi key required")
	}
	if config.DreamStdioApiKey == "" {
		logger.Danger("config error: DreamStdioApi key required")
	}
	return config
}
