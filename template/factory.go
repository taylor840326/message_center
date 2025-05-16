package template

const (
	FormatText     = "text"
	FormatHTML     = "html"
	FormatMarkdown = "markdown"
)

// GetRenderer 根据格式获取对应的渲染器
func GetRenderer(content string, format string) (Renderer, error) {
	switch format {
	case FormatText:
		return NewTextRenderer(content)
	case FormatHTML:
		return NewHTMLRenderer(content)
	case FormatMarkdown:
		return NewMarkdownRenderer(content)
	default:
		return NewTextRenderer(content) // 默认使用文本渲染器
	}
}
