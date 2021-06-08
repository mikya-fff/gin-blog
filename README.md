# gin-blog

## 环境准备
- 数据库可视化软件

https://www.xp.cn/linux.html

安装完成后会提示登录地址，用户名密码等信息

- vscode插件

作者.插件名-插件版本
```text
octref.vetur-0.31.3
visualstudioexptteam.vscodeintellicode-1.2.14
zhuangtongfa.material-theme-3.8.5
pkief.material-icon-theme-4.4.0
esbenp.prettier-vscode-5.8.0
formulahendry.auto-close-tag-0.5.10
formulahendry.auto-rename-tag-0.1.6
vincaslt.highlight-matching-tag-0.10.0
hollowtree.vue-snippets-1.0.4
whtouche.vscode-js-console-utils-0.7.0
vuetifyjs.vuetify-vscode-0.2.0
dxabikos.javascriptsnippets-1.8.0
ritwickdey.liveserver-5.6.1
```
- go env参数
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct // 使用七牛云的
```

## go ini -- 解析配置文件

https://ini.unknwon.io/

## gin 

https://gin-gonic.com/zh-cn/docs/

## gorm

https://gorm.io/zh_CN/docs/

### 错误

```txt
[error] invalid field found for struct gin-blog/model.Article's field Category, need to define a valid foreign key for relations or it need to implement the Valuer/Scanner interface
```

```go
type Article struct {
	gorm.Model
	Category Category
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int `gorm:"type:int;not null" json:"cid"`
	Desc string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img string `gorm:"type:varchar(100)" json:"img"`
}
```

缺少外键,添加下面的外键就可以了。

```go
type Article struct {
	gorm.Model
    //添加外键依赖
	Category Category `gorm:"foreignkey:Name"`
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int `gorm:"type:int;not null" json:"cid"`
	Desc string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img string `gorm:"type:varchar(100)" json:"img"`
}
```

