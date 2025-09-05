package core

import "time"

// GetHoldersOptions holds optional parameters for GetHolders.
type GetHoldersOptions struct {
	// TimeoutSeconds overrides per-request timeout if provided.
	TimeoutSeconds *int32
}

// HolderItem corresponds to one token entry in the holders response.
type HolderItem struct {
	Token   string   `json:"token"`
	Holders []Holder `json:"holders"`
}

type Holder struct {
	ProxyWallet           string     `json:"proxyWallet"`
	Bio                   string     `json:"bio"`
	Asset                 string     `json:"asset"`
	Pseudonym             string     `json:"pseudonym"`
	Amount                float64    `json:"amount"`
	DisplayUsernamePublic bool       `json:"displayUsernamePublic"`
	OutcomeIndex          int        `json:"outcomeIndex"`
	Name                  string     `json:"name"`
	ProfileImage          string     `json:"profileImage"`
	ProfileImageOptimized string     `json:"profileImageOptimized"`
	CreatedAt             *time.Time `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time `json:"updatedAt,omitempty"`
}

type GetHoldersQuery struct {
	Limit        int
	Market       []string
	ConditionIDs []string
}
