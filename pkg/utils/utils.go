package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
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

func UnmarshalYamlFile(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
