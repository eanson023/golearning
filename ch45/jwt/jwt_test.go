package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"testing"
	"time"
)

// 创建token
func CreateToken(SecretKey []byte, issuer string, subject string, isAdmin bool) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":     issuer,
		"sub":     subject,
		"isAdmin": isAdmin,
		// 生效时间
		"nbf": time.Now().Unix(),
		// 过期时间
		"exp": int64(time.Now().Add(time.Minute * 30).Unix()),
	})
	tokenString, err = token.SignedString(SecretKey)
	return
}

//解析token
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}

type MyCustomClaims struct {
	issuser string
	subject string
	isAdmin bool
}

func TestToken(t *testing.T) {
	sercretKey := []byte("idghsu&*%hvukvuy27t7&%#^&@fvidjsj")
	issuser := "eanson"
	uuid := uuid.Must(uuid.NewV4()).String()
	isAdmin := false
	tokenStr, _ := CreateToken(sercretKey, issuser, uuid, isAdmin)
	t.Log(tokenStr)
	if claims, err := ParseToken(tokenStr, sercretKey); err != nil {
		t.Fatal(err)
	} else {
		mapClaims := claims.(jwt.MapClaims)
		t.Log(mapClaims)
	}
}

func TestUUID(t *testing.T) {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)
	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err = uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
