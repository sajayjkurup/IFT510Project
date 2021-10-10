package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iancoleman/strcase"
)

type Response struct {
	Input  string `json:"input"`
	Output string `json:"output"`
	Format string `json:"format"`
}

func startServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.Path("/supportedformats").HandlerFunc(handleSupportedTypesRequestt)
	router.Path("/{str}/convert").Queries("format", "{format}").HandlerFunc(handleCaseConvertRequest)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleSupportedTypesRequestt(w http.ResponseWriter, r *http.Request) {
	formats := []string{
		"SNAKE",
		"SCREAMINGSNAKE",
		"KEBAB",
		"SCREAMINGKEBAB",
		"CAMEL",
		"LOWERCAMEL",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formats)
}

func handleCaseConvertRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	input := vars["str"]
	format := vars["format"]
	output, err := convertCase(format, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "text/plain")
	response := new(Response)
	response.Format = format
	response.Input = input
	response.Output = output
	json.NewEncoder(w).Encode(response)
}

func convertCase(format string, str string) (string, error) {
	switch format {
	case "SNAKE":
		return strcase.ToCamel(str), nil
	case "SCREAMINGSNAKE":
		return strcase.ToScreamingSnake(str), nil
	case "KEBAB":
		return strcase.ToKebab(str), nil
	case "SCREAMINGKEBAB":
		return strcase.ToScreamingKebab(str), nil
	case "CAMEL":
		return strcase.ToCamel(str), nil
	case "LOWERCAMEL":
		return strcase.ToLowerCamel(str), nil
	default:
		return "", errors.New("Unsupported Type. Supported types are SNAKE, SCREAMINGSNAKE, KEBAB, SCREAMINGKEBAB, CAMEL, LOWERCAMEL")
	}

}

func main() {
	startServer()
}
