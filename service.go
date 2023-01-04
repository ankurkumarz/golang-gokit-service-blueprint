package entitlementsvc

import (
	"context"
)

// Service is a simple CRUD interface for entitlement.
type Service interface {
	PostEntitlement(ctx context.Context, e EntitlementData) (string, error)
	GetEntitlement(ctx context.Context, userId string) (EntitlementData, error)
}

type EntitlementData struct {
	UserID      string      `json:"user_id"`
	Entitlement Entitlement `json:"entitlement,omitempty"`
}

type Entitlement struct {
	ID       string    `json:"id"`
	Features []Feature `json:"feature,omitempty"`
}
type Feature struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	AccessLevel string `json:"access_level,omitempty"`
}
