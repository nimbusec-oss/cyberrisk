package cyberrisk

import (
	"time"
)

type RequestSupplier struct {
	Name       string   `json:"name"`                 // mandatory
	Email      string   `json:"email"`                // mandatory
	VAT        string   `json:"vat"`                  // mandatory
	Street     string   `json:"street,omitempty"`     // optional
	ZipCode    string   `json:"zipCode,omitempty"`    // optional
	City       string   `json:"city,omitempty"`       // optional
	Country    string   `json:"country,omitempty"`    // optional
	Sector     string   `json:"sector,omitempty"`     // optional
	Tel        string   `json:"companyTel,omitempty"` // optional
	Scope      []string `json:"websites"`             // mandatory
	ExternalID string   `json:"externalID,omitempty"` // optional
	Contact    *Contact `json:"contact,omitempty"`    // optional
}

type Supplier struct {
	RequestSupplier
	ID     string  `json:"id"`
	Rating *Rating `json:"rating,omitempty"`
}

type Rating struct {
	SupplierID string     `json:"supplierID,omitempty"`
	ExternalID string     `json:"externalID,omitempty"`
	AScore     int        `json:"aScore,omitempty"`
	BScore     int        `json:"bScore,omitempty"`
	CScore     int        `json:"cScore,omitempty"`
	Status     string     `json:"status,omitempty"`
	ValidFrom  *time.Time `json:"validFrom,omitempty"`
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	Label      string     `json:"label,omitempty"`
}

type Contact struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
	Tel      string `json:"tel"`
	Language string `json:"language"`
}

type RequestRating struct {
	Category    string   `json:"category,omitempty"`
	SupplierIDs []string `json:"supplierIDs,omitempty"`
	ExternalIDs []string `json:"externalIDs,omitempty"`
}

type Account struct {
	Contingent int `json:"contingent,omitempty"`
}
