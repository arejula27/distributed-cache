package main

type server struct {
	cache      Cache
	listenPort string
}

// NewServer creates a server
func NewServer() *server {
	server := server{
		listenPort: "8000",
		cache:      NewMemoryCache(),
	}
	return &server
}
