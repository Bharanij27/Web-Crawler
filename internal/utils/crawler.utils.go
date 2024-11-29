package utils

import (
	"strings"

	"golang.org/x/net/html"
)

func ExtractMetadata(doc *html.Node) (string, string) {
	var title, content string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
		}
		if n.Type == html.TextNode {
			pageDetail := strings.TrimSpace(n.Data)
			if strings.Contains(n.Data, "function(") ||
				strings.Contains(n.Data, "<iframe") ||
				strings.Contains(n.Data, "<script") ||
				strings.Contains(n.Data, "window.") {
				return
			}
			if len(pageDetail) > 0 {
				content += pageDetail + " "
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if len(content) > 100 {
		content = content[:100]
	}
	return title, content
}

func ExtractLinks(doc *html.Node, baseURL string) []string {
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link := strings.TrimSpace(attr.Val)
					if !strings.HasPrefix(link, "http") {
						link = baseURL + link
					}
					links = append(links, link)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links
}
