package main

import (
	"context"

	"tiny-namer/backend/file"
	"tiny-namer/backend/highlight"
	"tiny-namer/backend/regex"
)

// App struct
type App struct {
	ctx           context.Context
	fileService   *file.Service
	hlService     *highlight.Service
	regexService  *regex.Service
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts up. The context here
// can be used to call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.fileService = file.NewService(ctx)
	a.hlService = highlight.NewService()
	a.regexService = regex.NewService()
}

// SelectFiles 选择文件
func (a *App) SelectFiles() file.SelectFilesResult {
	return a.fileService.SelectFiles()
}

// GetHighlightParts 获取高亮部分
func (a *App) GetHighlightParts(req highlight.HighlightRequest) []string {
	return a.hlService.GetHighlightParts(req)
}

// ValidateRegex 验证正则表达式
func (a *App) ValidateRegex(pattern string) regex.ValidateResult {
	return a.regexService.ValidateRegex(pattern)
}

// RegexRename 正则表达式重命名（预览模式）
func (a *App) RegexRename(req regex.RenameRequest) regex.RenameResult {
	return a.regexService.RegexRename(req)
}

// RegexRenameExecute 执行正则表达式重命名
func (a *App) RegexRenameExecute(req regex.RenameRequest) regex.RenameResult {
	return a.regexService.RegexRenameExecute(req)
}

// PreviewRegexRename 批量预览正则表达式重命名
func (a *App) PreviewRegexRename(req struct {
	Files       []file.FileInfo `json:"files"`
	Pattern     string          `json:"pattern"`
	Replacement string          `json:"replacement"`
	IgnoreExt   bool            `json:"ignoreExt"`
	IgnoreCase  bool            `json:"ignoreCase"`
	Global      bool            `json:"global"`
}) struct {
	Success bool                  `json:"success"`
	Items   []regex.RenameResult  `json:"items"`
	Message string                `json:"message"`
} {
	if len(req.Files) == 0 {
		return struct {
			Success bool                  `json:"success"`
			Items   []regex.RenameResult  `json:"items"`
			Message string                `json:"message"`
		}{
			Success: false,
			Message: "没有选择文件",
			Items:   []regex.RenameResult{},
		}
	}

	var results []regex.RenameResult
	for _, fileInfo := range req.Files {
		renameReq := regex.RenameRequest{
			FileInfo:    fileInfo,
			Pattern:     req.Pattern,
			Replacement: req.Replacement,
			IgnoreExt:   req.IgnoreExt,
			IgnoreCase:  req.IgnoreCase,
			Global:      req.Global,
		}
		result := a.regexService.RegexRename(renameReq)
		results = append(results, result)
	}

	return struct {
		Success bool                  `json:"success"`
		Items   []regex.RenameResult  `json:"items"`
		Message string                `json:"message"`
	}{
		Success: true,
		Items:   results,
		Message: "预览完成",
	}
}
