package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func PrintDebug(variable any, comment string) {
	fmt.Printf("!!! %s => %#v\n", comment, variable)
}

func GenerateRandData(length int) []byte {
	res := make([]byte, length)
	if _, err := rand.Read(res); err != nil {
		logrus.Fatalf("cannot generate token: %v", err)
	}
	return res
}

func GetIP(req *http.Request) net.IP {
	return net.ParseIP(echo.ExtractIPDirect()(req))
}

func PrettyPrint(i any) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
