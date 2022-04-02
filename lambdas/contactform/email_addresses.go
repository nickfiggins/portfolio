package contactform

type EmailAddress int
 
const (
	NoReplyEmail EmailAddress = iota
	PersonalEmail
	TutoringEmail
)
 
func (e EmailAddress) String() string {
	return [...]string{
	"no-reply@nickfiggins.com",
	"figginsn@gmail.com",
	"nickfigginstutoring@gmail.com"}[e]
}