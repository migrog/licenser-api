package entity

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type License struct {
	Id         primitive.ObjectID `bson:"_id"`
	Enabled    bool               `bson:"enabled"`
	LicenseKey string             `bson:"license_key"`
	Product    *Product           `bson:"product"`
	User       *User              `bson:"user"`
	Date       time.Time          `bson:"date"`
}

func NewLicense() *License {
	return &License{
		Id:   primitive.NewObjectID(),
		Date: time.Now(),
	}
}

func (d *License) AddProduct(m *Product) {
	d.Product = m
}

func (d *License) AddUser(m *User) {
	d.User = m
}

func (d *License) Enable() {
	d.Enabled = true
}

func (d *License) Disable() {
	d.Enabled = false
}

func (d *License) NewLicenseKey() {
	s, _ := generateRandomString(20)
	k := licenseKeyFormatting(s, 5)
	d.LicenseKey = k
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func licenseKeyFormatting(s string, k int) string {
	countAN := len(s)

	if countAN == 0 {
		return ""
	}
	s = strings.ToUpper(s)

	res := make([]byte, (countAN+k-1)/k-1+countAN)

	i, j := len(res), len(s)
	for 0 <= j-k {
		copy(res[i-k:i], s[j-k:j])

		if 0 <= i-k-1 {
			res[i-k-1] = '-'
		}

		i -= k + 1
		j -= k
	}

	if j > 0 {
		copy(res[:j], s[:j])
	}

	return string(res)
}
