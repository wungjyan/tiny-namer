package file

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// FileInfo 文件信息结构
type FileInfo struct {
	FullName string `json:"fullName"`
	Path     string `json:"path"`
	Dir      string `json:"dir"`
	Ext      string `json:"ext"`
	BaseName string `json:"baseName"`
}

// SelectFilesResult 文件选择结果
type SelectFilesResult struct {
	Success bool       `json:"success"`
	Files   []FileInfo `json:"files"`
	Message string     `json:"message"`
}

// Service 文件服务
type Service struct {
	ctx context.Context
}

// NewService 创建文件服务实例
func NewService(ctx context.Context) *Service {
	return &Service{ctx: ctx}
}

// SelectFiles 选择文件
func (s *Service) SelectFiles() SelectFilesResult {
	filesPath, err := runtime.OpenMultipleFilesDialog(s.ctx, runtime.OpenDialogOptions{
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