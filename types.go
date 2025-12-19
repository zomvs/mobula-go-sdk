package mobula

import "time"

// Common response wrapper
type Response struct {
	Data interface{} `json:"data"`
}

// TokenDetails represents detailed information about a token
type TokenDetails struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Symbol            string                 `json:"symbol"`
	Contracts         []string               `json:"contracts"`
	Blockchains       []string               `json:"blockchains"`
	Price             float64                `json:"price"`
	PriceChange24h    float64                `json:"price_change_24h"`
	MarketCap         float64                `json:"market_cap"`
	MarketCapDiluted  float64                `json:"market_cap_diluted"`
	Liquidity         float64                `json:"liquidity"`
	Volume            float64                `json:"volume"`
	Volume7d          float64                `json:"volume_7d"`
	IsListed          bool                   `json:"is_listed"`
	Logo              string                 `json:"logo"`
	Decimals          int                    `json:"decimals"`
	TotalSupply       float64                `json:"total_supply"`
	CirculatingSupply float64                `json:"circulating_supply"`
	Description       string                 `json:"description"`
	Website           string                 `json:"website"`
	Twitter           string                 `json:"twitter"`
	Discord           string                 `json:"discord"`
	Telegram          string                 `json:"telegram"`
	Github            string                 `json:"github"`
	Coingecko         string                 `json:"coingecko"`
	Coinmarketcap     string                 `json:"coinmarketcap"`
	Metadata          map[string]interface{} `json:"metadata"`
}

// MarketData represents market data for a token
type MarketData struct {
	Price          float64 `json:"price"`
	PriceChange24h float64 `json:"price_change_24h"`
	Volume24h      float64 `json:"volume_24h"`
	MarketCap      float64 `json:"market_cap"`
	Liquidity      float64 `json:"liquidity"`
}

// HistoricalData represents historical price data
type HistoricalData struct {
	Timestamp int64   `json:"timestamp"`
	Price     float64 `json:"price"`
	Volume    float64 `json:"volume"`
	MarketCap float64 `json:"market_cap"`
}

// OHLCVData represents OHLCV (Open, High, Low, Close, Volume) data
type OHLCVData struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
}

// Market represents a trading market/pair
type Market struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	BaseToken   string  `json:"base_token"`
	QuoteToken  string  `json:"quote_token"`
	Price       float64 `json:"price"`
	Volume24h   float64 `json:"volume_24h"`
	Liquidity   float64 `json:"liquidity"`
	Exchange    string  `json:"exchange"`
	Blockchain  string  `json:"blockchain"`
	PairAddress string  `json:"pair_address"`
}

// Trade represents a trade transaction
type Trade struct {
	Timestamp    int64   `json:"timestamp"`
	Type         string  `json:"type"` // buy or sell
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Total        float64 `json:"total"`
	TxHash       string  `json:"tx_hash"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Blockchain   string  `json:"blockchain"`
	Exchange     string  `json:"exchange"`
	PairAddress  string  `json:"pair_address"`
	TokenAddress string  `json:"token_address"`
	AmountUSD    float64 `json:"amount_usd"`
}

// TokenSecurity represents security information about a token
type TokenSecurity struct {
	IsOpenSource         bool    `json:"is_open_source"`
	IsProxy              bool    `json:"is_proxy"`
	IsMintable           bool    `json:"is_mintable"`
	CanTakeBackOwnership bool    `json:"can_take_back_ownership"`
	OwnerChangeBalance   bool    `json:"owner_change_balance"`
	HiddenOwner          bool    `json:"hidden_owner"`
	SelfDestruct         bool    `json:"selfdestruct"`
	ExternalCall         bool    `json:"external_call"`
	BuyTax               float64 `json:"buy_tax"`
	SellTax              float64 `json:"sell_tax"`
	IsHoneypot           bool    `json:"is_honeypot"`
	HolderCount          int     `json:"holder_count"`
	TotalSupply          float64 `json:"total_supply"`
	HolderConcentration  float64 `json:"holder_concentration"`
}

// TokenHolder represents information about a token holder
type TokenHolder struct {
	Address    string   `json:"address"`
	Balance    float64  `json:"balance"`
	BalanceUSD float64  `json:"balance_usd"`
	Percentage float64  `json:"percentage"`
	IsContract bool     `json:"is_contract"`
	Labels     []string `json:"labels"`
}

// TraderPosition represents a trader's position
type TraderPosition struct {
	Address         string    `json:"address"`
	TokenAddress    string    `json:"token_address"`
	Blockchain      string    `json:"blockchain"`
	Balance         float64   `json:"balance"`
	BalanceUSD      float64   `json:"balance_usd"`
	AverageBuyPrice float64   `json:"average_buy_price"`
	TotalBought     float64   `json:"total_bought"`
	TotalSold       float64   `json:"total_sold"`
	ProfitLoss      float64   `json:"profit_loss"`
	ProfitLossUSD   float64   `json:"profit_loss_usd"`
	FirstBuy        time.Time `json:"first_buy"`
	LastActivity    time.Time `json:"last_activity"`
}

// WalletNFT represents an NFT in a wallet
type WalletNFT struct {
	TokenID     string                 `json:"token_id"`
	Contract    string                 `json:"contract"`
	Blockchain  string                 `json:"blockchain"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Image       string                 `json:"image"`
	Collection  string                 `json:"collection"`
	FloorPrice  float64                `json:"floor_price"`
	LastSale    float64                `json:"last_sale"`
	Metadata    map[string]interface{} `json:"metadata"`
	Attributes  []NFTAttribute         `json:"attributes"`
}

// NFTAttribute represents an NFT attribute/trait
type NFTAttribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
	Rarity    float64     `json:"rarity,omitempty"`
}

// WalletHolding represents a crypto holding in a wallet
type WalletHolding struct {
	Asset      string  `json:"asset"`
	Symbol     string  `json:"symbol"`
	Address    string  `json:"address"`
	Blockchain string  `json:"blockchain"`
	Balance    float64 `json:"balance"`
	BalanceUSD float64 `json:"balance_usd"`
	Price      float64 `json:"price"`
	Logo       string  `json:"logo"`
	Decimals   int     `json:"decimals"`
}

// WalletNetWorth represents historical net worth data
type WalletNetWorth struct {
	Timestamp int64   `json:"timestamp"`
	NetWorth  float64 `json:"net_worth"`
	Assets    []struct {
		Symbol     string  `json:"symbol"`
		Value      float64 `json:"value"`
		Blockchain string  `json:"blockchain"`
	} `json:"assets"`
}

// WalletLabel represents a label associated with a wallet
type WalletLabel struct {
	Address    string   `json:"address"`
	Labels     []string `json:"labels"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	IsContract bool     `json:"is_contract"`
}

// PulseData represents pulse/trending data
type PulseData struct {
	Trending []TrendingToken `json:"trending"`
	Gainers  []TrendingToken `json:"gainers"`
	Losers   []TrendingToken `json:"losers"`
	Recent   []TrendingToken `json:"recent"`
}

// TrendingToken represents a trending token
type TrendingToken struct {
	Name           string  `json:"name"`
	Symbol         string  `json:"symbol"`
	Address        string  `json:"address"`
	Blockchain     string  `json:"blockchain"`
	Price          float64 `json:"price"`
	PriceChange24h float64 `json:"price_change_24h"`
	Volume24h      float64 `json:"volume_24h"`
	MarketCap      float64 `json:"market_cap"`
	Liquidity      float64 `json:"liquidity"`
	Logo           string  `json:"logo"`
	TrendingScore  float64 `json:"trending_score,omitempty"`
}
