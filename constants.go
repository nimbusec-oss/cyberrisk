package cyberrisk

type RatingStatus string

const (
	StatusRatingInProgress    RatingStatus = "in-progress"
	StatusRatingValidation    RatingStatus = "validation"
	StatusRatingJustification RatingStatus = "justification"
	StatusRatingRevalidation  RatingStatus = "revalidation"
	StatusRatingFinalization  RatingStatus = "finalization"
	StatusRatingDone          RatingStatus = "done"

	StatusRatingDeclined RatingStatus = "declined"
	StatusRatingBlocked  RatingStatus = "blocked"
	StatusRatingExpired  RatingStatus = "expired"
)
