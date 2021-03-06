package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ToDo struct{
	Kegiatan string `json:"kegiatan"`
	Waktu string `json:"waktu"`
}

type JSONResponse struct{
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string `json:"message"`
	// Data []ToDo `json:"data"`
	Data interface{} `json:"data"`
}

func main() { 

	daftarKegiatan := []ToDo{}
	daftarKegiatan = append(daftarKegiatan, ToDo{"Liburan Ke Bali", "2021-11-30"})
	daftarKegiatan = append(daftarKegiatan, ToDo{"Liburan Ke Karimunjawa", "2021-12-25"})
	
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// GET http://hocalhost:8080/
		if r.Method == "GET" {
			// method GET
			rw.Header().Add("Content-Type", "application/json")
			// res := JSONResponse{
			// 	http.StatusOK,
			// 	true,
			// 	"Uji coba Get Method pada postman",
			// 	[]ToDo{},
			// }
			// resJSON, err := json.Marshal(res)
			// if err != nil {
			// 	http.Error(rw, "Terjadi Kesalahan", http.StatusInternalServerError)
			// }
			// rw.Write(resJSON)

			res := JSONResponse{
				http.StatusOK,
				true,
				"Berhasil mendapatkan daftar aktifitas",
				daftarKegiatan,
			}
			resJSON, err := json.Marshal(res)
			if err != nil {
				fmt.Println("Terjadi Kesalahan")
				http.Error(rw, "Terjadi Kesalahan", http.StatusInternalServerError)
				return
			}
			rw.Write(resJSON)
			return

		} else if r.Method == "POST"{
			//method POST

			//menambahkan data baru

			// isi r.body(JSON)
			// {
			// 	"kegiatan": "main futsal",
			// 	"waktu": "2021-12-5"
			// }

			jsonDecode := json.NewDecoder(r.Body)
			aktifitasBaru := ToDo{}
			res := JSONResponse{}

			if err := jsonDecode.Decode(&aktifitasBaru); err != nil {
				fmt.Println("Terjadi Kesalahan")
				http.Error(rw, "Terjadi Kesalahan saat membaca input", http.StatusInternalServerError)
				return
			}

			res.Code = http.StatusCreated
			res.Success = true
			res.Message = "Berhasil menambahkan data"
			res.Data = aktifitasBaru

			resJSON, err := json.Marshal(res)
			if err != nil {
				fmt.Println("Terjadi Kesalahan")
				http.Error(rw, "Terjadi Kesalahan saat ubah json", http.StatusInternalServerError)
				return
			}
			rw.Header().Add("Content-Type", "application/json")
			rw.Write(resJSON)
			return

		}
	})
	
	//membuat lokal server
	fmt.Println("Lstening on: 8080 ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}