package x

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return RandomStringWithCharset(length, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func RandomDigit(length int) string {
	return RandomStringWithCharset(length, "0123456789")
}

func IsDigit(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func IsPhone(s string) bool {
	return len(s) == 11 && s[0] == '1' && IsDigit(s)
}

func UniqueStringSlice(strings []string) []string {
	m := map[string]bool{}

	for _, s := range strings {
		m[s] = true
	}

	r := make([]string, 0)
	for _, s := range strings {
		if _, ok := m[s]; ok {
			r = append(r, s)
		}
	}
	return r
}

func UniqueIntSlice(strings []int64) []int64 {
	m := map[int64]bool{}

	for _, s := range strings {
		m[s] = true
	}

	r := make([]int64, 0)
	for _, s := range strings {
		if _, ok := m[s]; ok {
			r = append(r, s)
		}
	}
	return r
}

func GenderName(g int) string {
	switch g {
	case 1:
		return "男"
	case 2:
		return "女"
	default:
		return "未知"
	}
}

func BoolName(b bool) string {
	switch b {
	case false:
		return "否"
	case true:
		return "是"
	default:
		return "未知"
	}
}

func RemoveStringSliceItem(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func DiffStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func DiffIntSlice(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func StripStringArray(list []string) []string {
	r := make([]string, 0)
	for _, s := range list {
		s = strings.TrimSpace(s)
		if s != "" {
			r = append(r, s)
		}
	}
	return r
}

func ParseSortParam(sort string) (string, string) {
	if sort == "" {
		return "", ""
	}
	seps := strings.Fields(sort)
	if len(seps) != 2 {
		return "", ""
	}
	field := seps[0]
	order := seps[1]
	if order == "descend" {
		order = "DESC"
	} else {
		order = "ASC"
	}
	return field, order
}
