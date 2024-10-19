package main

/*
#cgo LDFLAGS: -L/usr/local/lib -L/usr/lib64 -L../. -l_Site_Controller -lcjson -lfims -lspdlog -lstdc++
#include "../../git/hybridos/site_controller/c_interface/include.h"
*/
import "C"
import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var memeInstance = C.createAssetESS()

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
	// collect the form data
	ratedPwr, _ := strconv.ParseFloat(r.FormValue("ratedPower"), 64)
	socChgBegin, _ := strconv.ParseFloat(r.FormValue("socChgBegin"), 64)
	socChgEnd, _ := strconv.ParseFloat(r.FormValue("socChgEnd"), 64)
	socDischgBegin, _ := strconv.ParseFloat(r.FormValue("socDischgBegin"), 64)
	socDischgEnd, _ := strconv.ParseFloat(r.FormValue("socDischgEnd"), 64)
	socStep, _ := strconv.ParseFloat(r.FormValue("socStep"), 64)

	var chargeableFloatSlice = make([]float64, 0, 100)
	var dischargeableFloatSlice = make([]float64, 0, 100)
	var socValues []float64
	C.c_assign_derate_method(memeInstance, 0)
	C.c_assign_maint_mode(memeInstance, 0)
	C.c_assign_protection_buffers(memeInstance, 0)
	C.c_assign_rated_chargeable_power(memeInstance, C.float(ratedPwr))
	C.c_assign_rated_dischargeable_power(memeInstance, C.float(ratedPwr))
	C.c_assign_chargeable_power_raw(memeInstance, C.float(ratedPwr))
	C.c_assign_dischargeable_power_raw(memeInstance, C.float(ratedPwr*-1))
	C.c_assign_soc_derate_begin_end(memeInstance, C.float(socChgBegin), C.float(socChgEnd), C.float(socDischgBegin), C.float(socDischgEnd))

	for i := -10.0; i <= 110; i = i + socStep {
		C.c_assign_soc(memeInstance, C.float(i))
		C.c_limit_power(memeInstance)
		chargeable := C.c_get_chargeable_power(memeInstance)
		dischargeable := C.c_get_dischargeable_power(memeInstance)
		chargeableFloatSlice = append(chargeableFloatSlice, float64(chargeable))
		dischargeableFloatSlice = append(dischargeableFloatSlice, float64(dischargeable))
		socValues = append(socValues, i)
	}

	jsPayload := generateUpdateScript(socValues, chargeableFloatSlice, dischargeableFloatSlice)
	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprintf(w, jsPayload)
}

func generateUpdateScript(socValues []float64, chargeableSlice []float64, dischargeableSlice []float64) string {
	return fmt.Sprintf(`<script>updateChart(%s, [{ label: "ChargeablePwr", data: %v, hex_color: 'rgba(75, 192, 192, 1)' },{ label: "DischargeablePwr", data: %v, hex_color: 'rgba(75, 192, 192, 1)' }]);</script>`, slice_to_string(socValues), slice_to_string(chargeableSlice), slice_to_string(dischargeableSlice))
}

func slice_to_string(slice []float64) string {
	strData := make([]string, len(slice))
	for i, v := range slice {
		strData[i] = fmt.Sprintf("%f", v)
	}
	joinedData := strings.Join(strData, ", ")
	return "[" + joinedData + "]"
}
