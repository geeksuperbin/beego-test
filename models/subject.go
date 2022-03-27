package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"strings"
)

type Subject struct {
	Id int
	Option string
	AnswerKey string
	Status int8
	Img string
}


func init() {
	orm.RegisterModel(new(Subject))
}

func GetSubject(id int) (s Subject, err error) {
	o := orm.NewOrm()
	// 选择数据库
	//o.Using("pycontrol")
	o.Using("default")
	s = Subject{Id:id}
	err = o.Read(&s)

	if err != nil {
		return s, err
	}
	return
}

func Answer(sid int, answerKey string) bool {
	subject, err := GetSubject(sid)
	if err != nil {
		return false
	}
	return strings.Compare(strings.ToUpper(answerKey), subject.AnswerKey) == 0 // 表示匹配到
}