# Mobula Go SDK

A comprehensive Go SDK for the [Mobula API](https://docs.mobula.io), providing access to cryptocurrency market data, wallet analytics, and blockchain information.

## Features

- **Market Data**: Access real-time and historical price data, trading volumes, market caps, and more
- **Wallet Analytics**: Query wallet holdings, NFTs, transaction history, and net worth tracking
- **Token Information**: Get detailed token metadata, security analysis, and holder information
- **Trading Data**: Retrieve trade history, OHLCV data, and market pairs
- **Multi-chain Support**: Works across multiple blockchain networks
- **Easy to Use**: Simple, idiomatic Go API with comprehensive type definitions
- **Demo Mode**: Test without API key using the demo endpoint

## Installation

```bash
go get github.com/mobula-go-sdk
```

## Quick Start

### Using Demo Mode (No API Key Required)

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/mobula-go-sdk"
)

func main() {
    // Create a demo client
    client := mobula.NewDemoClient()

    ctx := context.Background()

    // Get token details
    details, err := client.Market.GetTokenDetails(ctx, &mobula.GetTokenDetailsRequest{
        Asset: "Bitcoin",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s (%s): $%.2f\n", details.Name, details.Symbol, details.Price)
}
```

### Using Production Mode (With API Key)

```go
client := mobula.NewClient(&mobula.Config{
    APIKey: "your-api-key-here",
})
```

## API Documentation

### Market Data Service

#### Get Token Details

```go
details, err := client.Market.GetTokenDetails(ctx, &mobula.GetTokenDetailsRequest{
    Asset: "Ethereum",
    Blockchain: "Ethereum", // Optional
})
```

#### Get Market Data

```go
data, err := client.Market.GetMarketData(ctx, &mobula.GetMarketDataRequest{
    Asset: "Bitcoin",
})
// Returns: price, volume, market cap, liquidity, etc.
```

#### Get Multiple Token Prices

```go
prices, err := client.Market.GetMarketPrices(ctx, &mobula.GetMarketPricesRequest{
    Assets: []string{"Bitcoin", "Ethereum", "BNB"},
})
// Returns: map[string]float64 of asset -> price
```

#### Get Historical Data

```go
history, err := client.Market.GetHistoricalData(ctx, &mobula.GetHistoricalDataRequest{
    Asset: "Ethereum",
    From: 1609459200, // Unix timestamp
    To: 1640995200,
})
```

#### Get OHLCV Data

```go
ohlcv, err := client.Market.GetOHLCVData(ctx, &mobula.GetOHLCVDataRequest{
    PairAddress: "0x...",
    Blockchain: "Ethereum",
    Interval: "1h", // 1m, 5m, 15m, 1h, 4h, 1d, etc.
})
```

#### Get Token Markets/Pairs

```go
markets, err := client.Market.GetTokenMarkets(ctx, &mobula.GetTokenMarketsRequest{
    Asset: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
    Blockchain: "Ethereum",
})
```

#### Get Token Trades

```go
trades, err := client.Market.GetTokenTrades(ctx, &mobula.GetTokenTradesRequest{
    Asset: "Bitcoin",
    Limit: 100,
    Offset: 0,
})
```

#### Get Token Security Info

```go
security, err := client.Market.GetTokenSecurity(ctx, &mobula.GetTokenSecurityRequest{
    Asset: "0x...",
    Blockchain: "Ethereum",
})
// Returns: honeypot status, taxes, holder concentration, etc.
```

#### Get Token Holders

```go
holders, err := client.Market.GetTokenHolders(ctx, &mobula.GetTokenHoldersRequest{
    Asset: "0x...",
    Blockchain: "Ethereum",
    Limit: 100,
})
```

#### Get Trending Tokens (Pulse)

```go
pulse, err := client.Market.GetPulse(ctx, &mobula.GetPulseRequest{
    Category: "trending", // trending, gainers, losers, recent
})
```

### Wallet Service

#### Get Wallet Holdings

```go
holdings, err := client.Wallet.GetWalletHoldings(ctx, &mobula.GetWalletHoldingsRequest{
    Wallet: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    Blockchains: []string{"Ethereum", "BSC"},
})
```

#### Get Wallet NFTs

```go
nfts, err := client.Wallet.GetWalletNFTs(ctx, &mobula.GetWalletNFTsRequest{
    Wallet: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    Blockchain: "Ethereum",
    Limit: 50,
})
```

#### Get Historical Net Worth

```go
netWorth, err := client.Wallet.GetWalletHistoricalNetWorth(ctx, &mobula.GetWalletHistoricalNetWorthRequest{
    Wallet: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    Blockchains: []string{"Ethereum"},
})
```

#### Get Wallet Labels

```go
labels, err := client.Wallet.GetWalletLabels(ctx, &mobula.GetWalletLabelsRequest{
    Wallet: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
})
```

#### Get Wallet Transactions

```go
txs, err := client.Wallet.GetWalletTransactions(ctx, &mobula.GetWalletTransactionsRequest{
    Wallet: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    Blockchain: "Ethereum",
    Limit: 100,
})
```

#### Get Specific Token Balance

```go
balance, err := client.Wallet.GetWalletBalance(ctx, &mobula.GetWalletBalanceRequest{
    Wallet: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    Asset: "USDT",
    Blockchain: "Ethereum",
})
```

## Configuration

### Custom HTTP Client

```go
import "net/http"
import "time"

client := mobula.NewClient(&mobula.Config{
    APIKey: "your-api-key",
    HTTPClient: &http.Client{
        Timeout: 60 * time.Second,
    },
})
```

### Custom Base URL

```go
client := mobula.NewClient(&mobula.Config{
    BaseURL: "https://custom-endpoint.mobula.io",
    APIKey: "your-api-key",
})
```

### Custom Timeout

```go
client := mobula.NewClient(&mobula.Config{
    APIKey: "your-api-key",
    Timeout: 60 * time.Second,
})
```

## Error Handling

The SDK provides structured error responses:

```go
data, err := client.Market.GetTokenDetails(ctx, &mobula.GetTokenDetailsRequest{
    Asset: "InvalidToken",
})
if err != nil {
    if apiErr, ok := err.(*mobula.APIError); ok {
        fmt.Printf("API Error (Status %d): %s\n", apiErr.StatusCode, apiErr.Message)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Examples

Complete examples are available in the [examples](./examples) directory:

- [Market Data Examples](./examples/market_data.go) - Token details, prices, trades, security
- [Wallet Data Examples](./examples/wallet_data.go) - Holdings, NFTs, transactions, net worth

Run examples:

```bash
cd examples
go run market_data.go
go run wallet_data.go
```

## Supported Blockchains

The Mobula API supports multiple blockchains including:

- Ethereum
- BNB Chain (BSC)
- Polygon
- Avalanche
- Arbitrum
- Optimism
- Fantom
- And many more...

## API Endpoints

The SDK implements all major Mobula API endpoints:

### Market Data
- `/api/1/metadata` - Token metadata and details
- `/api/1/market/data` - Current market data
- `/api/1/market/multi-data` - Batch price queries
- `/api/1/market/pairs` - Trading pairs
- `/api/1/market/trades` - Trade history
- `/api/1/market/history` - Historical price data
- `/api/1/market/multi-history` - Batch historical data
- `/api/1/market/history/pair` - OHLCV data
- `/api/1/market/security` - Token security analysis
- `/api/1/market/holders` - Token holder information
- `/api/1/market/pulse` - Trending tokens

### Wallet Data
- `/api/1/wallet/portfolio` - Wallet holdings
- `/api/1/wallet/nfts` - NFT holdings
- `/api/1/wallet/history` - Historical net worth
- `/api/1/wallet/labels` - Wallet labels
- `/api/1/wallet/transactions` - Transaction history
- `/api/1/wallet/balance` - Token balance

## Rate Limiting

The Mobula API has rate limits depending on your plan. The demo endpoint has lower limits. Consider:

- Using the production API with an API key for higher limits
- Implementing retry logic with exponential backoff
- Caching responses when appropriate

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Resources

- [Mobula API Documentation](https://docs.mobula.io/)
- [Mobula Website](https://mobula.io/)
- [Get API Key](https://mobula.io/dashboard)

## Support

For API-related questions, visit the [Mobula Documentation](https://docs.mobula.io/).

For SDK issues, please open an issue on GitHub.
