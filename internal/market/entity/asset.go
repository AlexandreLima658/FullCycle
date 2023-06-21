package entity

type Asset struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	MarketVolume int64  `json:"market_volume"`
}

func NewAsset(id string, name string, marketVolume int64) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
