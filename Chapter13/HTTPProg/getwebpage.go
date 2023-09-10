package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Searchrequest struct {
	base       *url.URL
	token      string
	searchterm string
}

type Main struct {
	Type  string `json:"type"`
	Index int    `json:"index"`
	All   bool   `json:"all"`
}

type Query struct {
	Original             string `json:"original"`
	ShowStrictWarning    bool   `json:"show_strict_warning"`
	IsNavigational       bool   `json:"is_navigational"`
	IsNewsBreaking       bool   `json:"is_news_breaking"`
	SpellcheckOff        bool   `json:"spellcheck_off"`
	Country              string `json:"country"`
	BadResults           bool   `json:"bad_results"`
	ShouldFallback       bool   `json:"should_fallback"`
	PostalCode           string `json:"postal_code"`
	City                 string `json:"city"`
	HeaderCountry        string `json:"header_country"`
	MoreResultsAvailable bool   `json:"more_results_available"`
	State                string `json:"state"`
}

type Mixed struct {
	Type string        `json:"type"`
	Main []Main        `json:"main"`
	Top  []interface{} `json:"top"`
	Side []interface{} `json:"side"`
}

type MetaURL struct {
	Scheme   string `json:"scheme"`
	Netloc   string `json:"netloc"`
	Hostname string `json:"hostname"`
	Favicon  string `json:"favicon"`
	Path     string `json:"path"`
}

type Results struct {
	Title          string  `json:"title"`
	URL            string  `json:"url"`
	IsSourceLocal  bool    `json:"is_source_local"`
	IsSourceBoth   bool    `json:"is_source_both"`
	Description    string  `json:"description"`
	Language       string  `json:"language"`
	FamilyFriendly bool    `json:"family_friendly"`
	Type           string  `json:"type"`
	Subtype        string  `json:"subtype"`
	MetaURL        MetaURL `json:"meta_url"`
	Age            string  `json:"age,omitempty"`
}

type Web struct {
	Type           string    `json:"type"`
	Results        []Results `json:"results"`
	FamilyFriendly bool      `json:"family_friendly"`
}

type Response struct {
	Query Query  `json:"query"`
	Mixed Mixed  `json:"mixed"`
	Type  string `json:"type"`
	Web   Web    `json:"web"`
}

func main() {
	// Ask For Token and Search Term
	searchrequest := Searchrequest{
		token:      "",
		searchterm: "",
	}

	flag.StringVar(&searchrequest.token, "token", "", "token")
	flag.StringVar(&searchrequest.searchterm, "searchterm", "", "searchterm")
	flag.Parse()

	// Build Request
	encodedsearchterm := url.PathEscape(searchrequest.searchterm)

	encodedurl := fmt.Sprintf("https://api.search.brave.com/res/v1/web/search")
	queryparam := url.Values{}
	queryparam.Set("q", encodedsearchterm)
	if len(searchrequest.searchterm) > 0 {
		encodedurl += "?" + queryparam.Encode()
	}

	fmt.Println(encodedurl)

	req, err := http.NewRequest(http.MethodGet, encodedurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = http.Header{
		"Accept":               {"application/json"},
		"X-Subscription-Token": {searchrequest.token},
	}
	// Send Request
	client := &http.Client{}
	repsonse, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer repsonse.Body.Close()
	// Recieve Request

	body, err := io.ReadAll(repsonse.Body)
	if err != nil {
		log.Fatal(err)
	}
	var res Response

	json.Unmarshal([]byte(body), &res)
	content := fmt.Sprintf("Query: %+v\n, Results: %+v\n", res.Query, res.Web.Results)
	bytes, _ := json.MarshalIndent(content, "", "")
	contentjsonerr := os.WriteFile("reply.json", bytes, 0644)
	if contentjsonerr != nil {
		log.Fatal(contentjsonerr)
	}
}

/*
curl -s --compressed "https://api.search.brave.com/res/v1/web/search?q=brave+search" -H "Accept: application/json"
    -H "Accept-Encoding: gzip" -H "X-Subscription-Token: BSAl9dM0lmw9MFeqKSgYPNmh8SBxPRp"
*/
