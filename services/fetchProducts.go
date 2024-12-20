package services

import "net/http"

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoaXZhbW1Ab3AuY29tIiwiZXhwIjoxNzM0NzYwMzc4LCJpc3N1ZWRfYXQiOjE3MzQ2NzM5NzgsInVzZXJfaWQiOjExfQ.7Rbn2reHwF5_Z7kuVdw-e5RkGArxDRFk59AP-QfL-S8"

func FetchProducts(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Authorization", "Bearer"+token)
	w.Header().Set("Content-Type", "application/json")

}
