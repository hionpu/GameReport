package lol

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RiotClient struct {
	apiKey string
	httpClient *http.Client
	region string
	regionalRouting string
}

// Region configuration
var regionConfig = map[string]string{
	"kr":   "asia",      // Korea
	"jp1":  "asia",      // Japan
	"na1":  "americas",  // North America
	"br1":  "americas",  // Brazil
	"la1":  "americas",  // Latin America North
	"la2":  "americas",  // Latin America South
	"oc1":  "americas",  // Oceania
	"euw1": "europe",    // Europe West
	"eun1": "europe",    // Europe Nordic & East
	"tr1":  "europe",    // Turkey
	"ru":   "europe",    // Russia
}

func NewRiotClient(apiKey string) *RiotClient {
	return &RiotClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		}
		baseURL: 