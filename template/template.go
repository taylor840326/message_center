package template

type Renderer interface {
	Render(data interface{}) (string, error)
}

type Template struct {
	content string
	format  string
}

func NewTemplate(content string, format string) *Template {
	return &Template{
		content: content,
		format:  format,
	}
}
