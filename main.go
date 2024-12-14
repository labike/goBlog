package main

import (
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	// 返回json格式数据
// 	w.Header().Set("Content-Type", "application/json")
// 	var indexData IndexData
// 	indexData.Title = "ohayo"
// 	indexData.Desc = "this is ohayo blog"
// 	jsonStr, _ := json.Marshal(indexData)

// 	w.Write(jsonStr)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	var indexData IndexData
// 	indexData.Title = "home page"
// 	indexData.Desc = "this is home page"

// 	t := template.New("index.html")
// 	path, _ := os.Getwd()
// 	t, _ = t.ParseFiles(path + "/template/index.html")
// 	t.Execute(w, indexData)
// }

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", homePage)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
