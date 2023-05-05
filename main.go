package main

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

/**
* This function counts how many unique normalized valid URLs were passed to the function
*
* Accepts a list of URLs
*
* Example:
*
* input: ['https://example.com']
* output: 1
*
* Notes:
*  - assume none of the URLs have authentication information (username, password).
*
* Normalized URL:
*  - process in which a URL is modified and standardized: https://en.wikipedia.org/wiki/URL_normalization
*
#    For example.
#    These 2 urls are the same:
#    input: ["https://example.com", "https://example.com/"]
#    output: 1
#
#    These 2 are not the same:
#    input: ["https://example.com", "http://example.com"]
#    output 2
#
#    These 2 are the same:
#    input: ["https://example.com?", "https://example.com"]
#    output: 1
#
#    These 2 are the same:
#    input: ["https://example.com?a=1&b=2", "https://example.com?b=2&a=1"]
#    output: 1
*/

func CountUniqueUrls(urls []string) int {
	uniqueUrls := make(map[string]bool)

	for _, urlStr := range urls {
		parsedUrl, err := url.Parse(urlStr)
		if err != nil {
			continue
		}

		// Normalize the URL
		parsedUrl.RawQuery = ""
		parsedUrl.Fragment = ""
		parsedUrl.Path = path.Clean(parsedUrl.Path)
		normalizedUrl := parsedUrl.String()

		// Count the unique normalized URLs
		if _, ok := uniqueUrls[normalizedUrl]; !ok {
			uniqueUrls[normalizedUrl] = true
		}
	}

	return len(uniqueUrls)
}

/**
 * This function counts how many unique normalized valid URLs were passed to the function per top level domain
 *
 * A top level domain is a domain in the form of example.com. Assume all top level domains end in .com
 * subdomain.example.com is not a top level domain.
 *
 * Accepts a list of URLs
 *
 * Example:
 *
 * input: ["https://example.com"]
 * output: Hash["example.com" => 1]
 *
 * input: ["https://example.com", "https://subdomain.example.com"]
 * output: Hash["example.com" => 2]
 *
 */

func CountUniqueUrlsPerTopLevelDomain(urls []string) map[string]int {
	uniqueUrls := make(map[string]bool)
	domains := make(map[string]int)

	for _, urlStr := range urls {
		parsedUrl, err := url.Parse(urlStr)
		if err != nil {
			continue
		}
		// Normalize the URL
		parsedUrl.RawQuery = ""
		parsedUrl.Fragment = ""
		normalizedUrl := parsedUrl.String()
		// Extract the top-level domain
		domainParts := strings.Split(parsedUrl.Hostname(), ".")
		if len(domainParts) > 1 {
			domain := domainParts[len(domainParts)-2] + "." + domainParts[len(domainParts)-1]
			// Count the unique normalized URLs per top-level domain
			if _, ok := uniqueUrls[normalizedUrl]; !ok {
				uniqueUrls[normalizedUrl] = true
				domains[domain]++
			}
		}
	}

	return domains
}

func main() {
	urls := []string{
		"https://example.com",
		"https://example.com/",
		"http://example.com",
		"https://example.com?a=1&b=2",
		"https://example.com?b=2&a=1",
		"https://subdomain.example.com",
		"https://subdomain.example.com/",
		"https://subdomain.example.com?a=1&b=2",
		"https://subdomain.example.com?b=2&a=1",
		"https://example.org",
		"https://example.net",
	}

	uniqueUrls := CountUniqueUrls(urls)
	fmt.Printf("Number of unique normalized URLs: %d\n", uniqueUrls)

	uniqueUrlsPerDomain := CountUniqueUrlsPerTopLevelDomain(urls)
	fmt.Printf("Number of unique normalized URLs per top level domain: %+v\n", uniqueUrlsPerDomain)
}
