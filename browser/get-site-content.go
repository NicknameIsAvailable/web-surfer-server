package browser

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetHTMLContent(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ошибка: %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func RemoveStylesAndScripts(htmlContent string) string {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	var cleanedBuffer bytes.Buffer
	traverseAndRemoveStylesAndScripts(doc, &cleanedBuffer)

	words := strings.Fields(cleanedBuffer.String())
	cleanedString := strings.Join(words, " ")

	return cleanedString
}

func traverseAndRemoveStylesAndScripts(node *html.Node, buffer *bytes.Buffer) {
	if node.Type == html.ElementNode {
		for i, attr := range node.Attr {
			if attr.Key == "style" {
				node.Attr = append(node.Attr[:i], node.Attr[i+1:]...)
				break
			}
		}

		if node.Data == "script" {
			node.FirstChild = nil
		}
	}

	if node.Type == html.TextNode {
		buffer.WriteString(node.Data)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverseAndRemoveStylesAndScripts(c, buffer)
	}
}
