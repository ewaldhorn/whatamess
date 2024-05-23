package main

import "net/url"

func escapeString(src string) string {
	return url.QueryEscape(src)
}

func escapeQueryParameters(p1 string, p2 string, p3 string) string {
	params := url.Values{}

	params.Add("param_1", p1)
	params.Add("param_2", p2)
	params.Add("param_3", p3)

	return params.Encode()
}

func allInOne(scheme string, host string, path string) string {
	finalURL := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}

	return finalURL.String()
}
