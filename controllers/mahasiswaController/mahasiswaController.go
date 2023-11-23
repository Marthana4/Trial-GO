package mahasiswacontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"fmt"
	"github.com/Marthana4/Trial-GO/entities"
	"github.com/Marthana4/Trial-GO/models"
	// "github.com/go-playground/validator/v10/translations/id"
)

var mahasiswaModel = models.NewMahasiswaModel()

func Index(response http.ResponseWriter, request *http.Request){
	mahasiswa, _ := mahasiswaModel.FindAll()

	data := map[string]interface{}{
		"mahasiswa" : mahasiswa,
	}

	temp, err := template.ParseFiles("views/mahasiswa/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request){
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/mahasiswa/add.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(response, nil)

	} else if request.Method == http.MethodPost {
		request.ParseForm()
	
		var mahasiswa entities.Mahasiswa
		mahasiswa.NIM = request.Form.Get("nim")
		mahasiswa.Nama = request.Form.Get("nama")
		mahasiswa.Jurusan = request.Form.Get("jurusan")
	
		isSuccess := mahasiswaModel.Create(mahasiswa)
		data := map[string]interface{}{
			"pesan": "Behasil! Data Mahasiswa berhasil ditambahkan!",
		}
	
		handleCreateError(isSuccess, response, data)
		temp, _ := template.ParseFiles("views/mahasiswa/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request){
	if request.Method == http.MethodGet {
        queryString := request.URL.Query()
        id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)
        if err != nil {
            http.Error(response, "Invalid ID format", http.StatusBadRequest)
            return
        }

        var mahasiswa entities.Mahasiswa
        err = mahasiswaModel.Find(id, &mahasiswa)
        if err != nil {
            http.Error(response, "Data Mahasiswa tidak ditemukan", http.StatusNotFound)
            return
        }

        data := map[string]interface{}{
            "mahasiswa": mahasiswa,
        }

        temp, err := template.ParseFiles("views/mahasiswa/edit.html")
        if err != nil {
            panic(err)
        }

        temp.Execute(response, data)

	} else if request.Method == http.MethodPost {
        request.ParseForm()

        var mahasiswa entities.Mahasiswa
        mahasiswa.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
        mahasiswa.NIM = request.Form.Get("nim")
        mahasiswa.Nama = request.Form.Get("nama")
        mahasiswa.Jurusan = request.Form.Get("jurusan")

        err := mahasiswaModel.Update(mahasiswa)
        data := map[string]interface{}{
            "pesan": "Behasil! Data Mahasiswa berhasil diedit!",
        }

        if err != nil {
            fmt.Println("Error:", err)
            data["pesan"] = "Gagal! Terjadi kesalahan saat mengedit data Mahasiswa."
        }

        temp, _ := template.ParseFiles("views/mahasiswa/edit.html")
        temp.Execute(response, data)
    }
}

func Delete(response http.ResponseWriter, request *http.Request){
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10,64)

	mahasiswaModel.Delete(id)
	http.Redirect(response, request, "/mahasiswa/index", http.StatusSeeOther)
}

func handleCreateError(isSuccess bool, response http.ResponseWriter, data map[string]interface{}) {
    if !isSuccess {
        data["pesan"] = "Gagal! Terjadi kesalahan saat menambahkan data Mahasiswa."
    }
}