package entitlementsvc

import (
	"context"
	"errors"
)

// Service is a simple CRUD interface for entitlement.
type Service interface {
	//PostEntitlement(ctx context.Context, e EntitlementData) (string, error)
	GetEntitlement(ctx context.Context, userId string) (Entitlement, error)
}

type Entitlement struct {
	ID       string    `json:"id"`
	Features []Feature `json:"features,omitempty"`
}
type Feature struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	AccessLevel string `json:"level,omitempty"`
}

type entitlementService struct {
}

func NewEntitlementService() Service {
	return &entitlementService{}
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

func (s *entitlementService) GetEntitlement(ctx context.Context, userId string) (Entitlement, error) {
	feature1 := Feature{ID: "F01", Name: "View Profile", AccessLevel: "ALL"}
	feature2 := Feature{ID: "F02", Name: "Update Profile", AccessLevel: "ALL"}

	features := []Feature{feature1, feature2}

	entitlement := Entitlement{"EN001", features}

	//entitlementData := EntitlementData{"001"}
	return entitlement, nil
}
