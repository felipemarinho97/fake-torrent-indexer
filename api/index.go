package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"time":  currentTime,
		"build": "v1.0.0",
		"endpoints": map[string]interface{}{
			"/indexers/comando_torrents": []map[string]interface{}{
				{
					"method":      "GET",
					"description": "Indexer for comando torrents",
					"query_params": map[string]string{
						"q":              "search query",
						"page":           "page number",
						"filter_results": "if results with similarity equals to zero should be filtered (true/false)",
					},
				},
			},
			"/indexers/bludv": []map[string]interface{}{
				{
					"method":      "GET",
					"description": "Indexer for bludv",
					"query_params": map[string]string{
						"q":              "search query",
						"page":           "page number",
						"filter_results": "if results with similarity equals to zero should be filtered (true/false)",
					}},
			},
			"/indexers/torrent-dos-filmes": []map[string]interface{}{
				{
					"method":      "GET",
					"page":        "page number",
					"description": "Indexer for Torrent dos Filmes",
					"query_params": map[string]string{
						"q":              "search query",
						"filter_results": "if results with similarity equals to zero should be filtered (true/false)",
					},
				},
			},
			"/indexers/comandohds": []map[string]interface{}{
				{
					"method":      "GET",
					"page":        "page number",
					"description": "Indexer for Comando HDs",
					"query_params": map[string]string{
						"q":              "search query",
						"filter_results": "if results with similarity equals to zero should be filtered (true/false)",
					},
				},
			},
			"/indexers/starck-filmes": []map[string]interface{}{
				{
					"method":      "GET",
					"page":        "page number",
					"description": "Indexer for Starck Filmes",
					"query_params": map[string]string{
						"q":              "search query",
						"filter_results": "if results with similarity equals to zero should be filtered (true/false)",
					},
				},
			},
			"/indexers/rede_torrent": []map[string]interface{}{
				{
					"method":      "GET",
					"description": "Indexer for rede torrent",
					"query_params": map[string]string{
						"q":              "search query",
						"page":           "page number",
						"filter_results": "if results with similarity equals to zero should be filtered (true/false)",
					},
				},
			},
			"/indexers/manual": []map[string]interface{}{
				{
					"method":      "POST",
					"description": "Add a manual torrent entry to the indexer for 12 hours",
					"body": map[string]interface{}{
						"magnetLink": "magnet link",
					}},
				{
					"method":      "GET",
					"description": "Get all manual torrents",
				},
			},
			"/search": []map[string]interface{}{
				{
					"method":      "GET",
					"description": "Search for cached torrents across all indexers",
					"query_params": map[string]string{
						"q": "search query",
					},
				},
			},
			"/ui/": []map[string]interface{}{
				{
					"method":      "GET",
					"description": "Show the unified search UI (only work if Meilisearch is enabled)",
				},
			},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
