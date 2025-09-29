# Fake Torrent Indexer

A lightweight mock torrent indexer server designed for testing and development purposes. This server simulates a real torrent indexer by generating realistic torrent search results with proper metadata, making it perfect for testing applications like Prowlarr, Sonarr, Radarr, and other *arr stack applications.

## Features

- 🔍 **Wildcard Indexer Support**: Works with any indexer name in the URL path
- 🎬 **Realistic Release Names**: Generates authentic-looking torrent titles with proper formatting (BluRay, WEB-DL, x264, etc.)
- 📊 **High Seed Counts**: Provides well-seeded torrents to enhance Prowlarr selection likelihood
- 🔄 **Multiple Results**: Returns 20+ varied results per search query
- 📝 **Comprehensive Metadata**: Includes IMDB IDs, audio tracks, file sizes, and tracker information
- 📡 **HTTP Logging**: Built-in request/response logging with zerolog
- 🌐 **Multiple Endpoints**: Supports various URL patterns for maximum compatibility

## Quick Start

### Prerequisites

- Go 1.24.5 or later

### Installation

1. Clone the repository:
```bash
git clone https://github.com/felipemarinho97/fake-torrent-indexer.git
cd fake-torrent-indexer
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

The server will start on port `7707`.

## API Endpoints

### Search Endpoints

The server supports multiple endpoint patterns for maximum compatibility:

- `GET /indexers/{indexer_name}?q={search_query}`
- `GET /indexers?q={search_query}`
- `GET /search?q={search_query}`

### Index Endpoint

- `GET /` - Returns basic server information

### Example Requests

```bash
# Search with specific indexer name
curl "http://localhost:7707/indexers/rarbg?q=Avengers"

# Search without indexer name
curl "http://localhost:7707/indexers?q=Breaking%20Bad"

# Direct search endpoint
curl "http://localhost:7707/search?q=The%20Matrix"
```

## Response Format

The server returns JSON responses with the following structure:

```json
{
  "results": [
    {
      "title": "The.Matrix.1999.1080p.BluRay.x264-SPARKS",
      "original_title": "The.Matrix.1999.1080p.BluRay.x264-SPARKS",
      "details": "Your IP is blocked. Please, deploy your torrent-indexer instance.",
      "year": "2023",
      "imdb": "tt0133093",
      "audio": ["english", "portuguese"],
      "magnet_link": "magnet:?xt=urn:btih:...",
      "date": "2025-09-29T10:30:00Z",
      "info_hash": "1234567890abcdef1234567890abcdef12345678",
      "trackers": [
        "udp://tracker.opentrackr.org:1337/announce",
        "udp://tracker.coppersurfer.tk:6969/announce",
        "https://academictorrents.com/announce.php"
      ],
      "size": "8.5 GB",
      "files": null,
      "leech_count": 15,
      "seed_count": 342,
      "similarity": 1.0
    }
  ],
  "count": 23
}
```

## Configuration

### Port Configuration

The server runs on port `7707` by default. To change the port, modify the `main.go` file:

```go
err := http.ListenAndServe(":8080", loggedIndexerMux) // Change to desired port
```

### Customizing Results

You can customize the generated torrent data by modifying the following files:

- `api/search.go` - Main search logic and result generation
- `misc/constants.go` - Magnet links and constants
- `schema/` - Data structures and audio configurations

## Use Cases

### Development Testing
Perfect for testing torrent client applications without relying on real indexers.

### Prowlarr Testing
Provides realistic, well-seeded results that Prowlarr will prioritize, making it ideal for testing Prowlarr configurations.

### *arr Stack Development
Great for testing Sonarr, Radarr, and other applications that integrate with torrent indexers.

### Load Testing
Use for performance testing of applications that consume torrent indexer APIs.

## Project Structure

```
├── api/           # HTTP handlers and API logic
├── logging/       # HTTP logging middleware
├── misc/          # Constants and utility data
├── schema/        # Data structures and types
├── main.go        # Server entry point
├── go.mod         # Go module definition
└── README.md      # This file
```

## Generated Content Features

- **Release Names**: Authentic torrent naming conventions with quality indicators
- **Episode Information**: Random season/episode data for TV shows
- **Quality Variations**: 720p, 1080p, 2160p, 4K options
- **Source Types**: BluRay, WEB-DL, WEBRip, DVDRip, BDRip, HDTV
- **Codecs**: x264, x265, H.264, H.265, HEVC
- **Audio**: DTS, AC3, AAC, TrueHD, DDP5.1
- **Release Groups**: YIFY, RARBG, FGT, SPARKS, AMZN, NTb, CMRG, EVO

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Disclaimer

This software is intended for testing and development purposes only. It generates fake torrent data and should not be used to distribute actual copyrighted content. Users are responsible for complying with all applicable laws and regulations.

## Acknowledgments

- Built with Go's standard HTTP library
- Logging powered by [zerolog](https://github.com/rs/zerolog)
- Inspired by the need for reliable torrent indexer testing tools