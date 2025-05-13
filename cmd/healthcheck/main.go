package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/wassupAnkit/go-health-checker/internal/checker"
	"github.com/wassupAnkit/go-health-checker/internal/config"
)

func main() {
	// Parse command-line flags
	cfgPath := flag.String("config", "services.json", "Path to config file")
	flag.Parse()

	// Load URLs from the config file
	urls, err := config.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println("🔍 Checking URLs...")

	// Run all URL checks concurrently
	results := checker.CheckAll(urls)

	// Print results
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("❌ %s (%v)\n", result.URL, result.Error)
		} else {
			fmt.Printf("✅ %s [%d]\n", result.URL, result.StatusCode)
		}
	}
}
