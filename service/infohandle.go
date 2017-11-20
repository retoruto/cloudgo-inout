package service

import (
	"net/http"

	"github.com/unrolled/render"
)
func InfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		var sex string

		if req.Form["sex"][0] == "1" {
			sex = "Male"
		} else if req.Form["sex"][0] == "2" {
			sex = "Female"
		} else if req.Form["sex"][0] == "3" {
                        sex = "Secret"
                }
                var mem string

                if req.Form["mem"][0] == "1" {
			mem = "Ohno Satoshi"
		} else if req.Form["mem"][0] == "2" {
			mem = "Sakurai Sho"
		} else if req.Form["mem"][0] == "3" {
                        mem = "Aiba Masaki"
                } else if req.Form["mem"][0] == "4" {
                        mem = "Ninomiya Kazunari"
                } else if req.Form["mem"][0] == "5" {
                        mem = "Matumoto Jun"
                }

		formatter.HTML(w, http.StatusOK, "table", struct {
			Name   string
			Sex    string
                        Mem    string
                        Addr   string
			Phone string
			Email string
		}{Name: req.Form["name"][0], Sex: sex, Mem: mem, Addr: req.Form["addr"][0], Phone: req.Form["phone"][0], Email: req.Form["email"][0]})
	}
}
