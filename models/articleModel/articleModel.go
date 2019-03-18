package articleModel

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"log"
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

const articleTable = "blog_articles_copy1"

//数据包的初始化
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	host := beego.AppConfig.String("db::host")
	user := beego.AppConfig.String("db::user")
	passwd := beego.AppConfig.String("db::passwd")
	port := beego.AppConfig.String("db::port")
	database := beego.AppConfig.String("db::database")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, passwd, host, port, database)
	// root@mysql
	err := orm.RegisterDataBase("default", "mysql", dataSource, 30)
	if err != nil {
		log.Fatal(err)
	}
	orm.RegisterModel(new(BlogArticlesCopy1), new(BlogContent))
	orm.RegisterModel(new(TestUser), new(TestProfile))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(err)
	}
}

//博客内容的结构体
type BlogContent struct {
	Id        int    `orm:"column(id);auto;pk;" description:"文章内容主键id"`
	ArticleId int    `orm:"column(article_id);size(11)"`
	Content   string `orm:"size(6000)"`
	CreatedAt string `orm:"size(20)"`
	UpdatedAt string `orm:"size(20)"`
}

// Model Struct //博客文章的结构体
type BlogArticles struct {
	Id            int    `orm:"column(id);auto;pk;" description:"文章主键id"`
	ClassId       int    `orm:"column(class_id);size(11);default(0);comment(文章频道id)" description:"文章频道id"`
	SubclassId    int    `orm:"column(subclass_id);size(11);default(0);comment(文章子频道id)"`
	Title         string `orm:"column(title);size(50);comment(文章标题)"`
	Date          int64  `orm:"column(date);size(20);comment(文章书写日期)"`
	PublishDate   int64  `orm:"column(publish_date);size(20);comment(文章发布日期)"`
	PublishStatus int    `orm:"column(publish_status);size(2);comment(文章状态)"`
	IsDel         int    `orm:"column(is_del);size(2);default(0);comment(文章是否已删除)"`
	CreatedAt     string `orm:"column(created_at);size(20);comment(文章创建时间)"`
	UserId        int    `orm:"column(user_id);size(11);default(0);comment(文章作者id)"`
	Pv            int    `orm:"column(pv);size(11);default(0);comment(文章的pv统计)"`
	Content       string
}

//取别名
type BlogArticlesCopy1 = BlogArticles

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

//新增数据
func Add(data *BlogArticles) (int64, error) {
	o := orm.NewOrm()
	insertId, err := o.Insert(data)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return insertId, nil
}

/**
查询数据
*/
func GetList(page, pageSize int, filters ...interface{}) ([]*BlogArticles, int64) {
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	list := make([]*BlogArticles, 0)
	o := orm.NewOrm()
	query := o.QueryTable(articleTable)

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter((filters[k]).(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	//根据列表获取文章内容
	if len(list) < 1 {
		return list, total
	}
	//log.Println(list[0])
	var ids []int
	for _, article := range list {
		ids = append(ids, article.Id)
	}
	contentList := make([]*BlogContent, 0)
	var contentMapData = make(map[int](*BlogContent), 0)
	query = o.QueryTable("blog_content")
	query = query.Filter("article_id__in", ids)
	query.All(&contentList)
	for _, oneContent := range contentList {
		contentMapData[oneContent.ArticleId] = oneContent
	}
	var tmpContent string
	for i, article := range list {
		if contentMapData[article.Id] == nil {
			continue
		}
		tmpContent = contentMapData[article.Id].Content
		if tmpContent != "" {
			list[i].Content = beego.Substr(contentMapData[article.Id].Content, 0, 120)
		} else {
			list[i].Content = ""
		}
	}

	return list, total
}

func GetRaw(param []ConditionType) (error, BlogArticles) {
	o := orm.NewOrm()
	qs := o.QueryTable(articleTable)
	//article := new(BlogArticles)
	var article BlogArticles
	if len(param) < 1 {
		return errors.New("查询条件为空！"), article
	}
	for _, oneCondition := range param {
		qs = qs.Filter(oneCondition.Column, oneCondition.Value)
	}
	err := qs.One(&article)
	if err == orm.ErrNoRows {
		return err, article
	}
	if err != nil {
		return err, article
	}
	contentObj := &BlogContent{}
	qs = o.QueryTable("blog_content")
	qs = qs.Filter("article_id", article.Id)
	qs.One(contentObj)
	article.Content = contentObj.Content

	return nil, article
}

/**
查询一条数据中的某个字段值
*/
func GetOne(param []ConditionType, field string) (error, string) {
	o1 := orm.NewOrm()
	qs := o1.QueryTable(articleTable)
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

func InsertArticle(articleObj *BlogArticlesCopy1, contentObj *BlogContent, other map[string]string) (int64, error) {
	o1 := orm.NewOrm()
	id, err := o1.Insert(articleObj)
	if err != nil {
		return 0, err
	}
	//查看文章是否已存在
	contentObj.ArticleId = int(id)
	_, err = o1.Insert(contentObj)
	if err != nil {
		return 0, err
	}
	return id, nil
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
