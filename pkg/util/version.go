package util

import "encoding/json"

var (
	Name      = "Not Provide"
	Version   = "Not Provide"
	BuildTime = "Not Provide"
	GitSHA1   = "Not Provide"
)

func VersionInfo() []byte {
	v := map[string]string{
		"Name":      Name,
		"Version":   Version,
		"BuildTime": BuildTime,
		"GitSHA1":   GitSHA1,
	}
	d, _ := json.MarshalIndent(v, "", "    ")
	return d
}
