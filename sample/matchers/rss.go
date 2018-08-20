package matchers

import (
	"base01/sample/search"
	"encoding/xml"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"log"
	"net/http"
	"regexp"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"string"`
		Link    string   `xml:"link"`
	}

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

type rssMatcher struct {
}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

//发送Http get请求获取rss数据源并解码
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("no rss feed URI provided")
	}

	resp, err := http.Get(feed.URI)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		//返回一个自定义错误
		return nil, fmt.Errorf("Http Response Error %d\n", resp.StatusCode)
	}

	var document rssDocument

	err = xml.NewDecoder(resp.Body).Decode(&document)

	return &document, err
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed type[%s] Site[%s] For Uri[%s]\n",
		feed.Type, feed.Name, feed.URI)

	document, err := m.retrieve(feed)

	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		//检查标题部分是否包含搜索项
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)

		if err != nil {
			return nil, err
		}

		//找到匹配项,将其作为结果保存
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		//检查描述部门是否包含搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)

		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}
