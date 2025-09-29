package schema

type Response struct {
	Results []IndexedTorrent `json:"results"`
	Count   int              `json:"count"`
}
