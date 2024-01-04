package browser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func Find(query string) SearchResponse {
	apiKey := os.Getenv("SEARCH_API")
	searchId := os.Getenv("SEARCH_ID")
	baseURL := "https://www.googleapis.com/customsearch/v1"
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("cx", searchId)
	params.Add("q", query)

	var searchResponse SearchResponse

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("Ошибка при выполнении запроса: %s\n", err)
		return searchResponse
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении тела ответа: %s\n", err)
		return searchResponse
	}

	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		fmt.Printf("Ошибка при декодировании JSON: %s\n", err)
		return searchResponse
	}

	return searchResponse
}
