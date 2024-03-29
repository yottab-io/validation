package validation

import (
	"errors"
	"regexp"
)

var (
	wordRe     = regexp.MustCompile(`^\w{0,32}$`) // used for log category
	userRe     = regexp.MustCompile(`^[a-z]([a-z0-9]){3,19}$`)
	emailRe    = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	codeNumber = regexp.MustCompile(`^\d{0,128}$`)
	phonNumber = regexp.MustCompile(`^[\-+\d]{0,16}$`)
	notPass    = regexp.MustCompile(`^(.{0,9}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`) //  ^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*~_+=|:;?,<>]).{10,64}$
	uuidRe     = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	domainRe   = regexp.MustCompile(`^[a-z0-9.\\-]{4,61}$`)
)

var (
	ErrBadWorkspace = errors.New("Err Bad Workspace")
	ErrBadUser      = errors.New("Err Bad User")
	ErrBadEmail     = errors.New("Err Bad Email")
	ErrBadPassword  = errors.New("Err Bad Password")
	ErrBadUUID      = errors.New("Err Bad UUID")
)

func ValidationWord(w string) bool           { return wordRe.MatchString(w) }
func ValidationWorkspaceName(ws string) bool { return userRe.MatchString(ws) }
func ValidationAppName(app string) bool      { return userRe.MatchString(app) }
func ValidationCompanyName(co string) bool   { return userRe.MatchString(co) }
func ValidationEmail(email string) bool      { return emailRe.MatchString(email) }
func ValidationCodeNumber(code string) bool  { return codeNumber.MatchString(code) }
func ValidationPhonNumber(num string) bool   { return phonNumber.MatchString(num) }
func ValidationPass(pass string) bool        { return !notPass.MatchString(pass) }
func ValidationUUID(uuid string) bool        { return uuidRe.MatchString(uuid) }
func ValidationDomain(domain string) bool    { return domainRe.MatchString(domain) }

func CheckWorkspaceName(ws string) error {
	if !ValidationWorkspaceName(ws) {
		return ErrBadWorkspace
	}

	return nil
}

func CheckSignupRequest(user, email, pass string) error {
	if !ValidationCompanyName(user) {
		return ErrBadUser
	} else if !ValidationEmail(email) {
		return ErrBadEmail
	} else if !ValidationPass(pass) {
		return ErrBadPassword
	}

	return nil
}

func CheckTokenRequest(user, pass string) error {
	if !ValidationCompanyName(user) {
		return ErrBadUser
	} else if !ValidationPass(pass) {
		return ErrBadPassword
	}

	return nil
}

func CheckResetPassRequest(newPass, oldPass string) error {
	if !ValidationPass(newPass) {
		return ErrBadPassword
	} else if !ValidationPass(oldPass) {
		return ErrBadPassword
	}

	return nil
}

func CheckNewPassRequest(user, newPass, uuidCode string) error {
	if !ValidationPass(newPass) {
		return ErrBadPassword
	} else if !ValidationCompanyName(user) {
		return ErrBadUser
	} else if !ValidationUUID(uuidCode) {
		return ErrBadUUID
	}

	return nil
}
