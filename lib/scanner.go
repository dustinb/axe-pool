package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

const DHCP_START = 1
const DHCP_END = 254

func GetSystemInfo(ip string) Bitaxe {
	client := http.Client{}
	client.Timeout = 2 * time.Second
	resp, err := client.Get("http://" + ip + "/api/system/info")
	if err != nil {
		return Bitaxe{}
	}
	body, _ := io.ReadAll(resp.Body)
	info := Bitaxe{}
	json.Unmarshal(body, &info)
	return info
}

// ScanNetwork scans the network for the Bitaxe
func ScanNetwork() []Bitaxe {
	var bitaxes = make(map[string]Bitaxe)
	var mutex sync.Mutex
	var waitgroup sync.WaitGroup

	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		host, _ := address.(*net.IPNet)
		if host.IP.IsLoopback() {
			continue
		}
		// Check for IPv4
		if host.IP.To4() == nil {
			continue
		}
		fmt.Printf("Scanning %s for Bitaxes...\n", host.IP.String())
		octets := strings.Split(host.IP.String(), ".")
		network := octets[0] + "." + octets[1] + "." + octets[2] + ".%d"

		for i := DHCP_START; i < DHCP_END+1; i++ {
			ip := fmt.Sprintf(network, i)
			waitgroup.Add(1)
			go func() {
				defer waitgroup.Done()
				bitaxe := GetSystemInfo(ip)
				if bitaxe.Hostname != "" {
					bitaxe.IP = ip
					mutex.Lock()
					bitaxes[ip] = bitaxe
					mutex.Unlock()
				}
			}()
		}
	}
	waitgroup.Wait()
	final := []Bitaxe{}
	for _, bitaxe := range bitaxes {
		final = append(final, bitaxe)
	}
	return final
}
