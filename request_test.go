package whois_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/linzeyan/whois"
)

const domain string = "google.com"

func TestWhoisXMLApi(t *testing.T) {
	result, err := whois.RequestWhoisXML(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(whois.ParserWhoisXML(result), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(out))
}

func TestIp2Whois(t *testing.T) {
	result, err := whois.RequestIp2Whois(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(whois.ParserIp2Whois(result), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(out))
}

func TestWhoApi(t *testing.T) {
	result, err := whois.RequestWhoApi(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(whois.ParserWhoApi(result), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(out))
}

func TestApiNinjas(t *testing.T) {
	result, err := whois.RequestApiNinjas(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(whois.ParserApiNinjas(result), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(out))
}

func TestVerisign(t *testing.T) {
	result, err := whois.RequestVerisign(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(whois.ParserVerisign(result), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
}

func TestIana(t *testing.T) {
	result, err := whois.RequestIana(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(whois.ParserIana(result), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
}

func TestDefault(t *testing.T) {
	resultVer, err := whois.RequestVerisign(domain)
	if err != nil {
		log.Println(err)
		return
	}
	outVer := whois.ParserVerisign(resultVer)

	result, err := whois.RequestIana(domain)
	if err != nil {
		log.Println(err)
		return
	}
	org := whois.ParserIana(result)
	outVer["Registrant"] = org
	out, err := json.MarshalIndent(outVer, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
}
