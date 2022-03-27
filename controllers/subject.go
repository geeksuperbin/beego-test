package controllers

import (
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"hello/models"
)

type SubjectController struct {
	beego.Controller
}

// 获取题目信息函数
func (c *SubjectController) Get() {
	var subject models.Subject // 设置变量
	// 数据校验模块
	err := func() error {
		// 从浏览器接收参数
		id, err := c.GetInt("id")
		// 记录日志
		logs.Info(id)
		if err != nil {
			id = 1
		}

		subject, err = models.GetSubject(id)

		if err != nil {
			return errors.New("subject not exist")
		}
		// 返回校验逻辑没有问题
		return nil
	}()

	// 判断数据校验模块是否整体有错误产生
	if err != nil {
		// 通过上下文模块返回错误信息
		c.Ctx.WriteString("wrong params")
	}

	// 将选项转成 map 格式，并将map 格式传递给模板文件
	var option map[string]string
	// 进行 json  decode 操作
	if err = json.Unmarshal([]byte(subject.Option), &option); err != nil {
		c.Ctx.WriteString("wrong params, json decode")
	}
	c.Data["ID"] = subject.Id
	c.Data["Option"] = option
	c.Data["Img"] = "/static" + subject.Img
	c.TplName = "guess.tpl"

}

// 接收答案提交
func (c *SubjectController) Post() {
	var subject models.Subject
	err := func() error {
		id, err := c.GetInt("id")
		logs.Info(id)

		if err != nil {
			id = 1
		}
		subject, err = models.GetSubject(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()

	if err != nil {
		c.Ctx.WriteString("wrong params")
	}

	answer := c.GetString("key") // 提交的答案
	right := models.Answer(subject.Id, answer) // 获取正确的答案匹配结果

	c.Data["Right"] = right
	c.Data["Next"] = subject.Id + 1
	c.Data["ID"] = subject.Id
	c.TplName = "guess.tpl"


}