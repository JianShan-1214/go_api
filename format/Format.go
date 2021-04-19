package format

import (
	. "fmt"
	"os"
	// "strings"

	"github.com/PuerkitoBio/goquery"
)

func Format(str []byte)(name string){
	file,_ := os.Create("test.txt")
	defer file.Close()
	Println(string(str))
	// Println(str)
	// str = strings.Replace(str,"\n","",-1)
	//Println(str+"\n\n")
	dom,err := goquery.NewDocument(string(str))
	//Println(string(str))
	if err != nil{
		Print("Error :")
		Println(err)
	}else{
	dom.Find("td").Siblings().Each(func(i int, s *goquery.Selection){
		file.WriteString(s.Text())
		Println(s.Text())
	})}
	return 
}

