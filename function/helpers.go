package function

type JSFetchOptionsArg struct {
	Method         string
	Headers        map[string]string
	Body           string
	Mode           string
	Credentials    string
	Cache          string
	Redirect       string
	Referrer       string
	ReferrerPolicy string
	Integrity      string
	Keepalive      string
	Signal         string
}

type HTTPResponse struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

func NewJSFetcthOptionArg() JSFetchOptionsArg {
	defaultOptions := JSFetchOptionsArg{
		Method:         "GET",
		Headers:        make(map[string]string, 0),
		Body:           "",
		Mode:           "no-cors",
		Credentials:    "omit",
		Cache:          "no-cache",
		Redirect:       "error",
		Referrer:       "",
		ReferrerPolicy: "",
		Integrity:      "",
		Keepalive:      "",
		Signal:         "",
	}
	return defaultOptions
}
