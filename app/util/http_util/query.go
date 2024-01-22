package httputil

import "strings"

type httpQueryParams struct {
	query map[string]string
}

func NewHTTPQuery() *httpQueryParams {
	return &httpQueryParams{
		query: make(map[string]string),
	}
}

func (q *httpQueryParams) SetQuery(key, value string) {
	q.query[key] = value
}

func (q *httpQueryParams) Reset() bool {
	if len(q.query) == 0 {
		return false
	}
	q.query = make(map[string]string)
	return true
}

func (q *httpQueryParams) QueryStringURL() string {
	if len(q.query) == 0 {
		return ""
	}
	var query string = "?"
	for k, v := range q.query {
		query += k + "=" + v + "&"
	}

	return strings.TrimSuffix(query, "&")
}
