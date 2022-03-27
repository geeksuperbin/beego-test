package test

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"hello/models"
	_ "hello/routers"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}


//// TestBeego is a sample to run an endpoint test
//func TestBeego(t *testing.T) {
//	r, _ := http.NewRequest("GET", "/", nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)
//
//	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
//
//	Convey("Subject: Test Station Endpoint\n", t, func() {
//	        Convey("Status Code Should Be 200", func() {
//	                So(w.Code, ShouldEqual, 200)
//	        })
//	        Convey("The Result Should Not Be Empty", func() {
//	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
//	        })
//	})
//}


func Test1(t *testing.T) {
	type Subject struct {
		Id int
		Option string
		AnswerKey string
		Status int8
		Img string
	}

	fmt.Println(new(Subject))
}

func Test_Subject_Model_GetSubject(t *testing.T) {
	//orm.RegisterDataBase("default", "mysql","root:qwerzxcv@tcp(127.0.0.1:8809)/pycontrol?charset=utf8")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql","root:qwerzxcv@tcp(127.0.0.1:8809)/pycontrol?charset=utf8")
	res,_ := models.GetSubject(1)
	fmt.Println(res)

	mjson,_ :=json.Marshal(res)
	mString :=string(mjson)
	fmt.Printf(mString)

}

func Test_Subject_Model_Answer(t *testing.T){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql","root:qwerzxcv@tcp(127.0.0.1:8809)/pycontrol?charset=utf8")

	s := models.Answer(1,"c")
	fmt.Println(s)
	s1 := models.Answer(1,"d")
	fmt.Println(s1)
}


func Test_Controller_Get(t *testing.T){
	req := httplib.Get("http://127.0.0.1:8080/subject?id=1")
	str, err := req.String()
	if err != nil {
		logs.Error(err)
	}
	logs.Info(str)

}
