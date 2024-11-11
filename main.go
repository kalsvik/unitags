package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

const (
	Red   = "\033[31m"
	White = "\033[97m"
)

type TrackPosition struct {
	Latitude  float32
	Longitude float32
	TimeStamp int64
}

type TrackAddress struct {
	FormattedAddressLines []string
}

type TrackRole struct {
	Emoji string
}

type TrackData struct {
	Name     string
	Battery  int
	Address  TrackAddress
	Location TrackPosition
	Role     TrackRole
}

func getFindMyData() []TrackData {
	home := os.Getenv("HOME")
	items, err := os.ReadFile(home + "/Library/Caches/com.apple.findmy.fmipcore/Items.data")
	if err != nil {
		println(Red + "ERROR:" + White + " Could not read Items Data... Make sure the Find My app is open.")
		panic(err)
	}
	var itemsTrackData []TrackData
	json.NewDecoder(bytes.NewReader(items)).Decode(&itemsTrackData)
	var devicesTrackData []TrackData
	devices, err := os.ReadFile(home + "/Library/Caches/com.apple.findmy.fmipcore/Devices.data")
	if err != nil {
		println(Red + "ERROR:" + White + " Could not read Devices Data... Make sure the Find My app is open.")
		panic(err)
	}
	json.NewDecoder(bytes.NewReader(devices)).Decode(&devicesTrackData)

	var mergedTrackData []TrackData

	mergedTrackData = append(mergedTrackData, itemsTrackData...)
	mergedTrackData = append(mergedTrackData, devicesTrackData...)

	return mergedTrackData
}

func Init() {
	pass, err := os.ReadFile("password.txt")
	if err != nil {
		println(Red + "ERROR:" + White + " Could not read 'password.txt'. Please make sure that this file exists and contains a password.")
		return
	}

	template, err := os.ReadFile("public/index.html")
	if err != nil {
		println(Red + "ERROR:" + White + " HTML served to front-end could not be found (target/public/index.html).")
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		buff, err := io.ReadAll(r.Body)
		if err != nil {
			println(Red + "ERROR:" + White + " Could not get request body.")
			http.Error(w, "could not get request body", http.StatusForbidden)
			return
		}
		if string(buff) != string(pass) {
			println(Red + "ERROR:" + White + " Input password did not match predicate. (" + string(buff) + ") -> " + string(pass))
			http.Error(w, "authentication failed", http.StatusForbidden)
			return
		}
		writer := new(bytes.Buffer)
		json.NewEncoder(writer).Encode(getFindMyData())
		web := string(template)
		web = strings.ReplaceAll(web, `øødataøø`, writer.String())
		w.Write([]byte(web))
	})

	handle := cors.Default().Handler(mux)

	err = http.ListenAndServe(":4849", handle)
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()
}
