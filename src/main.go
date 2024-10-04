package main

/*
#cgo LDFLAGS: -L../cppsrc -lmeme
#include "../cppsrc/meme.h"
#include "../cppsrc/meme_class.h"
*/
import "C"
import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var floatSlice = make([]float64, 100)
var memeInstance = C.createMeme(0.0)

func main() {
	http.HandleFunc("/", servePage)
	http.HandleFunc("/style/style.css", serveCSS)
	http.HandleFunc("/get_chart_data", getChart)
	http.ListenAndServe(":8080", nil)
}

func servePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/meme.html"))
	tmpl.Execute(w, nil)
}

func serveCSS(w http.ResponseWriter, r *http.Request) {
	style, _ := os.ReadFile("style/style.css")
	w.Header().Set("Content-Type", "text/css")
	w.Write(style)
}

func getChart(w http.ResponseWriter, r *http.Request) {
	if rand.Intn(2) == 0 {
		C.incrementMeme(memeInstance)
	} else {
		C.decrementMeme(memeInstance)
	}
	currentData := float64(C.getMemeData(memeInstance))
	copy(floatSlice, floatSlice[1:])
	floatSlice[len(floatSlice)-1] = currentData
	jsPayload := generateUpdateScript()
	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprintf(w, jsPayload)
}

func generateUpdateScript() string {
	return fmt.Sprintf(`<script>updateChart([{ src_uri: "meme", key: "data", data: %v, hex_color: 'rgba(75, 192, 192, 1)' }]);</script>`, slice_to_string(floatSlice))
}

func slice_to_string(slice []float64) string {
	strData := make([]string, len(slice))
	for i, v := range slice {
		strData[i] = fmt.Sprintf("%f", v)
	}
	joinedData := strings.Join(strData, ", ")
	return "[" + joinedData + "]"
}
