package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yael-castro/godi/internal/dependency"
	"github.com/yael-castro/godi/internal/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.SetFlags(log.Flags() | log.Lshortfile)

	h := handler.New()

	dependency.NewInjector(dependency.Default).Inject(h)

	log.Printf(`http server is running on port "%v" %v`, port, "🤘\n")
	log.Fatal(http.ListenAndServe(":"+port, h))
}
