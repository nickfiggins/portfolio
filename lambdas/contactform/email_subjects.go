package contactform

type EmailSubject int


const (
	TutoringSubjectToSelf EmailSubject = iota
	ContactFormSubjectToSelf
	SubjectToStudent
)
 
func (es EmailSubject) String() string {
	return [...]string{
	"Tutoring Inquiry",
	"Contact Form Submission from %s",
	"Thanks for contacting me"}[es]
}
