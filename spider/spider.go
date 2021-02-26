package spider

import (
	"log"
	"strings"
	"time"

	"fanfan.me/DoubanSpider/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

func Spider(db *gorm.DB) {
	startUrl := "https://movie.douban.com/top250"

	// 创建Collector
	collector := colly.NewCollector(
		// 设置用户代理
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36"),
	)

	// 设置抓取频率限制
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second, // 随机延迟
	})

	// 异常处理
	collector.OnError(func(response *colly.Response, err error) {
		log.Println(err.Error())
	})

	collector.OnRequest(func(request *colly.Request) {
		log.Println("start visit: ", request.URL.String())
	})

	// 解析列表
	collector.OnHTML("ol.grid_view", func(element *colly.HTMLElement) {
		// 依次遍历所有的li节点
		element.DOM.Find("li").Each(func(i int, selection *goquery.Selection) {
			href, found := selection.Find("div.hd > a").Attr("href")
			// 如果找到了详情页，则继续下一步的处理
			if found {
				parseDetail(collector, href, db)
				log.Println(href)
			}
		})
	})

	// 查找下一页
	collector.OnHTML("div.paginator > span.next", func(element *colly.HTMLElement) {
		href, found := element.DOM.Find("a").Attr("href")
		// 如果有下一页，则继续访问
		if found {
			element.Request.Visit(element.Request.AbsoluteURL(href))
		}
	})

	// 起始入口
	collector.Visit(startUrl)
}
func parseDetail(collector *colly.Collector, url string, db *gorm.DB) {
	collector = collector.Clone()
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second, // 随机延迟
	})
	collector.OnRequest(func(request *colly.Request) {
		log.Println("start visit :", request.URL.String())
	})
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		selection := element.DOM.Find("div#content")
		id := selection.Find("div.top250>span.top250-no").Text()
		title := selection.Find("hl>span").First().Text()
		year := selection.Find("h1 > span.year").Text()
		info := selection.Find("div#info").Text()
		info = strings.ReplaceAll(info, " ", "")
		info = strings.ReplaceAll(info, "\n", "; ")
		rating := selection.Find("strong.rating_num").Text()
		movieInfo := model.Movie{
			Id:     id,
			Title:  title,
			Year:   year,
			Info:   info,
			Rating: rating,
			Url:    url,
		}
		err := db.Save(&movieInfo).Error
		if err != nil {
			log.Printf("save movieinfo error:%s at id:%s at url:%s\n", err, id, url)
		}

	})
	collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)

	})
	collector.Visit(url)
}
