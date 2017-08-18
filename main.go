package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	/*str =
	req := &http.Request{}
	req, _ = http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(body))
	*/

	var (
		url  string
		mode string
	)

	flag.StringVar(&url, "u", "https://www.google.co.jp/", "URL")
	flag.StringVar(&mode, "m", "tw", "出力モードを選択 (tw|mail|rm|md)")
	flag.Parse()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	title := doc.Find("head>title").Text()

	//	fmt.Printf("url=%s\n", url)
	//	fmt.Printf("mode=%s\n", mode)
	//	fmt.Printf("title=%s\n", title)

	var str string

	switch mode {
	case "tw":
		str = fmt.Sprintf("%s %s", title, url)
	case "mail":
		str = fmt.Sprintf("%s\n%s", title, url)
	case "rm":
		str = fmt.Sprintf("\"%s\":%s", title, url)
	case "md":
		str = fmt.Sprintf("[%s](%s)", title, url)
	default:
		str = fmt.Sprintf("unknown mode")
	}

	fmt.Print(str + "\n")
}
