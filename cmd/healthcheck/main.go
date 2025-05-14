package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wassupAnkit/go-health-checker/internal/checker"
	"github.com/wassupAnkit/go-health-checker/internal/config"
)

func main() {
	// Parse command-line flags
	cfgPath := flag.String("config", "services.json", "Path to config file")
	format := flag.String("format", "text", "Output format: text or json")
	flag.Parse()

	// Load URLs from the config file
	urls, err := config.Load(*cfgPath)
	if err != nil {
		fmt.Println("❌ Error loading config:", err)
		os.Exit(1)
	}

	fmt.Println("🔍 Checking URLs...")

	// Run health checks with retries
	results := checker.CheckAll(urls.Services, 3)

	// Display results
	if *format == "json" {
		output, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("❌ Failed to marshal JSON:", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
	} else {
		for _, result := range results {
			if result.Error != nil {
				fmt.Printf("❌ %s (%v)\n", result.URL, result.Error)
			} else {
				fmt.Printf("✅ %s [%d]\n", result.URL, result.StatusCode)
			}
		}
	}
}
