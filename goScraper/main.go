package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type javData struct {
	Actresses       []actresses  `json:"actresses"`
	Categories      []categories `json:"categories"`
	Content_id      string       `json:"content_id"`
	Director        string       `json:"director"`
	Images          images       `json:"images"`
	Label           label        `json:"label"`
	Maker           maker        `json:"maker"`
	Release_date    string       `json:"release_date"`
	Runtime_minutes string       `json:"runtime_minutes"`
	Sample          sample       `json:"sample"`
	Series          series       `json:"series"`
	Title           string       `json:"title"`
}

type actresses struct {
	Name string `json:"name"`
}

type categories struct {
	Name string `json:"name"`
}

type images struct {
	Jacket_image jacket_image `json:"jacket_image"`
}

type jacket_image struct {
	Large  string `json:"large"`
	Large2 string `json:"large2"`
}

type label struct {
	Name string `json:"name"`
}

type maker struct {
	Name string `json:"name"`
}

type sample struct {
	High string `json:"high"`
}

type series struct {
	Name       string `json:"name"`
	Series_url string `json:"series_url"`
}

func main() {
	c := colly.NewCollector()

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	str := "https://r18.dev/videos/vod/movies/detail/-/dvd_id=SUN-084/json"

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", str)
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(">>", string(r.Body))
		str := string(r.Body)
		res := javData{}
		json.Unmarshal([]byte(str), &res)

		fmt.Println("****", res)
		fmt.Println("****", res.Series)
		fmt.Println("****", res.Series.Name)
		fmt.Println("****", res.Series.Series_url)
	})

	c.Visit(str)

	// str := `{"page": 1, "fruits": ["apple", "peach"]}`
	// res := response2{}
	// json.Unmarshal([]byte(str), &res)
	// fmt.Println(res)
	// fmt.Println(res.Fruits[0])
}
