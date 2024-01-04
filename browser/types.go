package browser

type SearchResponse struct {
	Items []struct {
		Kind             string `json:"kind"`
		Title            string `json:"title"`
		HtmlTitle        string `json:"htmlTitle"`
		Link             string `json:"link"`
		DisplayLink      string `json:"displayLink"`
		Snippet          string `json:"snippet"`
		HtmlSnippet      string `json:"htmlSnippet"`
		CacheId          string `json:"cacheId"`
		FormattedUrl     string `json:"formattedUrl"`
		HtmlFormattedUrl string `json:"htmlFormattedUrl"`
	} `json:"items"`
}
