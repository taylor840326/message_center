package template

import (
	"fmt"
	"message_center/message"
	"testing"
)

// 示例数据结构
type Data struct {
	Title   string
	Content string
}

func TestTemplateRenderers(t *testing.T) {
	// 模板内容
	const templateContent = `
    标题: {{.Title}}
    内容: {{.Content}}
    `
	/*  企业微信模板内容
	标题: PostgreSQL database has slow queries and the value is 87.16. (instance: platformdb-postgresql-0, database: juicefsdb) \n
	> 命名空间: <font color=\"comment\">prod-pg-arch</font>\n
	> 服务器节点: <font color=\"comment\"></font>\n
	> pod名称: <font color=\"comment\">platformdb-postgresql-0</font>
	\n> 状态: <font color=\"info\">resolved</font>\n
	\n"
	*/

	// 准备数据
	data := Data{
		Title:   "测试标题",
		Content: "这是测试内容",
	}

	// 获取HTML渲染器
	renderer, err := GetRenderer(templateContent, FormatHTML)
	if err != nil {
		panic(err)
	}

	// 渲染内容
	result, err := renderer.Render(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestMessageTemplate(t *testing.T) {
	// 企业微信消息模板
	const wecomTemplate = `标题: {{.Description.Default}} 
> 命名空间: <font color="comment">{{.Labels.namespace.Default}}</font>
> 服务器节点: <font color="comment">{{.Labels.node.Default}}</font>
> pod名称: <font color="comment">{{.Labels.pod.Default}}</font>
> 状态: <font color="info">{{.Status.Default}}</font>
> 严重程度: <font color="info">{{.Severity}}</font>
`

	// 创建测试数据
	msg := message.Message{
		Description: message.I10nField{
			Default: "PostgreSQL database has slow queries and the value is 87.16. (instance: platformdb-postgresql-0, database: juicefsdb)",
		},
		Labels: map[string]message.I10nField{
			"namespace": {Default: "prod-pg-arch"},
			"node":      {Default: "g001.local"},
			"pod":       {Default: "platformdb-postgresql-0"},
		},
		Status: message.I10nField{
			Default: "resolved",
		},
		Severity: message.I10nField{
			Default: "critical",
		},
	}

	// 获取文本渲染器（因为企业微信消息实际上是文本格式）
	renderer, err := GetRenderer(wecomTemplate, FormatText)
	if err != nil {
		t.Fatal(err)
	}

	// 渲染消息
	result, err := renderer.Render(msg)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}
