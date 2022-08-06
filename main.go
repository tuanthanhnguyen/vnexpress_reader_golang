package main

import (
	"fmt"
	"log"
	"strconv"
	//"strings"
)

//functions for later use
/*func filename(url string) string {
	rst := strings.TrimPrefix(url, "https://vnexpress.net/")
	rst = strings.TrimSuffix(rst, ".html")
	return rst
}*/
/*
func choice_splitter(sel_data string) []int {
	sel_data = strings.ReplaceAll(sel_data, " ", "")
	choices := strings.Split(sel_data, ",")
	var payload []int
	for _, item := range choices {
		intval, _ := strconv.Atoi(item)
		payload = append(payload, intval)
	}
	return payload

}*/
func main() {
	var choice int
	var article_data map[string]string
	var text string
	var paths []string
	url_list := make(map[int]string)
	counter := 1
	fmt.Println("Chào mừng đến phần mềm đọc tin trên Vnexpress.net")
	fmt.Println("############################################################")
	fmt.Println("1 : Chỉ đọc Top Story ( 3 bài )")
	fmt.Println("2 : Đọc tất cả ( gồm Top Story và khoảng 20 Editor's Picks )")
	fmt.Println("3 : Bỏ từ bài thứ n")
	fmt.Print("Lựa chọn : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		article_data = get_articles(3)
	case 2:
		article_data = get_articles(0)
	case 3:
		full_list := get_articles(0)
		fmt.Println("Danh sách bài đọc:")
		for link, item := range full_list {
			fmt.Println(counter, ":", full_list[item])
			url_list[counter] = link
			counter++
		}
	}
	for link := range article_data {
		fmt.Println("Đang đọc bài : ", article_data[link])
		text = get_text(link)
		segments := splitter(text)
		for index, chunk := range segments {
			pth := gg_recognise(chunk, strconv.Itoa(index))
			paths = append(paths, pth)
		}
		for _, path := range paths {
			err := play(path)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
