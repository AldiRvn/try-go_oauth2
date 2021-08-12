package util

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/kataras/golog"
)

func PrettyPrint(toJson interface{}) {
	byteJson, err := json.MarshalIndent(toJson, "", " ")
	golog.Debug(string(byteJson), "\nMarshalIndent error:", err)
}

func WriteJsonByMap(w http.ResponseWriter, mapData map[string]interface{}) {
	w.Header().Set("Content-type", "application/json")
	byteData, err := json.Marshal(mapData)
	if err != nil {
		golog.Errorf("json.Marshal:\n %v", err)
	}
	if _, err := w.Write(byteData); err != nil {
		golog.Errorf("w.Write:\n %v", err)
	}
}

func WriterOneLine(w http.ResponseWriter, contentToWrite string) {
	if _, err := w.Write([]byte(contentToWrite)); err != nil {
		golog.Errorf("w.Write:\n %v", err)
	}
}

func JsonUnmarshalOneLine(byteJson []byte) map[string]interface{} {
	newMap := map[string]interface{}{}
	err := json.Unmarshal(byteJson, &newMap)
	if err != nil {
		golog.Errorf("json.Unmarshal:\n %v", err)
	}
	return newMap
}

// ? Not used
func DecodeBase64(strBase64 string) string {
	afterDecode, err := base64.StdEncoding.DecodeString(strBase64)
	if err != nil {
		golog.Errorf("base64.StdEncoding.DecodeString:\n %v", err)
		golog.Debug(string(strBase64))
		return ""
	}
	return string(afterDecode)
}
