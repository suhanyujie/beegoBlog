package articleModel

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/orm"
	"log"
	"errors"
	"reflect"
)

//查询数据时的参数封装
type QueryParam struct {
	WhereParam []ConditionType
	OrderParam string
	GroupParam string
	LimitParam string
}

//查询的where的条件参数
type ConditionType struct {
	Column   string
	Operater string
	Value    string
}

//数据包的初始化
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", "root:root@mysql@tcp(45.118.254.107:3306)/laravel?charset=utf8", 30)
	if err != nil {
		log.Fatal(err)
	}
	orm.RegisterModel(new(BlogArticles), new(BlogContent))
	orm.RegisterModel(new(TestUser), new(TestProfile))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(err)
	}
}

//博客内容的结构体
type BlogContent struct {
	Id int `orm:"column(id);auto;pk;" description:"文章内容主键id"`
	//ArticleId  int           `orm:"size(11)"`
	Content   string        `orm:"size(6000)"`
	CreatedAt string        `orm:"size(20)"`
	UpdatedAt string        `orm:"size(20)"`
	Article   *BlogArticles `json:"id" orm:"rel(one)"`
}

// Model Struct //博客文章的结构体
type BlogArticles struct {
	Id             int          `orm:"column(id);auto;pk;" description:"文章主键id"`
	ClassId        int          `orm:"column(class_id);size(11);default(0);comment(文章频道id)" description:"文章频道id"`
	SubclassId     int          `orm:"column(subclass_id);size(11);default(0);comment(文章子频道id)"`
	Title          string       `orm:"column(title);size(50);comment(文章标题)"`
	Date           string       `orm:"column(date);size(20);comment(文章书写日期)"`
	PublishDate    string       `orm:"column(publish_date);size(20);comment(文章发布日期)"`
	PublishStatus  int          `orm:"column(publish_status);size(2);comment(文章状态)"`
	IsDel          int          `orm:"column(is_del);size(2);default(0);comment(文章是否已删除)"`
	CreatedAt      string       `orm:"column(created_at);size(20);comment(文章创建时间)"`
	UserId         int          `orm:"column(user_id);size(11);default(0);comment(文章作者id)"`
	Pv             int          `orm:"column(pv);size(11);default(0);comment(文章的pv统计)"`
	ArticleContent *BlogContent `json:"article_id" orm:"reverse(one)"`
}

type TestUser struct {
	Id      int
	Name    string
	Profile *TestProfile `orm:"rel(one)"` // OneToOne relation
}

type TestProfile struct {
	Id   int
	Age  int16
	User *TestUser `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

/**
测试 测试
 */
func GetTest(param []ConditionType) (error, BlogArticles) {
	o1 := orm.NewOrm()
	qs := o1.QueryTable("test_user")
	var article BlogArticles
	if len(param) < 1 {
		return errors.New("查询条件为空！"), article
	}
	for _, oneCondition := range param {
		qs = qs.Filter(oneCondition.Column, oneCondition.Value)
	}
	err := qs.RelatedSel().One(&article)
	if err != nil {
		log.Println(err)
	}
	log.Println(article)
	return nil, article
}

/**
查询数据
 */
func GetList(page,pageSize int,filters ...interface{}) ([]*BlogArticles,int64) {
	offset := (page-1)*pageSize
	list := make([]*BlogArticles,0)
	query := orm.NewOrm().QueryTable("blog_articles")
	if len(filters)>0 {
		l := len(filters)
		for k:=0;k<l;k+=2{
			query = query.Filter((filters[k]).(string), filters[k+1])
		}
	}
	total,_ := query.Count()
	query.OrderBy("-id").Limit(pageSize,offset).All(&list)
	return list,total
}

/**
查询一条数据中的某个字段值
 */
func GetOne(param []ConditionType, field string) (error, string) {
	o1 := orm.NewOrm()
	qs := o1.QueryTable("blog_articles")
	var article BlogArticles
	if len(param) < 1 {
		return errors.New("查询条件为空！"), ""
	}
	for _, oneCondition := range param {
		qs = qs.Filter(oneCondition.Column, oneCondition.Value)
	}
	err := qs.RelatedSel().One(&article)
	if err == orm.ErrNoRows {
		return err, ""
	}
	if err != nil {
		return err, ""
	}
	//通过反射 获取某个字段的值
	immutable := reflect.ValueOf(article)
	val := immutable.FieldByName(field).String()
	log.Println(val)
	return nil, val
}
