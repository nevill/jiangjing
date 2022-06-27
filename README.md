# JiangJing
`JiangJiang` is a Go client for elastic enterprise search. The implementation is inspired by [go-elasticsearch](https://github.com/elastic/go-elasticsearch/).
> The name 蒋敬 (jiǎng jìng) is a character from classical Chinese novel `Shui Hu Zhuan`.


## How to use it
A complete example:

```go
package main

import (
	"log"
	"net/http"

	jj "github.com/nevill/jiangjiang"
)

func main() {
	client, err := jj.NewClient(jj.Config{
		Address:  "http://localhost:3002",
		Username: "elastic",
		Password: "elastic823",
	})

	if err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}

	{
		response, err := client.Health()
		if err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}

		if response.StatusCode != http.StatusOK {
			log.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, response.StatusCode)
		}

		log.Printf("response is: %s\n", response)
	}

	{
		response, err := client.AppSearch.Engines.List()
		if err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}

		if response.StatusCode != http.StatusOK {
			log.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, response.StatusCode)
		}

		log.Printf("response is: %s\n", response)
	}

	name := "test-engine"

	{
		response, err := client.AppSearch.Engines.Create(name)
		if err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}

		if response.StatusCode != http.StatusOK {
			log.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, response.StatusCode)
		}

		log.Printf("response is: %s\n", response)
	}

	{
		response, err := client.AppSearch.Engines.Delete(name)
		if err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}

		if response.StatusCode != http.StatusOK {
			log.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, response.StatusCode)
		}

		log.Printf("response is: %s\n", response)
	}
}
```
