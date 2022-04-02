package enum

type EmailSubject int


const (
	TutoringSubjectToSelf EmailSubject = iota
	ContactFormSubjectToSelf
	SubjectToStudent
)
 
func (es EmailSubject) String() string {
	return [...]string{
	"Tutoring Inquiry",
	"Contact Form Submission",
	"Thanks for contacting me"}[es]
}
