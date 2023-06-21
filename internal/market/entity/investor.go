package entity

type Investor struct {
	ID            string                   `json:"id"`
	Name          string                   `json:"name"`
	AssetPosition []*InvestorAssetPosition `json:"assets"`
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdshares int64) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtdshares))
	} else{
		assetPosition.Shares += qtdshares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}

	}
	return nil
}

type InvestorAssetPosition struct {
	AssetID string `json:"asset_id"`
	Shares  int64  `json:"shares"`
}

func NewInvestorAssetPosition(assetID string, shares int64) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}
