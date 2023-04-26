package main

import (
	"flag"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/googege/gotools/id"
)

var (
	url string
)

func init() {
	flag.StringVar(&url, "url", "https://www.xvideos.com/video75472711/_av_", "key word")

	flag.Parse()
}

func main() {
	downloadVedio(dealwithJS(url))
}

// func do(root string) {
// 	url := make([]string,0)
// 	// getDownloadUrl(root, url)
// 	//
// 	wg := new(sync.WaitGroup)
// 	wg.Add(32)
// 	readUrl := make(chan string)
// 	for i := 0; i < 32; i++ {
// 		go func() {
// 			defer wg.Done()
// 			for v := range url {
// 				dealwithJS(v, readUrl)
// 			}
// 		}()
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(readUrl)
// 	}()
// 	//
// 	wg1 := new(sync.WaitGroup)
// 	wg1.Add(8)
// 	for i := 0; i < 8; i++ {
// 		go func() {
// 			defer wg1.Done()
// 			for v := range readUrl {
// 				downloadVedio(v)
// 			}
// 		}()
// 	}
// 	wg1.Wait()
// }

// get download url.
// func getDownloadUrl(root string, downloadUrl chan string) {
// 	go func() {
// 		defer close(downloadUrl)
// 		co := colly.NewCollector()
// 		fmt.Println("url is  :",root)
// 		re, _ := regexp.Compile("\\/video\\d+.*")
// 		ma := make(map[string]int)
// 		result := make([]string, 0)
// 		co.OnHTML("a[href]", func(element *colly.HTMLElement) {
// 			a := re.Find([]byte(element.Attr("href")))
// 			if len(a) != 0 {
// 				if _, ok := ma[string(a)]; !ok {
// 					d := fmt.Sprintf("https://www.xvideos.com%s", string(a))
// 					result = append(result, d)
// 				}
// 				ma[string(a)]++
// 			}
// 		})
// 		if err := co.Visit(root);err != nil {
// 			glog.Error(err," "+root)
// 		}
// 		for k, v := range result {
// 			if k <= number {
// 				downloadUrl <- v
// 			}
// 		}
// 	}()
// }

func dealwithJS(root string) string {
	co := colly.NewCollector()
	var thing []string
	co.OnHTML("body", func(element *colly.HTMLElement) {
		thing = element.ChildTexts("script")

	})
	co.Visit(root)
	re, err := regexp.Compile("html5player.setVideoUrlHigh\\(\\'.*\\'\\)")
	if err != nil {
		fmt.Println(err)
	}
	tt := ""
	for _, v := range thing {
		a := string(re.Find([]byte(v)))
		if a != "" {
			tt = a
		}
	}
	re1, err := regexp.Compile("https\\:\\/\\/.*\\'")
	ddd := re1.Find([]byte(tt))
	return string(ddd)
}
func downloadVedio(url string) {
	fmt.Println(url[:60]+"...", "is downloading...")
	u, _ := id.NewUUID(id.VERSION_1, nil)
	c := fmt.Sprintf("-O %s.mp4", u.String())
	cmd := exec.Command("wget", url, c)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("download info is :", string(out))
	cmd.Run()
}
