package opds

import (
	"fmt"
	"net/http"
	"os"
)

// Server represents the OPDS server.
type Server struct {
	// Add any configuration or dependencies here
}

// NewServer creates a new OPDS server.
func NewServer() *Server {
	return &Server{}
}

// Start starts the HTTP server on the given address.
func (s *Server) Start(addr string) error {
	// Use PORT environment variable if available
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	
	http.HandleFunc("/opds", s.handleRootCatalog)
	http.HandleFunc("/opds/search", s.handleSearch)
	http.HandleFunc("/opds/book/", s.handleBookDownload)
	fmt.Printf("OPDS server listening on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}

// handleRootCatalog serves the root OPDS catalog feed.
func (s *Server) handleRootCatalog(w http.ResponseWriter, r *http.Request) {
	// Generate and return OPDS root catalog XML
	w.Header().Set("Content-Type", "application/atom+xml;profile=opds-catalog;kind=acquisition")
	w.Write([]byte(GenerateRootCatalogFeed()))
}

// handleSearch serves the OPDS search endpoint.
func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	// Generate and return OPDS search results XML
	w.Header().Set("Content-Type", "application/atom+xml;profile=opds-catalog;kind=acquisition")
	w.Write([]byte(GenerateSearchResultsFeed(query)))
}

// handleBookDownload serves the EPUB file for the given book ID.
func (s *Server) handleBookDownload(w http.ResponseWriter, r *http.Request) {
	// Extract book ID from the URL path
	// For simplicity, we assume the URL is /opds/book/:id
	// In a real implementation, you would parse the URL and validate the ID
	bookID := "1" // Placeholder for book ID extraction
	if bookID != "1" {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	// TODO: In a real implementation, fetch the EPUB file for the given book ID
	// For now, serve a sample EPUB file
	w.Header().Set("Content-Type", "application/epub+zip")
	w.Header().Set("Content-Disposition", "attachment; filename=sample.epub")
	// Sample EPUB content (replace with actual file serving logic)
	w.Write([]byte("Sample EPUB content"))
} 