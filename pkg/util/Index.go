package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
)

var CstZone = time.FixedZone("CST", 8*3600) // 东八

func GenerateDatabaseID() string {
	uid := ksuid.New()
	return uid.String()
}

func GetTimeFromDatabaseId(id string) (t time.Time, err error) {
	kid, err := ksuid.Parse(id)
	if err != nil {
		return
	}
	t = kid.Time()
	return
}

func GenerateRequestID() string {
	uid := xid.New()
	return uid.String()
}

func GetTimeFromRequestId(id string) (t time.Time, err error) {
	x, err := xid.FromString(id)
	if err != nil {
		return
	}
	t = x.Time()
	return
}

func Unicode2Zh(form string) (to string, err error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(form), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}

// SnakeString XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// CamelString xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func HmacSha256(secret string, signString string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signString))
	encrypted := hex.EncodeToString(h.Sum(nil))
	return encrypted
}
