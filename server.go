package main

// Server handles the cache connections
type Server struct {
	cache      Cache
	listenPort string
}

// NewServer creates a server
func NewServer() *Server {
	server := Server{
		listenPort: "8000",
		cache:      NewMemoryCache(),
	}
	return &server
}
