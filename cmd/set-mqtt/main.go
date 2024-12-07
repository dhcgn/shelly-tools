package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

type ShellyResponse struct {
	Type string `json:"type"`
}

func isShellyDevice(ip string) bool {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("http://%s/shelly", ip))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var shellyResp ShellyResponse
	if err := json.NewDecoder(resp.Body).Decode(&shellyResp); err != nil {
		return false
	}
	return shellyResp.Type != ""
}

func main() {
	var wg sync.WaitGroup
	ipRange := make([]string, 0, 254)
	for i := 1; i <= 254; i++ {
		ipRange = append(ipRange, fmt.Sprintf("192.168.3.%d", i))
	}

	shellyDevices := make([]string, 0)
	var mu sync.Mutex

	for _, ip := range ipRange {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:80", ip), 1*time.Second)
			if err == nil {
				conn.Close()
				if isShellyDevice(ip) {
					mu.Lock()
					shellyDevices = append(shellyDevices, ip)
					mu.Unlock()
				}
			}
		}(ip)
	}

	wg.Wait()

	for _, ip := range shellyDevices {
		fmt.Println(ip)
	}
}
