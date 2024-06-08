package main

import (
	"encoding/json"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// ? fixed format json
	// // Create a fixed-format JSON response from a string. Notice how we're using a raw
	// // string literal (enclosed with backticks) so that we can include double-quote
	// // characters in the JSON without needing to escape them? We also use the %q verb to
	// // wrap the interpolated values in double-quotes.
	// js := `{"status": "available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)
	// // Set the "Content-Type: application/json" header on the response. If you forget to
	// // this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
	// // header instead.
	// w.Header().Set("Content-Type", "application/json")
	// // Write the JSON as the HTTP response body.
	// w.Write([]byte(js))

	// ? another approach | marshaling go struct or map or slice

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// js, err := json.Marshal(data)
	// if err != nil {
	// 	app.logger.Print(err)
	// 	http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	// 	return
	// }
	// // Append a newline to the JSON. This is just a small nicety to make it easier to
	// // view in terminal applications.
	// js = append(js, '\n')
	// // At this point we know that encoding the data worked without any problems, so we
	// // can safely set any necessary HTTP headers for a successful response.
	// w.Header().Set("Content-Type", "application/json")
	// // Use w.Write() to send the []byte slice containing the JSON as the response body.
	// w.Write(js)

	// ? helper function
	// err := app.writeJSON(w, http.StatusOK, data, nil)
	// if err != nil {
	// 	app.logger.Print(err)
	// 	http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	// }

	// ? json.Encoder

	// Set the "Content-Type: application/json" header on the response.
	w.Header().Set("Content-Type", "application/json")
	// Use the json.NewEncoder() function to initialize a json.Encoder instance that
	// writes to the http.ResponseWriter. Then we call its Encode() method, passing in
	// the data that we want to encode to JSON (which in this case is the map above). If
	// the data can be successfully encoded to JSON, it will then be written to our
	// http.ResponseWriter.

	// 	Imagine, for example, that you want to set a Cache-Control header on a successful
	// response, but notset a Cache-Control header if the JSON encoding fails and you have to
	// return an error response.

	// 	You could set the Cache-Control header and then delete it from the header map again in the
	// event of an error — but that’s pretty hacky.
	// Another option is to write the JSON to an interim bytes.Buffer instead of directly to the
	// http.ResponseWriter. You can then check for any errors, before setting the Cache-Control
	// header and copying the JSON from the bytes.Buffer to http.ResponseWriter. But once
	// you start doing that, it’s simpler and cleaner (as well as slightly faster) to use the alternative
	// json.Marshal() approach instead.
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
