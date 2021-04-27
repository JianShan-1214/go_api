package format

import (
	. "fmt"
	"os"
	// "strings"

	"github.com/PuerkitoBio/goquery"
)

func Format(data []byte)(name string){
	file,_ := os.Create("test.txt")
	defer file.Close()
	dom,err := goquery.NewDocument(string(data))
	Println("================================")
	if err != nil{
		Print("Error :")
		Println(err)
	}else{
	dom.Find("table").Siblings().Each(func(i int, s *goquery.Selection){
		file.WriteString(s.Text())
		Println(s.Text())
	})}
	return 
}

