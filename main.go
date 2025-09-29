package main

import (
	"net/http"

	"github.com/felipemarinho97/fake-torrent-indexer/api"
	"github.com/felipemarinho97/fake-torrent-indexer/logging"
)

func main() {
	logging.InitLogger()

	indexerMux := http.NewServeMux()

	// Handle any indexer name with wildcard pattern
	indexerMux.HandleFunc("/", api.HandlerIndex)
	indexerMux.HandleFunc("/indexers/", api.HandleSearch)
	indexerMux.HandleFunc("/indexers", api.HandleSearch)
	indexerMux.HandleFunc("/search", api.HandleSearch)

	loggedIndexerMux := logging.HTTPLoggingMiddleware(indexerMux)

	err := http.ListenAndServe(":7707", loggedIndexerMux)
	if err != nil {
		panic(err)
	}
}
