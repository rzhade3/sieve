package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/rzhade3/sieve/src"
)

func main() {
	// Get files config as a CLI arg
	var filesConfig string
	var domainList string
	flag.StringVar(&filesConfig, "files", "", "Files config file")
	flag.StringVar(&domainList, "domains", "", "Domain list file")

	flag.Parse()

	if filesConfig == "" {
		fmt.Println("No files config file specified")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if domainList == "" {
		fmt.Println("No domain list file specified")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Read files config
	files, err := src.ReadFilesConfigYaml(filesConfig)
	if err != nil {
		fmt.Printf("Could not read files config: %s", err)
		os.Exit(1)
	}

	// Read domain list
	domains, err := src.ReadFileLineByLine(domainList)
	if err != nil {
		fmt.Printf("Could not read domain list: %s", err)
		os.Exit(1)
	}

	// Create four threads
	//
	threads := 4
	domains_chan := make(chan string, threads)
	results_chan := make(chan src.DomainResult, threads)

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			for domain := range domains_chan {
				// domain_result := pkg.DomainResult{
				// 	Domain: domain,
				// }
				domain_result := src.DomainResult{
					Domain: domain,
					Files:  src.CheckDomain(domain, &files),
				}
				results_chan <- domain_result
			}
			wg.Done()
		}()
	}

	for _, domain := range domains {
		domains_chan <- domain
	}
	close(domains_chan)
	wg.Wait()
	close(results_chan)

	for result := range results_chan {
		fmt.Println(result.ToString())
	}
}
