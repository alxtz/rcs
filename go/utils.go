package ds

import "encoding/json"

func Readable(ds interface{}) string {
	j, _ := json.MarshalIndent(ds, "", "  ")
	return string(j)
}
