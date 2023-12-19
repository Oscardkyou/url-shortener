package handler

import 

func (us *URLShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]
	if shortKey == "" {
	http.Error(w, "Shortened key is missing", http.StatusBadRequest)
	return
	}
