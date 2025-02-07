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
