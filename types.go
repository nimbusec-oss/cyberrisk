package cyberrisk

import (
	"time"
)

type OrderRating struct {
	RequestedCategory string `json:"requestedCategory"` // mandatory
}

type RequestSupplier struct {
	VAT         string   `json:"vat"`         // mandatory
	CompanyName string   `json:"companyName"` // mandatory
	Language    string   `json:"language"`    // mandatory
	Email       string   `json:"email"`       // mandatory
	Websites    []string `json:"websites"`    // mandatory

	Forename string `json:"forename,omitempty"` // optional
	Surname  string `json:"surname,omitempty"`  // optional
	Tel      string `json:"tel,omitempty"`      // optional

	Street     string `json:"street,omitempty"`     // optional
	ZipCode    string `json:"zipCode,omitempty"`    // optional
	City       string `json:"city,omitempty"`       // optional
	Country    string `json:"country,omitempty"`    // optional
	Sector     string `json:"sector,omitempty"`     // optional
	ExternalID string `json:"externalID,omitempty"` // optional
	ExternalIDs []string `json:"externalIDs,omitempty"` // optional

	OrderRating *OrderRating `json:"orderRating,omitempty"` // optional
}

type Supplier struct {
	VAT         string   `json:"vat"`
	CompanyName string   `json:"companyName"`
	Language    string   `json:"language"`
	Email       string   `json:"email"`
	Websites    []string `json:"websites"`

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

type SupplierFilter struct {
	IncludeRatings bool `url:"include-ratings"`
}

type UnassignSupplier struct {
	SupplierIDs []string `url:"supplier-id"`
	ExternalIDs []string `url:"external-id"`
}

type Rating struct {
	CompanyName     string     `json:"companyName,omitempty"`
	SupplierID      string     `json:"supplierID"`
	ExternalID      string     `json:"externalID,omitempty"` // Deprecated
	ExternalIDs     []string   `json:"externalIDs,omitempty"`
	AScore          int        `json:"aScore,omitempty"`
	BScore          int        `json:"bScore,omitempty"`
	Status          string     `json:"status,omitempty"`
	ValidFrom       *time.Time `json:"validFrom,omitempty"`
	ValidUntil      *time.Time `json:"validUntil,omitempty"`
	CyberTrustLabel string     `json:"cyberTrustLabel,omitempty"`
	CScore          int        `json:"cScore,omitempty"`
	StatusCScore    string     `json:"statusCScore,omitempty"`
	Websites        []string   `json:"websites,omitempty"`
}

type RatingFilter struct {
	Status      []string `url:"status"`
	SupplierIDs []string `url:"supplier-id"`
	ExternalIDs []string `url:"external-id"`
}

type RequestRating struct {
	Category   string `json:"category"`
	SupplierID string `json:"supplierID,omitempty"`
	ExternalID string `json:"externalID,omitempty"`
}

type ContingentUsage struct {
	ContingentBefore int      `json:"contingentBefore"`
	ContingentAfter  int      `json:"contingentAfter"`
	OrderedSuppliers []string `json:"orderedSuppliers"`
}

type PostSuppliersReturn struct {
	Created int              `json:"created"`
	Existed int              `json:"existed"`
	Orders  *ContingentUsage `json:"orders,omitempty"`
}

type Account struct {
	Contingent int `json:"contingent,omitempty"`
}
