package regex

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"tiny-namer/backend/file"
)

// ValidateResult 验证结果
type ValidateResult struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

// RenameRequest 重命名请求
type RenameRequest struct {
	FileInfo    file.FileInfo `json:"fileInfo"`
	Pattern     string        `json:"pattern"`
	Replacement string        `json:"replacement"`
	IgnoreExt   bool          `json:"ignoreExt"`
	IgnoreCase  bool          `json:"ignoreCase"`
	Global      bool          `json:"global"`
}

// RenameResult 重命名结果
type RenameResult struct {
	Success     bool   `json:"success"`
	// 原始文件信息（平铺）
	FullName    string `json:"fullName"`    // 原始完整文件名
	Path        string `json:"path"`        // 原始路径
	Dir         string `json:"dir"`         // 原始目录
	Ext         string `json:"ext"`         // 原始扩展名
	BaseName    string `json:"baseName"`    // 原始基础名
	// 新的文件信息（平铺，只在有变化时设置）
	NewFullName string `json:"newFullName"` // 新的完整文件名（如果有变化）
	NewPath     string `json:"newPath"`     // 新的路径（如果有变化）
	NewDir      string `json:"newDir"`      // 新的目录（如果有变化）
	NewExt      string `json:"newExt"`      // 新的扩展名（如果有变化）
	NewBaseName string `json:"newBaseName"` // 新的基础名（如果有变化）
	Message     string `json:"message"`
	PreviewOnly bool   `json:"previewOnly"`
	Changed     bool   `json:"changed"` // 标识文件名是否发生变化
}

// Service 正则表达式服务
type Service struct{}

// NewService 创建正则表达式服务实例
func NewService() *Service {
	return &Service{}
}

// ValidateRegex 验证正则表达式
func (s *Service) ValidateRegex(pattern string) ValidateResult {
	if pattern == "" {
		return ValidateResult{
			Valid:   false,
			Message: "正则表达式不能为空",
		}
	}

	_, err := regexp.Compile(pattern)
	if err != nil {
		return ValidateResult{
			Valid:   false,
			Message: fmt.Sprintf("正则表达式语法错误: %s", err.Error()),
		}
	}

	return ValidateResult{
		Valid:   true,
		Message: "正则表达式语法正确",
	}
}

// RegexRename 正则表达式重命名（预览模式）
func (s *Service) RegexRename(req RenameRequest) RenameResult {
	return s.regexRename(req, true)
}

// RegexRenameExecute 执行正则表达式重命名
func (s *Service) RegexRenameExecute(req RenameRequest) RenameResult {
	return s.regexRename(req, false)
}

// regexRename 内部重命名方法
func (s *Service) regexRename(req RenameRequest, previewOnly bool) RenameResult {
	// 验证正则表达式
	validateResult := s.ValidateRegex(req.Pattern)
	if !validateResult.Valid {
		return RenameResult{
			Success:     false,
			// 原始文件信息（平铺）
			FullName:    req.FileInfo.FullName,
			Path:        req.FileInfo.Path,
			Dir:         req.FileInfo.Dir,
			Ext:         req.FileInfo.Ext,
			BaseName:    req.FileInfo.BaseName,
			Message:     validateResult.Message,
			PreviewOnly: previewOnly,
			Changed:     false,
		}
	}

	// 使用传入的文件信息
	oldPath := req.FileInfo.Path
	oldFileName := req.FileInfo.FullName
	ext := req.FileInfo.Ext
	baseName := req.FileInfo.BaseName
	dir := req.FileInfo.Dir

	// 确定要处理的文本
	text := oldFileName
	if req.IgnoreExt {
		text = baseName
	}

	// 构建正则表达式
	pattern := req.Pattern
	if req.IgnoreCase {
		pattern = "(?i)" + pattern
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return RenameResult{
			Success:     false,
			// 原始文件信息（平铺）
			FullName:    req.FileInfo.FullName,
			Path:        req.FileInfo.Path,
			Dir:         req.FileInfo.Dir,
			Ext:         req.FileInfo.Ext,
			BaseName:    req.FileInfo.BaseName,
			Message:     fmt.Sprintf("正则表达式编译失败: %v", err),
			PreviewOnly: previewOnly,
			Changed:     false,
		}
	}

	// 执行替换
	var newText string
	if req.Global {
		newText = regex.ReplaceAllString(text, req.Replacement)
	} else {
		// 只替换第一个匹配
		if match := regex.FindStringIndex(text); match != nil {
			newText = text[:match[0]] + req.Replacement + text[match[1]:]
		} else {
			newText = text
		}
	}

	// 构建新文件名
	var newFileName string
	if req.IgnoreExt {
		newFileName = newText + ext
	} else {
		newFileName = newText
	}

	// 检查是否有变化
	if newFileName == oldFileName {
		return RenameResult{
			Success:     false,
			// 原始文件信息（平铺）
			FullName:    req.FileInfo.FullName,
			Path:        req.FileInfo.Path,
			Dir:         req.FileInfo.Dir,
			Ext:         req.FileInfo.Ext,
			BaseName:    req.FileInfo.BaseName,
			Message:     "文件名没有变化",
			PreviewOnly: previewOnly,
			Changed:     false,
		}
	}

	// 创建新的文件信息变量
	newPath := filepath.Join(dir, newFileName)
	newExt := filepath.Ext(newFileName)
	newBaseName := newFileName
	if newExt != "" {
		newBaseName = newFileName[:len(newFileName)-len(newExt)]
	}

	// 构建返回结果，只在有变化时设置old*字段
	result := RenameResult{
		Success:     true,
		// 原始文件信息（平铺）
		FullName:    req.FileInfo.FullName,
		Path:        req.FileInfo.Path,
		Dir:         req.FileInfo.Dir,
		Ext:         req.FileInfo.Ext,
		BaseName:    req.FileInfo.BaseName,
		PreviewOnly: previewOnly,
		Changed:     true,
	}

	// 设置new*字段（只在有变化时设置）
	if oldFileName != newFileName {
		result.NewFullName = newFileName
	}
	if oldPath != newPath {
		result.NewPath = newPath
	}
	if req.FileInfo.Dir != dir {
		result.NewDir = dir
	}
	if req.FileInfo.Ext != newExt {
		result.NewExt = newExt
	}
	if req.FileInfo.BaseName != newBaseName {
		result.NewBaseName = newBaseName
	}

	// 如果是预览模式，直接返回结果
	if previewOnly {
		result.Message = "预览重命名结果"
		return result
	}

	// 检查新文件是否已存在
	if _, err := os.Stat(newPath); err == nil {
		result.Success = false
		result.Message = fmt.Sprintf("目标文件已存在: %s", newFileName)
		return result
	}

	// 执行重命名
	err = os.Rename(oldPath, newPath)
	if err != nil {
		result.Success = false
		result.Message = fmt.Sprintf("重命名失败: %s", err.Error())
		return result
	}

	result.Message = "重命名成功"
	return result
}