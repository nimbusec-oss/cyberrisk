package cyberrisk

const (
	StatusRatingInProgress    = "in-progress"
	StatusRatingValidation    = "validation"
	StatusRatingJustification = "justification"
	StatusRatingRevalidation  = "revalidation"
	StatusRatingFinalization  = "finalization"
	StatusRatingDone          = "done"
	StatusRatingDeclined      = "declined"
	StatusRatingBlocked       = "blocked"
	StatusRatingExpired       = "expired"
)

const (
	StatusCScoreInProgress = "in-progress"
	StatusCScoreDone       = "done"
	StatusCScoreFailed     = "failed"
	StatusCScoreMissing    = "missing"
)

const TYPE_CRR ProjectType = "crr"
const TYPE_DPR ProjectType = "dpr"
const TYPE_WEBRISK ProjectType = "webrisk"
const TYPE_DORA ProjectType = "dora"

const PROVIDED_BY_GROUP ProvidedBy = "group"
const PROVIDED_BY_SUPPLIER ProvidedBy = "supplier"

const IDENTIFIER_TYPE_LEI IdentifierType = "LEI"
const IDENTIFIER_TYPE_EUID IdentifierType = "EUID"
const IDENTIFIER_TYPE_VAT IdentifierType = "VAT"
const IDENTIFIER_TYPE_CRN IdentifierType = "CRN"
const IDENTIFIER_TYPE_PNR IdentifierType = "PNR"
const IDENTIFIER_TYPE_NIN IdentifierType = "NIN"
const IDENTIFIER_TYPE_KSVID IdentifierType = "KSVID"
