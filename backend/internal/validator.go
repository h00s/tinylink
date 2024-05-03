package internal

import (
	"errors"
	"net"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func IsURLValid(urlToCheck string) error {
	u, err := url.ParseRequestURI(urlToCheck)
	if err != nil {
		return errors.New("INVALID_URL")
	}

	if !u.IsAbs() {
		return errors.New("URL_NOT_ABSOLUTE")
	}

	if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "ftp" {
		return errors.New("URL_NOT_HTTP")
	}

	host := strings.ToLower(u.Host)

	_, err = net.LookupHost(host)
	if err != nil {
		return errors.New("HOST_NOT_FOUND")
	}

	domain, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return errors.New("INVALID_DOMAIN")
	}

	if isWhitelisted(domain) {
		return nil
	}
	if isRedirector(domain) {
		return errors.New("REDIRECTOR_DOMAIN")
	}
	if isBlacklisted(domain) {
		return errors.New("BLACKLISTED_DOMAIN")
	}

	return nil
}

func isWhitelisted(domain string) bool {
	switch domain {
	case
		"google.com",
		"yahoo.com":
		return true
	}
	return false
}

func isRedirector(domain string) bool {
	switch domain {
	case
		"skrati.link",
		"tnylnk.in",
		"adf.ly",
		"bc.vc",
		"bit.do",
		"bit.ly",
		"budurl.com",
		"buff.ly",
		"clicky.me",
		"goo.gl",
		"is.gd",
		"mcaf.ee",
		"ow.ly",
		"s2r.co",
		"soo.gd",
		"short.to",
		"tiny.cc",
		"tinyurl.com":
		return true
	}
	return false
}

func isBlacklisted(domain string) bool {
	_, err := net.LookupHost(domain + ".multi.surbl.org")
	return err == nil
}
