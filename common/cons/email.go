package cons

const EmailSendingFailed = "Email sending failed"

const (
	EmailExists = 1 + iota
	EmailNotExists
)

const (
	EmailExistsMsg    = "Email already exists"
	EmailNotExistsMsg = "Email not exists"
)
