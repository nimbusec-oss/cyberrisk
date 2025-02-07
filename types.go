package cyberrisk

import (
	"time"
)

type ProjectType string

type RequestSupplier struct {
	VAT         string `json:"vat"`         // mandatory
	CompanyName string `json:"companyName"` // mandatory
	Language    string `json:"language"`    // mandatory
	Email       string `json:"email"`       // mandatory

	Websites []string `json:"websites,omitempty"` // optional
	Forename string   `json:"forename,omitempty"` // optional
	Surname  string   `json:"surname,omitempty"`  // optional
	Tel      string   `json:"tel,omitempty"`      // optional

	Street      string   `json:"street,omitempty"`      // optional
	ZipCode     string   `json:"zipCode,omitempty"`     // optional
	City        string   `json:"city,omitempty"`        // optional
	Country     string   `json:"country,omitempty"`     // optional
	Sector      string   `json:"sector,omitempty"`      // optional
	ExternalIDs []string `json:"externalIDs,omitempty"` // optional

	OrderCRR string `json:"orderCRR,omitempty"` // optional
	OrderDPR bool   `json:"orderDPR,omitempty"` // optional
}

type SupplierFilter struct {
	IncludeRatings bool `url:"include-ratings"`
}

type UnassignSupplier struct {
	SupplierIDs []string `url:"supplier-id"`
	ExternalIDs []string `url:"external-id"`
}

type Rating struct {
	CompanyName string   `json:"companyName,omitempty"`
	SupplierID  string   `json:"supplierID"`
	ExternalIDs []string `json:"externalIDs,omitempty"`

	RatingCRR *RatingCRR `json:"ratingCRR,omitempty"`
	RatingDPR *RatingDPR `json:"ratingDPR,omitempty"`
	WebRisk   *WebRisk   `json:"webrisk,omitempty"`
}

type RatingCRR struct {
	AScore               int        `json:"aScore,omitempty"`
	BScore               int        `json:"bScore,omitempty"`
	Status               string     `json:"status,omitempty"`
	ValidFrom            *time.Time `json:"validFrom,omitempty"`
	ValidUntil           *time.Time `json:"validUntil,omitempty"`
	AssessmentValidUntil *time.Time `json:"assessmentValidUntil,omitempty"`
	CyberTrustLabel      string     `json:"cyberTrustLabel,omitempty"`
}

type RatingDPR struct {
	MinimumRequirementsMet bool       `json:"minimumRequirementsMet,omitempty"`
	Score                  int        `json:"score,omitempty"`
	Status                 string     `json:"status,omitempty"`
	ValidFrom              *time.Time `json:"validFrom,omitempty"`
	ValidUntil             *time.Time `json:"validUntil,omitempty"`
	AssessmentValidUntil   *time.Time `json:"assessmentValidUntil,omitempty"`
}

type WebRisk struct {
	Score    int      `json:"score,omitempty"`
	Status   string   `json:"status,omitempty"`
	Websites []string `json:"websites,omitempty"`
}

type RatingFilter struct {
	Status      []string      `url:"status"`
	ProjectType []ProjectType `url:"pt"`
	SupplierIDs []string      `url:"supplier-id"`
	ExternalIDs []string      `url:"external-id"`
}

type RequestRatingCRR struct {
	Category   string `json:"category"`
	SupplierID string `json:"supplierID,omitempty"`
	ExternalID string `json:"externalID,omitempty"`
}

type RequestRatingDPR struct {
	SupplierID string `json:"supplierID,omitempty"`
	ExternalID string `json:"externalID,omitempty"`
}

type ContingentUsage struct {
	Project          ProjectType `json:"project"`
	ContingentBefore int         `json:"contingentBefore"`
	ContingentAfter  int         `json:"contingentAfter"`
	OrderedSuppliers []string    `json:"orderedSuppliers"`
}

type PostSuppliersResponse struct {
	Created   int              `json:"created"`
	Existed   int              `json:"existed"`
	OrdersCRR *ContingentUsage `json:"ordersCRR,omitempty"`
	OrdersDPR *ContingentUsage `json:"ordersDPR,omitempty"`
}

type Account struct {
	ContingentCRR int `json:"contingentCRR,omitempty"`
	ContingentDPR int `json:"contingentDPR,omitempty"`
}
