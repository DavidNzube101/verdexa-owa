package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "strings"
	"time"
)

const (
	DUNE_API_KEY = "0tef04bg4lx9drtbjwj5jxlmzgfvep0d"
	DUNE_API_BASE = "https://api.dune.com/api/echo/beta"
)

// Types representing the data structures from dashboard-api.ts
type TransactionFlowData struct {
	Nodes []Node `json:"nodes"`
	Links []Link `json:"links"`
}

type Node struct {
	ID    string  `json:"id"`
	Label string  `json:"label"`
	Size  int     `json:"size"`
	Color string  `json:"color"`
	Type  string  `json:"type,omitempty"`
	Cluster string `json:"cluster,omitempty"`
}

type Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Value  int    `json:"value"`
}

type AnomalyData struct {
	Dates     []string   `json:"dates"`
	Values    []float64  `json:"values"`
	Anomalies []Anomaly  `json:"anomalies"`
}

type Anomaly struct {
	Date       string  `json:"date"`
	Value      float64 `json:"value"`
	Type       string  `json:"type"`
	Percentage float64 `json:"percentage"`
}

type OwnershipData []OwnershipEntry

type OwnershipEntry struct {
	ID    string  `json:"id"`
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type SellOffData struct {
	Dates   []string        `json:"dates"`
	Wallets []WalletBalance `json:"wallets"`
}

type WalletBalance struct {
	ID       string    `json:"id"`
	Label    string    `json:"label"`
	Balances []float64 `json:"balances"`
}

type VolumeBracketData []VolumeBracket

type VolumeBracket struct {
	Bracket string `json:"bracket"`
	Count   int    `json:"count"`
}

type BotVolumeData []BotVolumeEntry

type BotVolumeEntry struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type PostRugData struct {
	LpPull      float64         `json:"lpPull"`
	PriceData   PriceData       `json:"priceData"`
	ActivityData ActivityData   `json:"activityData"`
}

type PriceData struct {
	Dates    []string  `json:"dates"`
	Prices   []float64 `json:"prices"`
	RugEvent string    `json:"rugEvent"`
}

type ActivityData struct {
	Dates        []string `json:"dates"`
	Transactions []int    `json:"transactions"`
	RugEvent     string   `json:"rugEvent"`
}

type WalletClusteringData struct {
	Nodes    []Node      `json:"nodes"`
	Links    []Link      `json:"links"`
	Timeline []TimeEvent `json:"timeline"`
}

type TimeEvent struct {
	Time   string  `json:"time"`
	Wallet string  `json:"wallet"`
	Action string  `json:"action"`
	Target string  `json:"target,omitempty"`
	Amount float64 `json:"amount"`
}

type DashboardSummary struct {
	TotalTransactions  int     `json:"totalTransactions"`
	TransactionsChange float64 `json:"transactionsChange"`
	ActiveWallets      int     `json:"activeWallets"`
	WalletsChange      float64 `json:"walletsChange"`
	SuspiciousActivity int     `json:"suspiciousActivity"`
	SuspiciousChange   float64 `json:"suspiciousChange"`
	BotPercentage      float64 `json:"botPercentage"`
	WhaleConcentration float64 `json:"whaleConcentration"`
	AnomalyCount       int     `json:"anomalyCount"`
}

// Dune API response types
type DuneTransaction struct {
	BlockNumber    int64   `json:"block_number"`
	BlockTimestamp string  `json:"block_timestamp"`
	From           string  `json:"from"`
	To             string  `json:"to"`
	Value          float64 `json:"value"`
	TokenAddress   string  `json:"token_address,omitempty"`
	TokenName      string  `json:"token_name,omitempty"`
	TokenSymbol    string  `json:"token_symbol,omitempty"`
	TokenAmount    float64 `json:"token_amount,omitempty"`
}

type DuneBalance struct {
	Token       string  `json:"token"`
	TokenSymbol string  `json:"token_symbol"`
	Amount      float64 `json:"amount"`
}

// Helper function to make API calls to Dune
func callDuneAPI(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", DUNE_API_BASE, endpoint)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("X-Dune-Api-Key", DUNE_API_KEY)
	
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return body, nil
}

// Function to get transactions for a specific address
func getAddressTransactions(address string) ([]DuneTransaction, error) {
	endpoint := fmt.Sprintf("transactions/svm/%s", address)
	
	body, err := callDuneAPI(endpoint)
	if err != nil {
		return nil, err
	}
	
	var transactions []DuneTransaction
	if err := json.Unmarshal(body, &transactions); err != nil {
		return nil, err
	}
	
	return transactions, nil
}

// Function to get balances for a specific address
func getAddressBalances(address string) ([]DuneBalance, error) {
	endpoint := fmt.Sprintf("balances/svm/%s", address)
	
	body, err := callDuneAPI(endpoint)
	if err != nil {
		return nil, err
	}
	
	var balances []DuneBalance
	if err := json.Unmarshal(body, &balances); err != nil {
		return nil, err
	}
	
	return balances, nil
}

// Implementation for transaction flow mapping
func getTransactionFlow(tokenAddress string) (TransactionFlowData, error) {
	// In a real implementation, you would:
	// 1. Query transactions related to the token address
	// 2. Build a graph of wallet interactions
	// 3. Format the data according to your TransactionFlowData structure
	
	// For now, returning sample data similar to your dashboard-api.ts
	nodes := []Node{
		{ID: "wallet1", Label: "Deployer", Size: 25, Color: "#FF6B6B"},
		{ID: "wallet2", Label: "Insider 1", Size: 20, Color: "#FF6B6B"},
		{ID: "wallet3", Label: "Early Buyer", Size: 15, Color: "#82e0aa"},
		{ID: "wallet4", Label: "Whale", Size: 22, Color: "#aed6f1"},
		{ID: "exchange1", Label: "Exchange 1", Size: 30, Color: "#f5cba7"},
	}
	
	links := []Link{
		{Source: "wallet1", Target: "wallet2", Value: 8},
		{Source: "wallet2", Target: "wallet3", Value: 5},
		{Source: "wallet3", Target: "wallet4", Value: 3},
		{Source: "wallet4", Target: "exchange1", Value: 10},
	}
	
	return TransactionFlowData{
		Nodes: nodes,
		Links: links,
	}, nil
}

// Implementation for anomaly detection
func getAnomalyData(tokenAddress string) (AnomalyData, error) {
	// In a real implementation, you would:
	// 1. Query historical price and volume data for the token
	// 2. Apply anomaly detection algorithms to identify spikes
	// 3. Format the data according to your AnomalyData structure
	
	// Generating sample data for now
	dates := []string{
		"2023-04-01", "2023-04-02", "2023-04-03", "2023-04-04",
		"2023-04-05", "2023-04-06", "2023-04-07", "2023-04-08",
		"2023-04-09", "2023-04-10", "2023-04-11", "2023-04-12",
	}
	
	values := []float64{
		120, 125, 130, 220, 190, 185, 250, 280, 275, 190, 350, 320,
	}
	
	anomalies := []Anomaly{
		{Date: "2023-04-04", Value: 220, Type: "spike", Percentage: 69.2},
		{Date: "2023-04-11", Value: 350, Type: "spike", Percentage: 84.2},
	}
	
	return AnomalyData{
		Dates:     dates,
		Values:    values,
		Anomalies: anomalies,
	}, nil
}

// Implementation for ownership concentration data
func getOwnershipData(tokenAddress string) (OwnershipData, error) {
	// In a real implementation, you would:
	// 1. Query token balances for top holders
	// 2. Calculate percentages
	// 3. Format the data according to your OwnershipData structure
	
	// Sample data for now
	data := OwnershipData{
		{ID: "wallet1", Label: "Deployer", Value: 25.3},
		{ID: "wallet2", Label: "Insider 1", Value: 18.7},
		{ID: "wallet3", Label: "Whale 1", Value: 12.4},
		{ID: "wallet4", Label: "Whale 2", Value: 8.9},
		{ID: "wallet5", Label: "Whale 3", Value: 6.2},
		{ID: "wallet6", Label: "Whale 4", Value: 4.8},
		{ID: "wallet7", Label: "Whale 5", Value: 3.5},
		{ID: "wallet8", Label: "Whale 6", Value: 2.9},
		{ID: "wallet9", Label: "Whale 7", Value: 2.1},
		{ID: "wallet10", Label: "Whale 8", Value: 1.8},
		{ID: "others", Label: "Others", Value: 13.4},
	}
	
	return data, nil
}

// Implementation for sell-off pattern data
func getSellOffData(tokenAddress string) (SellOffData, error) {
	// In a real implementation, you would:
	// 1. Query historical balance data for major token holders
	// 2. Track balance changes over time
	// 3. Format the data according to your SellOffData structure
	
	// Sample data for now
	dates := []string{
		"2023-04-01", "2023-04-02", "2023-04-03", "2023-04-04",
		"2023-04-05", "2023-04-06", "2023-04-07",
	}
	
	wallets := []WalletBalance{
		{
			ID:       "wallet1",
			Label:    "Deployer",
			Balances: []float64{1000000, 1000000, 950000, 800000, 500000, 200000, 0},
		},
		{
			ID:       "wallet2",
			Label:    "Insider 1",
			Balances: []float64{800000, 800000, 800000, 750000, 600000, 300000, 100000},
		},
		{
			ID:       "wallet3",
			Label:    "Whale 1",
			Balances: []float64{600000, 600000, 600000, 600000, 550000, 400000, 200000},
		},
		{
			ID:       "wallet4",
			Label:    "Whale 2",
			Balances: []float64{400000, 400000, 400000, 400000, 400000, 350000, 300000},
		},
	}
	
	return SellOffData{
		Dates:   dates,
		Wallets: wallets,
	}, nil
}

// Implementation for volume bracket distribution
func getVolumeBracketData(tokenAddress string) (VolumeBracketData, error) {
	// In a real implementation, you would:
	// 1. Query transaction data for the token
	// 2. Group transactions by volume brackets
	// 3. Count transactions in each bracket
	
	// Sample data for now
	data := VolumeBracketData{
		{Bracket: "$0-$100", Count: 1245},
		{Bracket: "$100-$500", Count: 842},
		{Bracket: "$500-$1K", Count: 433},
		{Bracket: "$1K-$5K", Count: 287},
		{Bracket: "$5K-$10K", Count: 126},
		{Bracket: "$10K-$50K", Count: 64},
		{Bracket: "$50K-$100K", Count: 28},
		{Bracket: "$100K+", Count: 12},
	}
	
	return data, nil
}

// Implementation for bot volume detection
func getBotVolumeData(tokenAddress string) (BotVolumeData, error) {
	// In a real implementation, you would:
	// 1. Apply heuristics to identify bot transactions
	// 2. Calculate percentages
	
	// Sample data for now
	data := BotVolumeData{
		{Type: "Bot Transactions", Value: 42.7},
		{Type: "Organic Transactions", Value: 57.3},
	}
	
	return data, nil
}

// Implementation for post-rug indicators
func getPostRugData(tokenAddress string) (PostRugData, error) {
	// In a real implementation, you would:
	// 1. Analyze price charts for sudden drops
	// 2. Check liquidity pool changes
	// 3. Track activity changes before and after suspected rug pull
	
	// Sample data for now
	dates := []string{
		"2023-04-01", "2023-04-02", "2023-04-03", "2023-04-04",
		"2023-04-05", "2023-04-06", "2023-04-07",
	}
	
	prices := []float64{0.00012, 0.00011, 0.00010, 0.000095, 0.000025, 0.0000032, 0.0000008}
	
	transactions := []int{1245, 1322, 1187, 1402, 1523, 245, 32}
	
	return PostRugData{
		LpPull: 87.5,
		PriceData: PriceData{
			Dates:    dates,
			Prices:   prices,
			RugEvent: "2023-04-05",
		},
		ActivityData: ActivityData{
			Dates:        dates,
			Transactions: transactions,
			RugEvent:     "2023-04-05",
		},
	}, nil
}

// Implementation for wallet clustering data
func getWalletClusteringData(tokenAddress string) (WalletClusteringData, error) {
	// In a real implementation, you would:
	// 1. Apply clustering algorithms to group similar wallets
	// 2. Identify transaction patterns between clusters
	// 3. Label clusters based on behavior patterns
	
	// Sample data for now
	nodes := []Node{
		{ID: "cluster1", Label: "Deployer Group", Size: 25, Color: "#FF6B6B", Type: "cluster"},
		{ID: "cluster2", Label: "Market Makers", Size: 20, Color: "#f5cba7", Type: "cluster"},
		{ID: "cluster3", Label: "Regular Traders", Size: 15, Color: "#aed6f1", Type: "cluster"},
		{ID: "wallet1", Label: "Deployer", Size: 10, Color: "#FF6B6B", Type: "wallet", Cluster: "cluster1"},
		{ID: "wallet2", Label: "Insider 1", Size: 10, Color: "#FF6B6B", Type: "wallet", Cluster: "cluster1"},
		{ID: "wallet3", Label: "Insider 2", Size: 10, Color: "#FF6B6B", Type: "wallet", Cluster: "cluster1"},
		{ID: "wallet4", Label: "MM Bot 1", Size: 10, Color: "#f5cba7", Type: "wallet", Cluster: "cluster2"},
		{ID: "wallet5", Label: "MM Bot 2", Size: 10, Color: "#f5cba7", Type: "wallet", Cluster: "cluster2"},
		{ID: "wallet6", Label: "Trader 1", Size: 10, Color: "#aed6f1", Type: "wallet", Cluster: "cluster3"},
		{ID: "wallet7", Label: "Trader 2", Size: 10, Color: "#aed6f1", Type: "wallet", Cluster: "cluster3"},
	}
	
	links := []Link{
		{Source: "wallet1", Target: "cluster1", Value: 1},
		{Source: "wallet2", Target: "cluster1", Value: 1},
		{Source: "wallet3", Target: "cluster1", Value: 1},
		{Source: "wallet4", Target: "cluster2", Value: 1},
		{Source: "wallet5", Target: "cluster2", Value: 1},
		{Source: "wallet6", Target: "cluster3", Value: 1},
		{Source: "wallet7", Target: "cluster3", Value: 1},
		{Source: "wallet1", Target: "wallet2", Value: 3},
		{Source: "wallet2", Target: "wallet3", Value: 2},
		{Source: "wallet4", Target: "wallet5", Value: 4},
		{Source: "wallet6", Target: "wallet7", Value: 1},
		{Source: "cluster1", Target: "cluster2", Value: 5},
		{Source: "cluster2", Target: "cluster3", Value: 3},
	}
	
	timeline := []TimeEvent{
		{Time: "2023-04-01 08:23", Wallet: "wallet1", Action: "buy", Amount: 50000},
		{Time: "2023-04-01 09:45", Wallet: "wallet2", Action: "buy", Amount: 75000},
		{Time: "2023-04-01 12:12", Wallet: "wallet4", Action: "buy", Amount: 120000},
		{Time: "2023-04-02 14:30", Wallet: "wallet1", Action: "transfer", Target: "wallet3", Amount: 25000},
		{Time: "2023-04-02 15:22", Wallet: "wallet4", Action: "transfer", Target: "wallet5", Amount: 60000},
		{Time: "2023-04-03 10:15", Wallet: "wallet6", Action: "buy", Amount: 90000},
		{Time: "2023-04-03 11:45", Wallet: "wallet7", Action: "buy", Amount: 45000},
		{Time: "2023-04-04 09:30", Wallet: "wallet3", Action: "sell", Amount: 15000},
		{Time: "2023-04-04 16:20", Wallet: "wallet5", Action: "sell", Amount: 30000},
	}
	
	return WalletClusteringData{
		Nodes:    nodes,
		Links:    links,
		Timeline: timeline,
	}, nil
}

// Implementation for dashboard summary data
func getDashboardSummary(tokenAddress string) (DashboardSummary, error) {
	// In a real implementation, you would:
	// 1. Aggregate data from various analyses
	// 2. Calculate summary metrics
	
	// Sample data for now
	summary := DashboardSummary{
		TotalTransactions:  1234,
		TransactionsChange: 12.5,
		ActiveWallets:      567,
		WalletsChange:      8.3,
		SuspiciousActivity: 89,
		SuspiciousChange:   -5.2,
		BotPercentage:      42.7,
		WhaleConcentration: 84.7,
		AnomalyCount:       5,
	}
	
	return summary, nil
}

// API handlers

func handleTransactionFlow(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getTransactionFlow(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleAnomalyData(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getAnomalyData(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleOwnershipData(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getOwnershipData(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleSellOffData(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getSellOffData(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleVolumeBracketData(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getVolumeBracketData(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleBotVolumeData(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getBotVolumeData(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handlePostRugData(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getPostRugData(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleWalletClusteringData(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getWalletClusteringData(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleDashboardSummary(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("token")
	if tokenAddress == "" {
		http.Error(w, "Missing token address", http.StatusBadRequest)
		return
	}
	
	data, err := getDashboardSummary(tokenAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Cross-Origin Resource Sharing (CORS) middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Define API routes
	http.Handle("/api/transaction-flow", corsMiddleware(http.HandlerFunc(handleTransactionFlow)))
	http.Handle("/api/anomaly-data", corsMiddleware(http.HandlerFunc(handleAnomalyData)))
	http.Handle("/api/ownership-data", corsMiddleware(http.HandlerFunc(handleOwnershipData)))
	http.Handle("/api/sell-off-data", corsMiddleware(http.HandlerFunc(handleSellOffData)))
	http.Handle("/api/volume-bracket-data", corsMiddleware(http.HandlerFunc(handleVolumeBracketData)))
	http.Handle("/api/bot-volume-data", corsMiddleware(http.HandlerFunc(handleBotVolumeData)))
	http.Handle("/api/post-rug-data", corsMiddleware(http.HandlerFunc(handlePostRugData)))
	http.Handle("/api/wallet-clustering-data", corsMiddleware(http.HandlerFunc(handleWalletClusteringData)))
	http.Handle("/api/dashboard-summary", corsMiddleware(http.HandlerFunc(handleDashboardSummary)))
	
	// Start the server
	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}