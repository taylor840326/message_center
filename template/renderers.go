package template

import (
	"bytes"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

type TextRenderer struct {
	template *textTemplate.Template
}

type HTMLRenderer struct {
	template *htmlTemplate.Template
}

type MarkdownRenderer struct {
	template *textTemplate.Template
}

// NewTextRenderer 创建文本渲染器
func NewTextRenderer(content string) (*TextRenderer, error) {
	tmpl, err := textTemplate.New("text").Parse(content)
	if err != nil {
		return nil, err
	}
	return &TextRenderer{template: tmpl}, nil
}

// NewHTMLRenderer 创建HTML渲染器
func NewHTMLRenderer(content string) (*HTMLRenderer, error) {
	tmpl, err := htmlTemplate.New("html").Parse(content)
	if err != nil {
		return nil, err
	}
	return &HTMLRenderer{template: tmpl}, nil
}

// NewMarkdownRenderer 创建Markdown渲染器
func NewMarkdownRenderer(content string) (*MarkdownRenderer, error) {
	tmpl, err := textTemplate.New("markdown").Parse(content)
	if err != nil {
		return nil, err
	}
	return &MarkdownRenderer{template: tmpl}, nil
}

// Render 实现TextRenderer的渲染方法
func (r *TextRenderer) Render(data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := r.template.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Render 实现HTMLRenderer的渲染方法
func (r *HTMLRenderer) Render(data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := r.template.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Render 实现MarkdownRenderer的渲染方法
func (r *MarkdownRenderer) Render(data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := r.template.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
