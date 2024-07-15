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

const SCORE_CRR ScoreType = "crr"
const SCORE_DPR ScoreType = "dpr"
const SCORE_WEBRISK ScoreType = "webrisk"
