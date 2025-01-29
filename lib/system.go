package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Patch struct {
	StratumURL          string `json:"stratumURL"`
	StratumPort         int    `json:"stratumPort"`
	StratumUser         string `json:"stratumUser"`
	FallbackStratumURL  string `json:"fallbackStratumURL"`
	FallbackStratumPort int    `json:"fallbackStratumPort"`
	FallbackStratumUser string `json:"fallbackStratumUser"`
}

func PatchAxe(ip string, patch Patch) {
	client := http.Client{}
	json, _ := json.Marshal(patch)
	fmt.Printf("Patching Axe: %s\n", ip)
	req, _ := http.NewRequest("PATCH", "http://"+ip+"/api/system", bytes.NewBuffer(json))
	client.Do(req)
	time.Sleep(2 * time.Second)
	client.Post("http://"+ip+"/api/system/restart", "application/json", nil)
}
