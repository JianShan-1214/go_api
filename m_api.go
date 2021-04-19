package main

import(
	"fmt"
	"go-api/login"
	"io/ioutil"
	//"github.com/gin-gonic/gin"
)

func main() {
	cookie,_ := ioutil.ReadFile("cookie.txt")
	fmt.Println(login.Login(string(cookie)))
}