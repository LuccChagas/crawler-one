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

func NewCollector() *colly.Collector {

	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	return c
}

//TODO: O rastreador deve ser uma ferramenta de linha de comando que aceita um URL inicial e um diretório de destino.
func (cc *CrawlerConstructor) NewCrawler(c *colly.Collector, args *Args) {

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Pega o link dentro do href
		link := e.Attr("href")

		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// visita o link dentro da pagina encontrada
		e.Request.Visit(link)

		//TODO: salvar conteudo do body do link no diretorio recebido
		SaveContent(e.Request, e.Response, args)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ->", r.URL.String())
	})

	c.Visit(args.Url)
	c.Wait()
}

func SaveContent(req *colly.Request, resp *colly.Response, args *Args) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileName := util.SanitizeUrlFileName(req.URL.String())
	file := util.CheckAndFileCreation(fileName, resp.Body)

	if err := os.Rename(
		fmt.Sprintf("%s/%s", dir, file.Name()),
		fmt.Sprintf("%s/%s", args.Source, file.Name()),
	); err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	return nil
}
