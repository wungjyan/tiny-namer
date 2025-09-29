package highlight

import (
	"log"
	"regexp"
)

// HighlightPart 高亮部分
type HighlightPart struct {
	Text        string `json:"text"`
	IsHighlight bool   `json:"isHighlight"`
}

// HighlightRequest 高亮请求
type HighlightRequest struct {
	FileName   string `json:"fileName"`   // 完整文件名
	BaseName   string `json:"baseName"`   // 基础文件名（不含扩展名）
	Pattern    string `json:"pattern"`    // 匹配模式
	IgnoreExt  bool   `json:"ignoreExt"`  // 是否忽略扩展名
	IgnoreCase bool   `json:"ignoreCase"` // 是否忽略大小写
	Global     bool   `json:"global"`     // 是否全局匹配
}

// HighlightResult 高亮结果
type HighlightResult struct {
	Success bool            `json:"success"`
	Parts   []HighlightPart `json:"parts"`
	Message string          `json:"message"`
}

// Service 高亮服务
type Service struct{}

// NewService 创建高亮服务实例
func NewService() *Service {
	return &Service{}
}

// GetHighlightParts 获取高亮部分，返回匹配到的文本片段
func (s *Service) GetHighlightParts(req HighlightRequest) []string {
	if req.Pattern == "" {
		return []string{} // 空模式返回空切片
	}

	// 确定要处理的文本
	text := req.FileName
	if req.IgnoreExt && req.BaseName != "" {
		text = req.BaseName
	}

	// 构建正则表达式
	pattern := req.Pattern
	if req.IgnoreCase {
		pattern = "(?i)" + pattern
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return []string{} // 正则表达式编译失败返回空切片
	}

	var matches []string
	log.Println("正则表达式:", regex.String())
	log.Println("要处理的文本:", text)

	if req.Global {
		// 全局匹配所有
		matches = regex.FindAllString(text, -1)
	} else {
		// 只匹配第一个
		match := regex.FindString(text)
		if match != "" {
			matches = []string{match}
		}
	}

	return matches
}
