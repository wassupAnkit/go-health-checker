package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"go-health-checker/internal/checker"
	"go-health-checker/internal/config"
)

func main() {
	// Command-line flags
	cfgPath := flag.String("config", "services.json", "Path to config file")
	format := flag.String("format", "text", "Output format: 'text' or 'json'")
	flag.Parse()

	// Load URLs from config
	urls, err := config.Load(*cfgPath)
	if err != nil {
		fmt.Println("❌ Error loading config:", err)
		os.Exit(1)
	}

	fmt.Println("🔍 Checking URLs...")

	// Run checks concurrently with retries
	results := checker.CheckAll(urls.Services)

	// Output results in desired format
	if *format == "json" {
		jsonOutput, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("❌ Failed to encode JSON:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonOutput))
	} else {
		for _, r := range results {
			if r.Error == nil {
				fmt.Printf("✅ %s [%d]\n", r.URL, r.StatusCode)
			} else {
				fmt.Printf("❌ %s (%s)\n", r.URL, r.Error)
			}
		}
	}

	// Determine exit code for CI/CD
	hasError := false
	for _, r := range results {
		if r.Error != nil {
			hasError = true
			break
		}
	}
	if hasError {
		os.Exit(1) // At least one failure
	}
	os.Exit(0) // All passed
}
