package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseUrl string, pathPrefix string) http.HandlerFunc {

	target, err := url.Parse(targetBaseUrl)
	fmt.Println("Original target path:", target)

	if err != nil {
		fmt.Println("Error parsing target URL:", err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	fmt.Println("Original proxy path:", proxy)

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)

		originalPath := r.URL.Path
		fmt.Println("Original url path:", originalPath)

		strippedPath := strings.TrimPrefix(originalPath, pathPrefix)
		fmt.Println("strippedPath request path:", strippedPath)

		r.URL.Host = target.Host
		r.URL.Path = target.Path + strippedPath
		r.Host = target.Host

		fmt.Println("host:", r.URL.Host)
		fmt.Println("r host:", r.Host)
		fmt.Println("path:", r.URL.Path)
		fmt.Println("target host:", target.Host)
		fmt.Println("target Path:", target.Path)

		if userId, ok := r.Context().Value("userID").(string); ok {
			r.Header.Set("X-User-ID", userId)
		}

	}

	return proxy.ServeHTTP

}
