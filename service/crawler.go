package service

import (
	"fmt"
	"log"
	"os"

	"github.com/LuccChagas/crawler-one/util"
	"github.com/gocolly/colly"
)

type CrawlerImpl interface {
	NewCrawler(c *colly.Collector, arg *Args)
}

type CrawlerConstructor struct {
	args *Args
}

func NewCrawlerConstructor(arg *Args) CrawlerImpl {
	return &CrawlerConstructor{
		args: &Args{
			Url:    arg.Url,
			Source: arg.Source,
		},
	}
}

func NewCollector() (*colly.Collector, error) {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)

	err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})
	if err != nil {
		return nil, err
	}

	return c, nil
}

//TODO: O rastreador deve ser uma ferramenta de linha de comando que aceita um URL inicial e um diretório de destino.
func (cc *CrawlerConstructor) NewCrawler(c *colly.Collector, args *Args) {
	var u string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(link)

		SaveContent(u, e.Response, args)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ->", r.URL.String())
		u = fmt.Sprint(r.URL.String())

	})

	c.Visit(args.Url)
	c.Wait()
}

func SaveContent(visiting string, resp *colly.Response, args *Args) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileName := util.SanitizeUrlFileName(visiting)
	if err != nil {
		return err
	}
	file, err := util.CheckAndFileCreation(fileName, resp.Body)
	if err != nil {
		return nil
	}

	if err := os.Rename(
		fmt.Sprintf("%s/%s", dir, file.Name()),
		fmt.Sprintf("%s/%s", args.Source, file.Name()),
	); err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	return nil
}
