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