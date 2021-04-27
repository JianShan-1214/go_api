package login

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Login(username string, password string)(cookie string) {
	var Login_url = "https://www.mcu.edu.tw/student/new-query/Chk_Pass_New_v1.asp"
	var formData = map[string]string{
		"t_tea_no": username,
		"t_tea_pass": password,
	}
	c := colly.NewCollector()
	err := c.Post(Login_url, formData)
	if err != nil {
		fmt.Println(err)
	}
	c.OnResponse(func(r *colly.Response){
		s := r.Request.Headers.Get("Cookie")
		// fmt.Println(s)
		cookie = s;
	})
	c.Visit(Login_url)
	return 
}