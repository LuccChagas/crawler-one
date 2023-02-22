package service

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type CrawlerImpl interface {
	NewCrawler(c *colly.Collector)
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

//TODO: O rastreador deve ser uma ferramenta de linha de comando que aceita um URL inicial e um diretório de destino.
func (cc *CrawlerConstructor) NewCrawler(c *colly.Collector) {

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Pega o link dentro do href
		link := e.Attr("href")

		// visita o link dentro da pagina encontrada
		e.Request.Visit(link)

		//TODO: salvar conteudo do body do link no diretorio recebido
		SaveContent(e.Request)
		c.Wait()

	})

}

func SaveContent(req *colly.Request) error {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DIR:", dir)
	fmt.Println("OPAQUE:", req.URL.Opaque)

	file, err := os.Create(req.URL.Opaque)
	if err != nil {
		// TODO: create a nice Log message with error details (Try to put in a recovery logic)
		log.Fatal(err)
	}

	err = file.Chmod(0777)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.Rename(
		fmt.Sprintf("%s/%s", dir, file.Name()),
		fmt.Sprintf("%s/%s", "", file.Name()),
	); err != nil {
		log.Fatal(err)
	}

	return nil
}
