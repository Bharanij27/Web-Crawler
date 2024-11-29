package internal

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/BharaniJ27/Web-Crawler/internal/model"
	"github.com/BharaniJ27/Web-Crawler/internal/utils"
	"golang.org/x/net/html"
)

type Crawler struct {
	visitedURLs   map[string]bool
	urlVisitCount int
	mu            sync.Mutex
}

func NewCrawler() *Crawler {
	return &Crawler{
		visitedURLs:   make(map[string]bool),
		urlVisitCount: 1,
	}
}

func (c *Crawler) Crawl(url string, maxDepth int) {
	urlList := make(chan string)
	results := make(chan model.SiteInfo)
	var wg sync.WaitGroup

	for i := 0; i < maxDepth; i++ {
		go c.siteTraverser(urlList, results, &wg, maxDepth)
	}

	wg.Add(1)
	go func() {
		urlList <- url
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for page := range results {
		fmt.Printf("URL: %s \nTitle: %s \nContent: %s\n \n", page.URL, page.Title, page.Content)
	}
}

func Crawl(startURL string, maxDepth int) {
	crawler := NewCrawler()
	crawler.Crawl(startURL, maxDepth)
}

func (c *Crawler) siteTraverser(urls chan string, results chan<- model.SiteInfo, wg *sync.WaitGroup, maxDepth int) {
	for url := range urls {
		c.mu.Lock()
		if c.visitedURLs[url] {
			c.mu.Unlock()
			wg.Done()
			continue
		}
		c.visitedURLs[url] = true
		c.mu.Unlock()

		page, links, err := c.fetchPage(url)
		if err != nil {
			fmt.Printf("Error fetching URL %s: %v\n", url, err)
			wg.Done()
			continue
		}

		results <- page

		// Process links if within depth
		if c.urlVisitCount < maxDepth {
			for _, link := range links {
				if c.urlVisitCount >= maxDepth {
					break
				} else {
					c.urlVisitCount++
				}

				wg.Add(1)
				go func(link string) {
					urls <- link
				}(link)
			}
		}
		wg.Done()
	}
}

func (c *Crawler) fetchPage(url string) (model.SiteInfo, []string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.SiteInfo{}, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.SiteInfo{}, nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return model.SiteInfo{}, nil, err
	}

	title, content := utils.ExtractMetadata(doc)
	links := utils.ExtractLinks(doc, url)

	return model.SiteInfo{
		URL:     url,
		Title:   title,
		Content: content,
	}, links, nil
}
