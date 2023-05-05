package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountUniqueUrls(t *testing.T) {
	tests := []struct {
		name     string
		urls     []string
		expected int
	}{
		{
			name:     "Empty input",
			urls:     []string{},
			expected: 0,
		},
		{
			name:     "All URLs are the same",
			urls:     []string{"https://example.com", "https://example.com", "https://example.com"},
			expected: 1,
		},
		{
			name:     "All URLs are different",
			urls:     []string{"https://example.com", "https://example.com?", "http://example.org", "ftp://example.net"},
			expected: 4,
		},
		{
			name:     "Different URLs, but normalized URLs are the same",
			urls:     []string{"https://example.com", "https://example.com/", "https://example.com/../././"},
			expected: 2,
		},
		{
			name:     "Different URLs wih params, but normalized URLs are the same",
			urls:     []string{"https://example.com?a=1&b=2", "https://example.com?b=2&a=1"},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountUniqueUrls(test.urls)
			if result != test.expected {
				t.Errorf("CountUniqueUrls(%v) returned %d, expected %d", test.urls, result, test.expected)
			}
		})
	}
}

func TestCountUniqueUrlsPerTopLevelDomain(t *testing.T) {
	tests := []struct {
		name        string
		urls        []string
		expectedMap map[string]int
	}{
		{
			name: "single URL with valid top level domain",
			urls: []string{
				"https://example.com",
			},
			expectedMap: map[string]int{
				"example.com": 1,
			},
		},
		{
			name: "multiple URLs with valid top level domain",
			urls: []string{
				"https://example.com",
				"https://subdomain.example.com",
				"https://example.com/page1",
				"https://example.com/page2",
				"https://anotherexample.com",
			},
			expectedMap: map[string]int{
				"example.com":        4,
				"anotherexample.com": 1,
			},
		},
		{
			name: "multiple URLs with invalid URLs mixed in",
			urls: []string{
				"https://example.com",
				"https://subdomain.example.com",
				"http://example.com",
				"http://example",
				"invalidurl",
				"https://example.com/page1",
				"https://example.com/page2",
				"https://anotherexample.com",
			},
			expectedMap: map[string]int{
				"example.com":        5,
				"anotherexample.com": 1,
			},
		},
		{
			name:        "empty list of URLs",
			urls:        []string{},
			expectedMap: map[string]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountUniqueUrlsPerTopLevelDomain(test.urls)
			assert.Equal(t, test.expectedMap, result)
		})
	}
}
