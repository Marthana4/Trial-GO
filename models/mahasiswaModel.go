package models

import (
	"database/sql"
	"fmt"

	"github.com/Marthana4/Trial-GO/config"
	"github.com/Marthana4/Trial-GO/entities"
)

type MahasiswaModel struct{
	conn *sql.DB
}

func NewMahasiswaModel() *MahasiswaModel{
	conn, err := config.DBconnection()
	if err != nil {
		panic(err)
	}

	return &MahasiswaModel{
		conn: conn,
	}
}

func (p *MahasiswaModel) FindAll()([]entities.Mahasiswa, error) {
	rows, err := p.conn.Query("select * from mahasiswa")
	
	if err != nil {
		return []entities.Mahasiswa{}, err
	}
	defer rows.Close()

	var dataMahasiswa []entities.Mahasiswa
	for rows.Next(){
		var mahasiswa entities.Mahasiswa
		rows.Scan(
			&mahasiswa.Id,
			&mahasiswa.NIM,
			&mahasiswa.Nama,
			&mahasiswa.Jurusan,
		)		

		dataMahasiswa = append(dataMahasiswa, mahasiswa)

	}

	return dataMahasiswa, nil
}

func (p *MahasiswaModel) Create(mahasiswa entities.Mahasiswa) bool {
	result, err := p.conn.Exec("insert into Mahasiswa (nim, nama, jurusan) values(?,?,?)",
				   mahasiswa.NIM, mahasiswa.Nama, mahasiswa.Jurusan)
	
	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
}

func (p *MahasiswaModel) Find(id int64, mahasiswa *entities.Mahasiswa)error {
	return p.conn.QueryRow("select * from mahasiswa where id = ?", id).Scan(
		&mahasiswa.Id,
		&mahasiswa.NIM,
		&mahasiswa.Nama,
		&mahasiswa.Jurusan,
	)
}

func (p *MahasiswaModel) Update(mahasiswa entities.Mahasiswa) error {
	_, err := p.conn.Exec("UPDATE mahasiswa set nim = ? , nama = ? , jurusan = ? WHERE id = ?",
				   mahasiswa.NIM, mahasiswa.Nama, mahasiswa.Jurusan, mahasiswa.Id)
	
	if err != nil {
		return err
	}

	return nil
}

func (p *MahasiswaModel) Delete(id int64) {
    p.conn.Exec("DELETE FROM mahasiswa WHERE id = ?", id)
}
