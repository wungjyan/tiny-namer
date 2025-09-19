package main

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type FileInfo struct {
	FullName string `json:"fullName"`
	Path     string `json:"path"`
	Dir      string `json:"dir"`
	Ext      string `json:"ext"`
	BaseName string `json:"baseName"`
}

type SelectFilesResult struct {
	Success bool       `json:"success"`
	Files   []FileInfo `json:"files"`
	Message string     `json:"message"`
}

func (a *App) SelectFiles() SelectFilesResult {
	filesPath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择文件",
	})
	if err != nil {
		println("Error:", err.Error())
		return SelectFilesResult{
			Success: false,
			Files:   []FileInfo{},
			Message: err.Error(),
		}
	}
	if len(filesPath) == 0 {
		return SelectFilesResult{
			Success: false,
			Files:   []FileInfo{},
			Message: "未选择文件",
		}
	}
	var files []FileInfo
	for _, filePath := range filesPath {
		dir := filepath.Dir(filePath)
		fileName := filepath.Base(filePath)
		ext := filepath.Ext(filePath)
		baseName := fileName
		if ext != "" {
			baseName = fileName[:len(fileName)-len(ext)]
		}
		fileInfo := FileInfo{
			FullName: fileName,
			BaseName: baseName,
			Path:     filePath,
			Dir:      dir,
			Ext:      ext,
		}
		files = append(files, fileInfo)
	}
	return SelectFilesResult{
		Success: true,
		Files:   files,
		Message: fmt.Sprintf("成功选择了 %d 个文件", len(files)),
	}
}

// RegexValidationResult 正则表达式校验结果
type RegexValidationResult struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

// ValidateRegex 校验正则表达式的有效性
func (a *App) ValidateRegex(pattern string) RegexValidationResult {
	if pattern == "" {
		return RegexValidationResult{
			Valid:   false,
			Message: "正则表达式不能为空",
		}
	}

	// 尝试编译正则表达式
	_, err := regexp.Compile(pattern)
	if err != nil {
		return RegexValidationResult{
			Valid:   false,
			Message: fmt.Sprintf("语法错误: %s", err.Error()),
		}
	}

	return RegexValidationResult{
		Valid:   true,
		Message: "正则表达式有效",
	}
}
