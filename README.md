## 基于beego的个人博客

## 环境
* go1.11
* beego 1.10.*
* 模板 layui轻博客

## 关于beego
### orm关联模型的一些理解
* 在关联模型中，t1表中有个成员属性Profile，它是关联到表t2的一个字段
* 如果在t1表中有一个字段是profile_id，那么，在声明t1的结构模型时，对应的成员属性Profile，设置为正向关联到t2表。也就是说，这种情况是正向关联`rel(one)`
* 如果这个关联字段在t1中不存在，而是存储在t2表，那么这种关系就是反向关联。

## 一些知识点
### 1.时间和时间戳
* 时间字符串的写法：`time.Now().Format("2006-01-02 15:04:05")`
* 时间戳的写法：`time.Now().Unix()`

### 2.分页的示例如下：
* https://beego.me/docs/mvc/view/page.md

## 踩坑全书
### 1.beego orm的in查询
* 在beego orm文档中 https://beego.me/docs/mvc/model/query.md
* 它的in查询如下：

```go
qs.Filter("profile__age__in", 17, 18, 19, 20)
// WHERE profile.age IN (17, 18, 19, 20)


ids:=[]int{17,18,19,20}
qs.Filter("profile__age__in", ids)
// WHERE profile.age IN (17, 18, 19, 20)

// 同上效果
```

* 按照这个文档尝试了多次，发现都是提示
> unknown field/column name ....

* 后来，尝试追源码，发现还是自己写法问题：in 查询，只需要 {字段名__in} 就可以了，不需要 {表名__字段名__in}

### 2.部署beegoBlog时模板找不到
* 因为在beego中，执行可执行文件后，找模板的规则，是在当前路径下找，因此，需要将项目根目录和模板目录保持一致
* 也就是将可执行文件放在项目根目录。（之前一致以为，只需要将可执行文件放到服务器上部署就可以了。。）



