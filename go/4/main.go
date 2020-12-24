package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}
	passports := strings.Split(string(contents), "\n\n")

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	// optionalFields := []string{"cid"}

	numValid := 0
	for _, p := range passports {
		fields := extractFields(p)
		if hasAllRequired(fields, requiredFields) {
			numValid++
		}
	}

	fmt.Println(numValid)
}

func hasAllRequired(fields map[string]struct{}, requiredFields []string) bool {
	for _, req := range requiredFields {
		if _, ok := fields[req]; !ok {
			return false
		}
	}
	return true
}

func extractFields(passport string) map[string]struct{} {
	fields := map[string]struct{}{}

	passportElements := strings.FieldsFunc(passport, func(r rune) bool {
		return r == '\n' || r == ' '
	})
	fmt.Println(passportElements)
	for _, el := range passportElements {
		f := strings.Split(el, ":")
		name := string(f[0])
		value := string(f[1])
		if validateField(name, value) {
			fields[name] = struct{}{}
		} else {
			fmt.Printf("INVALID: %s (%s)\n", name, value)
		}
	}

	return fields
}

func validateField(name, value string) bool {
	type validator func(string) bool
	validators := map[string]validator{
		"byr": func(s string) bool {
			y, _ := strconv.ParseInt(s, 10, 64)
			if y <= 2002 && y >= 1920 {
				return true
			}
			return false
		},
		"iyr": func(s string) bool {
			y, _ := strconv.ParseInt(s, 10, 64)
			if y <= 2020 && y >= 2010 {
				return true
			}
			return false
		},
		"eyr": func(s string) bool {
			y, _ := strconv.ParseInt(s, 10, 64)
			if y <= 2030 && y >= 2020 {
				return true
			}
			return false
		},
		"hgt": func(s string) bool {
			if x := strings.Index(s, "cm"); x > 0 {
				h, _ := strconv.ParseInt(s[:x], 10, 64)
				if h <= 193 && h >= 150 {
					return true
				}
				return false
			}
			if x := strings.Index(s, "in"); x > 0 {
				h, _ := strconv.ParseInt(s[:x], 10, 64)
				if h <= 76 && h >= 59 {
					return true
				}
				return false
			}
			return false
		},
		"hcl": func(s string) bool {
			ss := strings.Split(s, "#")
			if ss[0] != "" || len(ss[1]) != 6 {
				return false
			}
			for _, r := range ss[1] {
				if (r <= '9' && r >= '0') || (r <= 'f' && r >= 'a') {
					continue
				}
				return false
			}
			return true
		},
		"ecl": func(s string) bool {
			return s == "amb" ||
				s == "blu" ||
				s == "brn" ||
				s == "gry" ||
				s == "grn" ||
				s == "hzl" ||
				s == "oth"
		},
		"pid": func(s string) bool {
			if len(s) != 9 {
				return false
			}
			if _, err := strconv.ParseInt(s, 10, 64); err != nil {
				return false
			}
			return true
		},
	}
	if validator, ok := validators[name]; !ok || !validator(value) {
		return false
	}
	return true
}
