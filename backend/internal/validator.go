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
		return errors.New("invalid URL")
	}

	if !u.IsAbs() {
		return errors.New("URL is not absolute")
	}

	if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "ftp" {
		return errors.New("URL does not have http(s) prefix")
	}

	host := strings.ToLower(u.Host)

	_, err = net.LookupHost(host)
	if err != nil {
		return errors.New("host doesn't exist")
	}

	domain, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return errors.New("error while getting domain")
	}

	if isWhitelisted(domain) {
		return nil
	}
	if isRedirector(domain) {
		return errors.New("domain found in redirectors list")
	}
	if isBlacklisted(domain) {
		return errors.New("domain is blacklisted")
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
