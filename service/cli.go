package service

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Args struct {
	Url    string `json:"url"`
	Source string `json:"source"`
}

func NewCli() error {

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	// TODO: set this content in viper via .env
	app := &cli.App{
		Name:      "My Crawler App",
		UsageText: "This app runs a crawler to extract html content the website of your chose",
		Version:   "v1.0.0",
		Authors: []*cli.Author{
			{
				Name:  "Luccas Machado",
				Email: "luccaa.chagas23@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "url",
				Usage:       "URL for start the crawler",
				DefaultText: "https://pt.wikipedia.org/wiki/Wikipédia",
				Required:    true,
			},
			&cli.StringFlag{
				Name:     "source",
				Aliases:  []string{"s"},
				Usage:    "Destination folder to put html Content of the crawler",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "start",
				Action: StartCrawler,
			},
		},
		DefaultCommand: "start",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	return nil
}

func StartCrawler(cCtx *cli.Context) error {
	arg := &Args{
		Url:    cCtx.String("url"),
		Source: cCtx.String("source"),
	}

	// verifica se a source existe, caso n existe tenta criar, caso n consiga criar
	_, err := os.Open(arg.Source)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Source directory doen't exist, so it will be created")
			errMkdir := os.Mkdir(arg.Source, os.ModePerm)
			if errMkdir != nil {
				log.Fatal("An error occurred while trying to create a Dir based in `source` flag: ", err)
			}
			log.Printf("Folder '%s' was created with success", arg.Source)
		}
	}
	c, err := NewCollector()
	if err != nil {
		log.Fatal(err)
	}
	cc := NewCrawlerConstructor(arg)
	cc.NewCrawler(c, arg)

	return nil
}
