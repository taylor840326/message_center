package template

import (
	"fmt"
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
