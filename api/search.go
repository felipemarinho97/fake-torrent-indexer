package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/felipemarinho97/fake-torrent-indexer/misc"
	"github.com/felipemarinho97/fake-torrent-indexer/schema"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	// get "q" query param
	query := r.URL.Query().Get("q")

	// Initialize random generator for this request
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	results := []schema.IndexedTorrent{}
	// Generate 20-25 results by cycling through magnet links
	numResults := 20 + rng.Intn(6) // 20-25 results
	for i := 0; i < numResults; i++ {
		// Cycle through available magnet links
		magnet := misc.Magnets[i%len(misc.Magnets)]
		title := generateReleaseTitle(query, i)

		// Generate high seed counts for better Prowlarr attraction
		baseSeedCount := 50 + rng.Intn(500) // 50-550 seeds
		baseLeechCount := rng.Intn(50)      // 0-50 leechers

		results = append(results, schema.IndexedTorrent{
			Title:         fmt.Sprintf("%s (brazilian, eng)", title),
			OriginalTitle: title,
			Details:       "This is for testing purposes. Please, deploy your torrent-indexer (https://github.com/felipemarinho97/torrent-indexer) instance.",
			Year:          getRandomYear(),
			IMDB:          fmt.Sprintf("tt%07d", rng.Intn(9999999)),
			Audio:         []schema.Audio{schema.AudioEnglish, schema.AudioPortuguese},
			MagnetLink:    magnet,
			Date:          time.Now().Add(-time.Duration(rng.Intn(72)) * time.Hour), // Random dates within last 3 days
			InfoHash:      generateRandomInfoHash(rng),
			Trackers: []string{
				"udp://tracker.opentrackr.org:1337/announce",
				"udp://tracker.coppersurfer.tk:6969/announce",
			},
			Size:       getRandomSize(i),
			Files:      nil,
			LeechCount: baseLeechCount,
			SeedCount:  baseSeedCount,
			Similarity: 1.0 - float32(i)*0.02, // Decrease similarity more gradually
		})
	}

	// shuffle results to add randomness
	rng.Shuffle(len(results), func(i, j int) {
		results[i], results[j] = results[j], results[i]
	})

	resp := schema.Response{
		Results: results,
		Count:   len(results),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func getRandomSize(i int) string {
	sizes := []string{
		"700 MB",
		"1.4 GB",
		"2.1 GB",
		"3.5 GB",
		"4.7 GB",
		"8.5 GB",
		"12.3 GB",
		"25.7 GB",
	}
	return sizes[i%len(sizes)]
}

func generateReleaseTitle(query string, seed int) string {
	// Initialize random with seed for some consistency
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(seed)))

	// Clean the query and use it as base title
	baseTitle := strings.TrimSpace(query)
	if baseTitle == "" {
		baseTitle = "Unknown"
	}

	// Simple title case conversion (first letter of each word uppercase)
	words := strings.Fields(baseTitle)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	baseTitle = strings.Join(words, " ")

	// Check if query already contains a 4-digit year (1900-2099)
	yearRegex := regexp.MustCompile(`\b(19|20)\d{2}\b`)
	hasYear := yearRegex.MatchString(query)

	// Random qualities
	qualities := []string{"720p", "1080p", "2160p", "4K"}
	sources := []string{"BluRay", "WEBRip", "WEB-DL", "DVDRip", "BDRip", "HDTV"}
	codecs := []string{"x264", "x265", "H.264", "H.265", "HEVC"}
	audioCodecs := []string{"DTS", "AC3", "AAC", "TrueHD", "DDP5.1"}
	releaseGroups := []string{"YIFY", "RARBG", "FGT", "SPARKS", "AMZN", "NTb", "CMRG", "EVO", "KOGi", "BATV", "Tigole"}

	// Random episode info for TV shows (sometimes)
	episodeInfo := ""
	if r.Intn(3) == 0 { // 33% chance of being a TV episode
		season := r.Intn(10) + 1
		episode := r.Intn(24) + 1
		episodeInfo = fmt.Sprintf(".S%02dE%02d", season, episode)
	}

	// Only add year if query doesn't already contain one
	yearStr := ""
	if !hasYear {
		year := 2010 + r.Intn(15)
		yearStr = fmt.Sprintf(".%d", year)
	}

	// Build the release name
	title := fmt.Sprintf("%s%s%s.%s.%s.%s-%s",
		strings.ReplaceAll(baseTitle, " ", "."),
		episodeInfo,
		yearStr,
		qualities[r.Intn(len(qualities))],
		sources[r.Intn(len(sources))],
		codecs[r.Intn(len(codecs))],
		releaseGroups[r.Intn(len(releaseGroups))],
	)

	// Sometimes add audio codec
	if r.Intn(2) == 0 {
		parts := strings.Split(title, "-")
		if len(parts) >= 2 {
			title = strings.Join(parts[:len(parts)-1], "-") + "." + audioCodecs[r.Intn(len(audioCodecs))] + "-" + parts[len(parts)-1]
		}
	}

	return title
}

func getRandomYear() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	years := []string{"2020", "2021", "2022", "2023", "2024", "2025"}
	return years[r.Intn(len(years))]
}

func generateRandomInfoHash(r *rand.Rand) string {
	hexChars := "0123456789abcdef"
	hash := make([]byte, 40)
	for i := range hash {
		hash[i] = hexChars[r.Intn(len(hexChars))]
	}
	return string(hash)
}
