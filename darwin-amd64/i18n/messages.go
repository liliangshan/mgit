package i18n

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// Messages 存储所有语言的消息
var Messages map[string]map[string]string

// CurrentLang 当前语言
var CurrentLang string

// 初始化语言包
func init() {
	Messages = make(map[string]map[string]string)

	// 加载所有语言文件
	loadLanguage("zh-CN") // 简体中文
	loadLanguage("zh-TW") // 繁体中文
	loadLanguage("en-US") // 英语
	loadLanguage("ja-JP") // 日语
	loadLanguage("ko-KR") // 韩语
	loadLanguage("fr-FR") // 法语

	// 默认使用简体中文
	CurrentLang = "zh-CN"
}

// 获取语言文件的绝对路径
func getI18nPath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exePath := filepath.Dir(ex)

	// 如果环境变量不存在，使用相对路径
	return filepath.Join(exePath, "i18n")
}

// 加载语言文件
func loadLanguage(lang string) {
	// 获取语言文件的绝对路径
	i18nPath := getI18nPath()
	filePath := filepath.Join(i18nPath, lang+".json")

	file, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	var messages map[string]string
	if err := json.Unmarshal(file, &messages); err != nil {
		return
	}

	Messages[lang] = messages
}

// T 翻译消息
func T(key string) string {
	if messages, ok := Messages[CurrentLang]; ok {
		if msg, ok := messages[key]; ok {
			return msg
		}
	}
	return key
}

// SetLanguage 设置当前语言
func SetLanguage(lang string) {
	if _, ok := Messages[lang]; ok {
		CurrentLang = lang
		// 重新加载语言文件以确保最新内容
		loadLanguage(lang)
	}
}
