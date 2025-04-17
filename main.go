package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type result struct {
	statusCode int
	err        error
}

type reportData struct {
	URL           string
	TotalRequests int
	Success200    int
	StatusCodes   map[int]int
	Errors        int
	TotalTime     string
	Timestamp     string
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("Erro: url é obrigatório")
		flag.Usage()
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	results := make(chan result, *requests)
	var wg sync.WaitGroup
	sem := make(chan struct{}, *concurrency)
	start := time.Now()

	for i := 0; i < *requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			resp, err := client.Get(*url)
			if err != nil {
				results <- result{err: err}
				return
			}
			results <- result{statusCode: resp.StatusCode}
			resp.Body.Close()
		}()
	}

	wg.Wait()
	close(results)

	totalTime := time.Since(start)
	statusCount := make(map[int]int)
	errorCount := 0

	for res := range results {
		if res.err != nil {
			errorCount++
		} else {
			statusCount[res.statusCode]++
		}
	}

	report := reportData{
		URL:           *url,
		TotalRequests: *requests,
		Success200:    statusCount[200],
		StatusCodes:   statusCount,
		Errors:        errorCount,
		TotalTime:     totalTime.String(),
		Timestamp:     time.Now().Format(time.RFC3339),
	}

	printReport(report)
}

func printReport(r reportData) {
	fmt.Println("\n====== RELATÓRIO DE TESTE DE CARGA ======")
	fmt.Printf("URL testada: %s\n", r.URL)
	fmt.Printf("Tempo total: %s\n", r.TotalTime)
	fmt.Printf("Total de requests: %d\n", r.TotalRequests)
	fmt.Printf("Respostas com status 200: %d\n", r.Success200)
	fmt.Println("Outros códigos de status:")
	for code, count := range r.StatusCodes {
		if code != 200 {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}
	if r.Errors > 0 {
		fmt.Printf("Erros de requisição: %d\n", r.Errors)
	}
}
