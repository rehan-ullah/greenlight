package main

// func BenchmarkEncoder(b *testing.B) {
// 	w := httptest.NewRecorder()
// 	r := new(http.Request)
// 	for n := 0; n < b.N; n++ {
// 		healthcheckHandlerEncoder(w, r)
// 	}
// }

// func BenchmarkMarshal(b *testing.B) {
// 	w := httptest.NewRecorder()
// 	r := new(http.Request)
// 	for n := 0; n < b.N; n++ {
// 		healthcheckHandlerMarshal(w, r)
// 	}
// }

// func BenchmarkMarshalIndent(b *testing.B) {
// 	w := httptest.NewRecorder()
// 	r := new(http.Request)
// 	for n := 0; n < b.N; n++ {
// 		healthcheckHandlerMarshalIndent(w, r)
// 	}
// }

// func BenchmarkMarshal(b *testing.B) {
// 	w := httptest.NewRecorder()
// 	r := new(http.Request)
// 	for n := 0; n < b.N; n++ {
// 		healthcheckHandlerMarshal(w, r)
// 	}
// }

// func createMovieHandlerUnmarshal(w http.ResponseWriter, r *http.Request) {
// 	var input struct {
// 		Title   string   `json:"title"`
// 		Year    int32    `json:"year"`
// 		Runtime int32    `json:"runtime"`
// 		Genres  []string `json:"genres"`
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.Unmarshal(body, &input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }

// func createMovieHandlerDecoder(w http.ResponseWriter, r *http.Request) {
// 	var input struct {
// 		Title   string   `json:"title"`
// 		Year    int32    `json:"year"`
// 		Runtime int32    `json:"runtime"`
// 		Genres  []string `json:"genres"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }

// const body = `{"title":"Moana","year":2016,"runtime":107, "genres":["animation","adventure"]}`

// func BenchmarkUnmarshal(b *testing.B) {
// 	w := httptest.NewRecorder()

// 	for n := 0; n < b.N; n++ {
// 		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(body))
// 		if err != nil {
// 			b.Fatal(err)
// 		}

// 		createMovieHandlerUnmarshal(w, r)
// 	}
// }

// func BenchmarkDecoder(b *testing.B) {
// 	w := httptest.NewRecorder()

// 	for n := 0; n < b.N; n++ {
// 		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(body))
// 		if err != nil {
// 			b.Fatal(err)
// 		}

// 		createMovieHandlerDecoder(w, r)
// 	}
// }
