package validation

import "regexp"

const (
	K8S_NAME_PATTERN   = `^[a-z]([-a-z0-9]*[a-z0-9])?$`
	K8S_DOMAIN_PATTERN = `^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z]{2,20}$`
)

var (
	// Compile the regular expression
	reNamePattern   = regexp.MustCompile(K8S_NAME_PATTERN)
	reDomainPattern = regexp.MustCompile(K8S_DOMAIN_PATTERN)
)

// Kubernetes namespace name must match this regular expression:
// ^[a-z]([-a-z0-9]*[a-z0-9])?$
// - Must start with a lowercase letter or digit.
// - Can contain lowercase letters, digits, and dashes.
// - Must end with a lowercase letter or digit.
// - Can be up to 63 characters long.
func K8sCheckName(name string) bool {
	// check long
	if len(name) < 1 || len(name) > 63 {
		return false
	}

	// Check if the name matches the pattern
	return reNamePattern.MatchString(name)
}

func K8sCheckDomain(domain string) bool {
	// check long
	if len(domain) < 4 || len(domain) > 63 {
		return false
	}

	// Check if the name matches the pattern
	return reDomainPattern.MatchString(domain)
}
