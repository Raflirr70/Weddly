package entity

type Gift struct {
	Order           int    `json:"order"`
	Provider        string `json:"provider"`
	ProviderAccount string `json:"providerAccount"`
	Type            string `json:"type"`
	No              string `json:"no"`
}