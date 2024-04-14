package shortner

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"sync"
)

var (
	shortenedDomains = make(map[string]int)
	mu               sync.RWMutex
)

// This function records the domain of the original URL each time a URL is shortened.
// It takes the original URL as input
func recordDomain(url string) {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		return // Skip if URL format is invalid
	}
	domain := parts[2]
	mu.Lock()
	defer mu.Unlock()
	shortenedDomains[domain]++
}

// This function is added to retrieve the top 3 domains with the highest
// counts and return them as JSON.
func GetTopDomains(w http.ResponseWriter, r *http.Request) {
	type domainCount struct {
		Domain string `json:"domain"`
		Count  int    `json:"count"`
	}

	mu.RLock()
	defer mu.RUnlock()

	var domainCounts []domainCount
	/*
		It iterates over the shortenedDomains map and constructs a slice of domainCount structs,
		where each struct contains a domain and its corresponding count
	*/
	for domain, count := range shortenedDomains {
		domainCounts = append(domainCounts, domainCount{Domain: domain, Count: count})
	}

	// Sort domain counts by count in descending order
	sort.Slice(domainCounts, func(i, j int) bool {
		return domainCounts[i].Count > domainCounts[j].Count
	})

	// Take only top 3 domains
	if len(domainCounts) > 3 {
		domainCounts = domainCounts[:3]
	}

	// Marshal domain counts to JSON
	data, err := json.Marshal(domainCounts)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
