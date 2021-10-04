package model

import (
	"strconv"
	"time"
)

type Authentication struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`

	Name                    string                 `json:"name,omitempty"`
	AuthType                string                 `json:"authtype"`
	Username                string                 `json:"username"`
	Password                string                 `json:"password"`
	Extra                   map[string]interface{} `json:"extra,omitempty"`
	Version                 string                 `json:"version"`
	AvailabilityStatus      string                 `json:"availability_status,omitempty"`
	AvailabilityStatusError string                 `json:"availability_status_error,omitempty"`

	SourceID int64 `json:"source_id"`
	Source   Source
	TenantID int64 `json:"tenant_id"`
	Tenant   Tenant

	ResourceType string `json:"resource_type"`
	ResourceID   int64  `json:"resource_id"`

	ApplicationAuthentications []ApplicationAuthentication
}

func (auth *Authentication) ToResponse() *AuthenticationResponse {
	resourceID := strconv.FormatInt(auth.ResourceID, 10)
	return &AuthenticationResponse{
		ID:        auth.ID,
		CreatedAt: auth.CreatedAt,
		Name:      auth.Name,
		Version:   auth.Version,
		AuthType:  auth.AuthType,
		Username:  auth.Username,
		// TODO: remove this?
		Password:                auth.Password,
		Extra:                   auth.Extra,
		AvailabilityStatus:      auth.AvailabilityStatus,
		AvailabilityStatusError: auth.AvailabilityStatusError,
		ResourceType:            auth.ResourceType,
		ResourceID:              resourceID,
	}
}

/*
	This method translates an Authentication struct to a hash that will be
	accepted by vault, this format will also be deserialized properly by
	dao.authFromVault, so if we are to add more fields they will need to be
	added here as well.
*/
func (auth *Authentication) ToVaultMap() (map[string]interface{}, error) {
	data := map[string]interface{}{
		"name":                      auth.Name,
		"authtype":                  auth.AuthType,
		"username":                  auth.Username,
		"password":                  auth.Password,
		"extra":                     auth.Extra,
		"availability_status":       auth.AvailabilityStatus,
		"availability_status_error": auth.AvailabilityStatusError,
		"resource_type":             auth.ResourceType,
		"resource_id":               strconv.FormatInt(auth.ResourceID, 10),
	}

	// Vault requires the hash to be wrapped in a "data" object in order to be accepted.
	return map[string]interface{}{"data": data}, nil
}
