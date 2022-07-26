package cyberrisk

type AssessmentStatus string

const (
	StatusAssessmentInProgress    AssessmentStatus = "in-progress"
	StatusAssessmentValidation    AssessmentStatus = "validation"
	StatusAssessmentJustification AssessmentStatus = "justification"
	StatusAssessmentRevalidation  AssessmentStatus = "revalidation"
	StatusAssessmentFinalization  AssessmentStatus = "finalization"
	StatusAssessmentDone          AssessmentStatus = "done"

	StatusAssessmentDeclined AssessmentStatus = "declined"
	StatusAssessmentBlocked  AssessmentStatus = "blocked"
	StatusAssessmentExpired  AssessmentStatus = "expired"
)
