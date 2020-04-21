package main

import (
	"context"
	"fmt"

	"github.com/ns3777k/go-shodan/shodan"
)

func main() {
	client := shodan.NewClient(nil, "enter_api_key_here")
	host := &shodan.HostQueryOptions{
		Query: "product:MySQL country:TR",
		Page:  480,
	}

	querySearch, err := client.GetHostsForQuery(context.Background(), host)

	if err != nil {
		panic(err)
	}

	fmt.Print("Total: ")
	fmt.Println(querySearch.Total)
	matches := querySearch.Matches

	fmt.Println(len(matches))

	for i := 0; i < len(matches); i++ {
		ipStr := matches[i].IP
		port := matches[i].Port
		version := matches[i].Version
		fmt.Printf("IP Address: %s\nPort: %d\nVersion: %s\n\n", ipStr, port, version)
	}

}
