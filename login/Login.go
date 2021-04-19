package login

import (
	//"fmt"
	"os"
	"github.com/gocolly/colly"	
	"go-api/format"
	"io/ioutil"
	"bytes"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Login(cookie string)(name string){
	var data []byte
	file,_ := os.Create("test.txt")
	defer file.Close()
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request){
		r.Headers.Set("Cookie",cookie)
	})
	c.OnResponse(func (r *colly.Response)  {
		data = r.Body
		//fmt.Println(string(r.Body))
	})
	// c.OnHTML("table[cellspacing]",func (e *colly.HTMLElement) {
	// 	fmt.Println(e.ChildText("td"))
	// 	name = e.Text
	// })
	c.Visit("https://www.mcu.edu.tw/student/new-query/default.asp")	
	data,_ = DecodeBig5(data)
	name = format.Format(data)
	return
}

func DecodeBig5(s []byte) ([]byte, error) {
    I := bytes.NewReader(s)
    O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
    b, err := ioutil.ReadAll(O)
    if err != nil {
        return nil, err
    }
    return b, nil
}
