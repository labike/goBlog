package models

import (
	"goBlog/config"
	"html/template"
	"time"
)

type Post struct {
	// 文章id
	Pid int `json:"pid"`
	// 文章标题
	Title string `json:"title"`
	// 自定义页面path
	Slug string `json:"slug"`
	// 文章内容
	Content string `json:"content"`
	// 文章markdonw
	Markdown string `json:"markdown"`
	// 文章分类id
	CategoryId int `json:"categoryId"`
	// 用户id
	UserId int `json:"userId"`
	// 浏览次数
	ViewCount int `json:"viewCount"`
	// 文章类型 0 普通 1 自定义文章
	Type int `json:"type"`
	// 创建时间
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type PostReq struct {
	// 文章id
	Pid int `json:"pid"`
	// 文章标题
	Title string `json:"title"`
	// 自定义页面path
	Slug string `json:"slug"`
	// 文章内容
	Content string `json:"content"`
	// 文章 markdonw
	Markdown string `json:"markdown"`
	// 文章分类id
	CategoryId int `json:"categoryId"`
	// 用户id
	UserId int `json:"userId"`
	// 文章类型 0 普通 1 自定义文章
	Type int `json:"type"`
	// 浏览次数
	ViewCount int `json:"viewCount"`
	// 创建时间
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type PostMore struct {
	// 文章id
	Pid int `json:"pid"`
	// 文章标题
	Title string `json:"title"`
	// 自定义页面path
	Slug string `json:"slug"`
	// 文章内容
	Content template.HTML `json:"content"`
	// 文章分类id
	CategoryId int `json:"categoryId"`
	// 分类名称
	CategoryName string `json:"categoryName"`
	// 用户id
	UserId int `json:"userId"`
	// 文章类型 0 普通 1 自定义文章
	Type int `json:"type"`
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"`
	Title string `orm:"title" json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
