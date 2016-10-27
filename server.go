// The server program issue Google search requests and demonstrates the use
// of the go context API. It servers on port 8004
//
// The /search endpoint accepts these query params:
//	q=the google search query
//	timeout=a timeout for the request, in time.Duration format
//
// For example, http://localhost:8000/search?q=golang&timeout=1s servers the
// first few google search results for "golang" or a "deadline exceeded"
// error if the timeout expires.
package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"./google"
	"./userip"
)

func main() {
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/", showIndex)
	log.Fatal(http.ListenAndServe(":8004", nil))
}

// showIndex return the index to client
func showIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // normal html header
	w.WriteHeader(http.StatusOK)                               // Write http status code
	w.Write([]byte(index))
}

// handleSearch handles URLs like /search>q=golang&timeout=1s by forwarding the
// query to google.Search. If the query param includes timeout, the search is
// canceled after that duration elapses
func handleSearch(w http.ResponseWriter, req *http.Request) {
	// ctx is the Context for this handler. Calling cancel closes the ctx.Done
	// channel, which is the cancellation signal for requests started by this handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Cancel ctx as soon as handleSearch returns

	// check the search query
	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	// Store the user IP in ctx for use by code in other packages
	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)

	// Run the Google search and print the request
	start := time.Now()
	results, err := google.Search(ctx, query)
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := resultsTemplate.Execute(w, struct {
		Results          google.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<body>
<ol>
{{range .Results}}
<li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
{{end}}
</ol>
<p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))

var index = `
<html>
  <head>
	<title> 自定义搜索引擎--Linux相关 </title>
  </head>
<body>
  <h2> 自定义搜索引擎--Linux相关</h2>
  <form action="/search">
    <input type="text" name="q"> <input type="submit" value="搜一下">
  </form>
</body>
</html>
`
