package _package

import (
	"fmt"
	"google.golang.org/api/urlshortener/v1"
	"net/http"
	"text/template"
)

func Main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/short", short)
	http.HandleFunc("/long", long)

	http.ListenAndServe("localhost:8080", nil)
}

var rootTemplate = template.Must(template.New("rootHtml").Parse(`
<html><body>
<h1>URL SHORTENER</h1>
{{if .}}{{.}}<br /><br />{{end}}
<form action="/short" type="POST">
Shorten this: <input type="text" name="longUrl" />
<input type="submit" value="Give me the short URL" />
</form>
<br />
<form action="/long" type="POST">
Expand this: http://goo.gl/<input type="text" name="shortUrl" />
<input type="submit" value="Give me the long URL" />
</form>
</body></html>
`))

func root(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, r)
}

func short(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("longUrl")
	urlShortenSrv, _ := urlshortener.New(http.DefaultClient)
	url, _ := urlShortenSrv.Url.Insert(&urlshortener.Url{LongUrl: longUrl}).Do()
	rootTemplate.Execute(w, fmt.Sprintf("Shortened version of %s is : %s", longUrl, url.Id))
}

func long(w http.ResponseWriter, r *http.Request) {
	shortUrl := "http://goo.gl/" + r.FormValue("shortUrl")
	urlShortenSrv, _ := urlshortener.New(http.DefaultClient)
	url, err := urlShortenSrv.Url.Get(shortUrl).Do()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	rootTemplate.Execute(w, fmt.Sprintf("Long version of %s is : %s", shortUrl, url.LongUrl))
}
