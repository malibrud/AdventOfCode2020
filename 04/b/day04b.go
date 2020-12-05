package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr     int
	iyr     int
	eyr     int
	hgt     int
	hgtu    string
	hcl     string
	ecl     string
	pid     string
	cid     string
	parseOK bool
}

func main() {
	fmt.Println(doit("data.txt"))
}

func doit(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	passports := strings.Split(string(dat), "\r\n\r\n")

	validCount := 0
	for _, rawPassport := range passports {
		p := newPassport()
		p.parse(rawPassport)
		valid := p.validate()
		if valid {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}

func newPassport() passport {
	p := passport{0, 0, 0, 0, "", "", "", "", "", true}
	return p
}

func (p *passport) parse(raw string) {
	fields := strings.Fields(raw)
	for _, f := range fields {
		p.parseField(f)
	}
}

func (p *passport) validate() bool {
	// Check for successful parse
	if !p.parseOK {
		return false
	}

	// Birth year
	if !(1920 <= p.byr && p.byr <= 2002) {
		return false
	}

	// Issue year
	if !(2010 <= p.iyr && p.iyr <= 2020) {
		return false
	}

	// Expiration year
	if !(2020 <= p.eyr && p.eyr <= 2030) {
		return false
	}

	// Height
	switch p.hgtu {
	case "":
		return false
	case "cm":
		if !(150 <= p.hgt && p.hgt <= 193) {
			return false
		}
	case "in":
		if !(59 <= p.hgt && p.hgt <= 76) {
			return false
		}
	}

	// Hair color - checked in the parse
	if p.hcl == "" {
		return false
	}

	// Eye color
	vEcl := false
	validEcls := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, ecl := range validEcls {
		if p.ecl == ecl {
			vEcl = true
		}
	}
	if !vEcl {
		return false
	}

	// Passport ID - Checked in the parse
	if p.pid == "" {
		return false
	}

	return true
}

func (p *passport) parseField(f string) {
	kv := strings.Split(f, ":")
	k := kv[0]
	v := kv[1]
	e := errors.New("")

	switch k {
	case "byr":
		p.byr, e = strconv.Atoi(v)
		if e != nil {
			p.parseOK = false
		}
	case "iyr":
		p.iyr, e = strconv.Atoi(v)
		if e != nil {
			p.parseOK = false
		}
	case "eyr":
		p.eyr, e = strconv.Atoi(v)
		if e != nil {
			p.parseOK = false
		}
	case "hgt":
		p.hgt, p.hgtu, e = parseHgt(v)
		if e != nil {
			p.parseOK = false
		}
	case "hcl":
		p.hcl, e = parseHcl(v)
		if e != nil {
			p.parseOK = false
		}
	case "ecl":
		p.ecl = v
	case "pid":
		p.pid, e = parsePid(v)
		if e != nil {
			p.parseOK = false
		}
	case "cid":
	}
}

func parseHgt(s string) (int, string, error) {
	h := -1
	u := ""

	re := regexp.MustCompile("^([0-9]+)(cm|in)$")
	match := re.FindStringSubmatch(s)
	if match == nil {
		return h, u, errors.New("Could not parse height")
	}
	h, e := strconv.Atoi(match[1])
	u = match[2]
	return h, u, e
}

func parseHcl(s string) (string, error) {

	re := regexp.MustCompile("^#([0-9a-f]{6})$")
	match := re.FindStringSubmatch(s)
	if match == nil {
		return "", errors.New("Could not parse Hlc")
	}
	return match[1], nil
}

func parsePid(s string) (string, error) {
	re := regexp.MustCompile("^([0-9]{9})$")
	match := re.FindStringSubmatch(s)
	if match == nil {
		return "", errors.New("Could not parse Pid")
	}
	return match[1], nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
