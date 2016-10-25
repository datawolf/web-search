package google

import (
	"../userip"
	"context"
	"encoding/json"
	"net/http"
)

// Results is an ordered list of search results.
type Results []Result

// A Result contains the title and URL of a search result.
type Result struct {
	Title, URL string
}

// Search sends query to Google search and returns the results.
func Search(ctx context.Context, query string) (Results, error) {
	// Prepare the Google Search API request
	req, err := http.NewRequest("GET", "https://www.googleapis.com/customsearch/v1?", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("q", query)
	q.Set("cx", "002129777885552800241:v0akzp9nbnw")
	q.Set("key", "AIzaSyBfSPSMkvVBHbdWAMyQ27oLnivtzQseSlE")

	// If ctx is acrrying the user IP address. forward it to the searver.
	// Google APIs use the user IP to distinguish server-initiated requests
	// from end-user request
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userIp", userIP.String())
	}

	req.URL.RawQuery = q.Encode()

	// Issue the HTTP request and handle the response. The httpDo function
	// cancels the request if ctx.Done is closed.
	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Parse the JSON search result
		var data struct {
			Items struct {
				Results []struct {
					HtmlTitle        string
					HtmlFormattedUrl string
				}
			}
		}

		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}

		for _, res := range data.Items.Results {
			results = append(results, Result{Title: res.HtmlTitle, URL: res.HtmlFormattedUrl})
		}

		return nil
	})

	// httpDo waits for the closure we provided to return. so it's safe to read results here.
	return results, err
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return
		return ctx.Err()
	case err := <-c:
		return err
	}

}
