package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/wassupAnkit/go-health-checker/internal/checker"
	"github.com/wassupAnkit/go-health-checker/internal/config"
)

func main() {
	cfgPath := flag.String("config", "services.json", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	results := make(chan checker.Result)
	for _, url := range cfg.Services {
		go func(u string) {
			result := checker.Check(u)
			results <- result
		}(url)
	}

	// Collect results
	for i := 0; i < len(cfg.Services); i++ {
		res := <-results
		if res.Alive {
			fmt.Printf("✅ %s [%d]\n", res.URL, res.Status)
		} else {
			fmt.Printf("❌ %s (%v)\n", res.URL, res.Err)
		}
	}
}
