package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCrawlerConstructor(t *testing.T) {
	arg := &Args{
		Url:    "https://wiki.hackerspaces.org/",
		Source: "/Users/luccas/go/projects/crawler-one/service/temp/",
	}

	cc := &CrawlerConstructor{
		args: arg,
	}

	require.NotEmpty(t, cc)
}

func TestNewCollector(t *testing.T) {
	c, err := NewCollector()
	require.NotEmpty(t, c)
	require.NoError(t, err)
}

func TestNewCrawler(t *testing.T) {
	c, err := NewCollector()
	require.NoError(t, err)

	arg := &Args{
		Url:    "https://wiki.hackerspaces.org/",
		Source: "/Users/luccas/go/projects/crawler-one/service/temp/",
	}

	a := NewCrawlerConstructor(arg)
	a.NewCrawler(c, arg)
}
