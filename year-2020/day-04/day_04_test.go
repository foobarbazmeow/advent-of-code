package day_04

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var (
	re1, _ = regexp.Compile("^#[a-fA-F\\d]{6}$")
	re2, _ = regexp.Compile("^\\d{9}$")
)

func Test04Part1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day_04.in")
	assert.Equal(t, 182, validate(string(bytes)))
}

func Test04Part2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day_04.in")
	assert.Equal(t, 109, validateWithFelds(string(bytes)))
}

func validate(str string) int {
	result := 0
	for _, passport := range strings.Split(str, "\n\n") {
		fields := strings.Split(strings.Join(strings.Split(passport, "\n"), " "), " ")
		m := map[string]struct{}{}
		for _, field := range fields {
			parts := strings.Split(field, ":")
			m[parts[0]] = struct{}{}
		}
		if len(m) == 8 {
			result += 1
		} else {
			if _, ok := m["cid"]; len(m) == 7 && !ok {
				result += 1
			}
		}
	}
	return result
}

func validateWithFelds(str string) int {
	validateFields := func(key, val string) bool {
		switch key {
		case "byr":
			{
				return validate_byr(val)
			}
		case "iyr":
			{
				return validate_iyr(val)
			}
		case "eyr":
			{
				return validate_eyr(val)
			}
		case "hgt":
			{
				return validate_hgt(val)
			}
		case "hcl":
			{
				return validate_hcl(val)
			}
		case "ecl":
			{
				return validate_ecl(val)
			}
		case "pid":
			{
				return validate_pid(val)
			}
		case "cid":
			{
				return true
			}
		default:
			{
				return false
			}
		}
	}

	result := 0
	for _, passport := range strings.Split(str, "\n\n") {
		fields := strings.Split(strings.Join(strings.Split(passport, "\n"), " "), " ")
		m := map[string]struct{}{}
		for _, field := range fields {
			parts := strings.Split(field, ":")
			if !validateFields(parts[0], parts[1]) {
				continue
			}
			m[parts[0]] = struct{}{}
		}
		if len(m) == 8 {
			result += 1
		} else {
			if _, ok := m["cid"]; len(m) == 7 && !ok {
				result += 1
			}
		}
	}
	return result
}

func validate_pid(val string) bool {
	return re2.MatchString(val)
}

func validate_ecl(val string) bool {
	m := map[string]struct{}{
		"amb": {}, "blu": {}, "brn": {}, "gry": {}, "grn": {}, "hzl": {}, "oth": {},
	}

	_, ok := m[val]
	return ok
}

func validate_hcl(val string) bool {
	return re1.MatchString(val)
}

func validate_hgt(val string) bool {
	if len(val) != 5 && len(val) != 4 {
		return false
	}
	m := val[len(val)-2:]
	num, err := strconv.Atoi(val[:len(val)-2])
	if m == "cm" {
		return err == nil && num >= 150 && num <= 193
	} else if m == "in" {
		return err == nil && num >= 59 && num <= 76
	}
	return false
}

func validate_eyr(val string) bool {
	num, err := strconv.Atoi(val)
	return err == nil && num >= 2020 && num <= 2030
}

func validate_iyr(val string) bool {
	num, err := strconv.Atoi(val)
	return err == nil && num >= 2010 && num <= 2020
}

func validate_byr(val string) bool {
	num, err := strconv.Atoi(val)
	return err == nil && num >= 1920 && num <= 2002
}
