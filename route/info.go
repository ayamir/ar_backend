package route

import "fmt"

type Info struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Motto       string `json:"motto"`
	LastUpdated string `json:"last_updated"`
}

type InfoMotto struct {
	Code  string `json:"code"`
	Motto string `json:"motto"`
}

func (info Info) toString() string {
	return fmt.Sprintf(
		"Info: code=%s, name=%s, motto=%s, last_updated=%s.\n",
		info.Code, info.Name, info.Motto, info.LastUpdated)
}
