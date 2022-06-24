package cyberrisk

import (
	"time"
)

type OrderRating struct {
	requestedCategory string `json:"requestedCategory"` // mandatory
}

type RequestSupplier struct {
	VAT         string   `json:"vat"`         // mandatory
	CompanyName string   `json:"companyName"` // mandatory
	Language    string   `json:"language"`    // mandatory
	Email       string   `json:"email"`       // mandatory
	Scope       []string `json:"websites"`    // mandatory

	Forename string `json:"forename,omitempty"` // optional
	Surname  string `json:"surname,omitempty"`  // optional
	Tel      string `json:"tel,omitempty"`      // optional

	Street     string `json:"street,omitempty"`     // optional
	ZipCode    string `json:"zipCode,omitempty"`    // optional
	City       string `json:"city,omitempty"`       // optional
	Country    string `json:"country,omitempty"`    // optional
	Sector     string `json:"sector,omitempty"`     // optional
	ExternalID string `json:"externalID,omitempty"` // optional

	OrderRating *OrderRating `json:"orderRating,omitempty"` // optional
}

type Supplier struct {
	VAT         string   `json:"vat"`
	CompanyName string   `json:"companyName"`
	Language    string   `json:"language"`
	Email       string   `json:"email"`
	Scope       []string `json:"websites"`

	Forename string `json:"forename,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Tel      string `json:"tel,omitempty"`

	Street  string `json:"street,omitempty"`
	ZipCode string `json:"zipCode,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Sector  string `json:"sector,omitempty"`

	ExternalID string `json:"externalID,omitempty"`
	ID         string `json:"id"`

	Rating *Rating `json:"rating,omitempty"`
}

type Rating struct {
	SupplierID      string     `json:"supplierID"`
	ExternalID      string     `json:"externalID,omitempty"`
	AScore          int        `json:"aScore,omitempty"`
	BScore          int        `json:"bScore,omitempty"`
	CScore          int        `json:"cScore,omitempty"`
	Status          string     `json:"status,omitempty"`
	ValidFrom       *time.Time `json:"validFrom,omitempty"`
	ValidUntil      *time.Time `json:"validUntil,omitempty"`
	CyberTrustLabel string     `json:"cyberTrustLabel,omitempty"`
}

type RequestRating struct {
	Category   string `json:"category"`
	SupplierID string `json:"supplierID,omitempty"`
	ExternalID string `json:"externalID,omitempty"`
}

type Account struct {
	Contingent int `json:"contingent,omitempty"`
}
