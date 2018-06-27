package fetcher

import "fmt"
import "time"
import "io/ioutil"
import "net/http"
import "log"
import "context"
import "strings"
import "encoding/json"
import "github.com/PuerkitoBio/goquery"

type Links struct {
	Url   string
	Links  map[string]string
}
type Link struct {
	Url   string
	Fetch_date string
	Is_fetched string
}

type NewLink struct {
	Url   string
	Find_date string
}

type Fetcher struct {
	EsClient *elastic.Client
}

func (f *Fetcher)FetchUrl(url string) (html string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("fetcher %s failed", url)
		return ""
	}
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body);
		resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		html = string(robots);
	} else {
		html = ""
	}
	return		
}

func (f *Fetcher)ParseUrl(url string) (content string, links string, err error) {
	currentTime := time.Now().Local()
	timeStr := currentTime.Format("2006-01-02 15:04:05")
	content_map := map[string]string{"url": url}
	links_map := &Links{
		Url: url,
		Links: map[string]string{}}
	
	p, err := goquery.NewDocument(url)
	if err != nil {
		contentbyte, _ := json.Marshal(content_map)
		linksbyte, _ := json.Marshal(links_map)
		content = string(contentbyte)
		links = string(linksbyte)
		return
	} 
	pTitle := p.Find("title").Text()
	content_map["title"] = pTitle
	doc := p.Find("p").Text()
	content_map["content"] = doc
	content_map["fetch_date"] = timeStr

	contentbyte, _ := json.Marshal(content_map)
	linksbyte, _ := json.Marshal(links_map)
	content = string(contentbyte)
	links = string(linksbyte)
	return
}

