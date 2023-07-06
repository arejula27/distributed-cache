// Entrypoint fot the cache
package main

import "log"

func main() {

	s := NewServer()
	log.Fatalln(s.Start())
}
