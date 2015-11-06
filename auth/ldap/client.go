package ldap

import (
	"crypto/tls"
	"errors"
	"fmt"

	ldaplib "github.com/go-ldap/ldap"
)

type LDAPUser struct {
	Username string
	DistinguishedName string
	Emails []string
	ProfileNames []string
}

func connectLDAP() (l *ldaplib.Conn, err error) {
	addr := fmt.Sprintf("%s:%d", Config.Host, Config.Port)
	if Config.TLS {
		l, err = ldaplib.DialTLS("tcp", addr, &tls.Config{InsecureSkipVerify: true})
	} else {
		l, err = ldaplib.Dial("tcp", addr)
	}
	return
}

func VerifyConfig() error {
	bindusername := Config.SearchUser
	bindpassword := Config.SearchPassword

	l, err := connectLDAP()
	if err != nil {
		return err
	}
	defer l.Close()

	// Verify that bind works for search user
	return l.Bind(bindusername, bindpassword)
}

func VerifyLogin(username, password string) (*LDAPUser, error) {
	bindusername := Config.SearchUser
	bindpassword := Config.SearchPassword

	l, err := connectLDAP()
	if err != nil {
		return nil, err
	}
	defer l.Close()

	// First bind with a read only user
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		return nil, err
	}

	var filterTemplate string
	if Config.Filter != "" {
		filterTemplate = fmt.Sprintf("&(%s)", Config.Filter)
	}
	queryTemplate := fmt.Sprintf("(%s&(%s=%s))", filterTemplate, Config.UserIDField, username)

	attributes := []string{"dn"}
	if Config.EmailField != "" {
		attributes = append(attributes, Config.EmailField)
	}
	if Config.ProfileNameField != "" {
		attributes = append(attributes, Config.ProfileNameField)
	}

	// Search for the given username
	searchRequest := ldaplib.NewSearchRequest(
		Config.DomainBase,
		ldaplib.ScopeWholeSubtree, ldaplib.NeverDerefAliases, 0, 0, false,
		queryTemplate,
		attributes,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	if len(sr.Entries) != 1 {
		return nil, errors.New("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	emails := sr.Entries[0].GetAttributeValues(Config.EmailField)

	if len(emails) == 0 {
		return nil, errors.New("email not found for LDAP user")
	}

	// Bind as the user to verify their password
	err = l.Bind(userdn, password)
	if err != nil {
		return nil, err
	}

	// Rebind as the read only user for any futher queries
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		return nil, err
	}

	return &LDAPUser{Username: username, DistinguishedName: userdn, Emails: emails}, nil
}
