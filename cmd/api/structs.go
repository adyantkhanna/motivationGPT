package main

// type client struct {
// 	baseURL       string
// 	apiKey        string
// 	userAgent     string
// 	httpClient    *http.Client
// 	defaultEngine string
// 	idOrg         string
// }

type GPTResponse struct {
	Response string `json:"Response"`
	Good     bool   `json:"Good"`
}
