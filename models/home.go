package models

import "goBlog/config"

type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []Post
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}