package google

type ResponseData struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind             string `json:"kind"`
		Title            string `json:"title"`
		HTMLTitle        string `json:"htmlTitle"`
		Link             string `json:"link"`
		DisplayLink      string `json:"displayLink"`
		Snippet          string `json:"snippet"`
		HTMLSnippet      string `json:"htmlSnippet"`
		CacheID          string `json:"cacheId"`
		FormattedURL     string `json:"formattedUrl"`
		HTMLFormattedURL string `json:"htmlFormattedUrl"`
		Pagemap          struct {
			CseThumbnail []struct {
				Width  string `json:"width"`
				Height string `json:"height"`
				Src    string `json:"src"`
			} `json:"cse_thumbnail"`
			Code []struct {
				Programminglanguage string `json:"programminglanguage"`
				Name                string `json:"name"`
				Coderepository      string `json:"coderepository"`
				Description         string `json:"description"`
			} `json:"code"`
			Organization []struct {
				Image string `json:"image"`
				URL   string `json:"url"`
			} `json:"organization"`
			Metatags []struct {
				Viewport                     string `json:"viewport"`
				FbAppID                      string `json:"fb:app_id"`
				TwitterImageSrc              string `json:"twitter:image:src"`
				TwitterSite                  string `json:"twitter:site"`
				TwitterCard                  string `json:"twitter:card"`
				TwitterTitle                 string `json:"twitter:title"`
				OgImage                      string `json:"og:image"`
				OgSiteName                   string `json:"og:site_name"`
				OgType                       string `json:"og:type"`
				OgTitle                      string `json:"og:title"`
				OgURL                        string `json:"og:url"`
				ProfileUsername              string `json:"profile:username"`
				BrowserStatsURL              string `json:"browser-stats-url"`
				BrowserErrorsURL             string `json:"browser-errors-url"`
				PjaxTimeout                  string `json:"pjax-timeout"`
				RequestID                    string `json:"request-id"`
				MsapplicationTileimage       string `json:"msapplication-tileimage"`
				MsapplicationTilecolor       string `json:"msapplication-tilecolor"`
				GoogleAnalytics              string `json:"google-analytics"`
				OctolyticsHost               string `json:"octolytics-host"`
				OctolyticsAppID              string `json:"octolytics-app-id"`
				OctolyticsDimensionRequestID string `json:"octolytics-dimension-request_id"`
				AnalyticsLocation            string `json:"analytics-location"`
				Dimension1                   string `json:"dimension1"`
				Hostname                     string `json:"hostname"`
				ExpectedHostname             string `json:"expected-hostname"`
				JsProxySiteDetectionPayload  string `json:"js-proxy-site-detection-payload"`
				HTMLSafeNonce                string `json:"html-safe-nonce"`
				FormNonce                    string `json:"form-nonce"`
			} `json:"metatags"`
			CseImage []struct {
				Src string `json:"src"`
			} `json:"cse_image"`
		} `json:"pagemap"`
	} `json:"items"`
}
