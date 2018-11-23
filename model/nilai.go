package model

import (
	"CRUD-Golang/lib"
	"database/sql"
	"fmt"
)

type Nilai struct {
	IdNilai int    `json:"id"`
	Kd_mk   string `json:"kd_mk"`
	NPM     string `json:"npm"`
	UAS     int    `json:"uas"`
	UTS     int    `json:"uts"`
	Total   int    `json:"total"`
	Ipk     string `json:"ipk"`
}

var TbNilai string = `
CREATE TABLE nilai (
	idNilai SERIAL PRIMARY KEY NOT NULL,
	kd_mk VARCHAR(10) NOT NULL,
	NPM VARCHAR(10) NOT NULL,
	UAS INT NOT NULL,
	UTS INT NOT NULL,
	total INT NOT NULL,
	ipk VARCHAR(5) NOT NULL
);
`

func (n *Nilai) Name() string {
	return "nilai"
}

func (n *Nilai) Field() ([]string, []interface{}) {
	fields := []string{"IdNilai", "kd_mk", "NPM", "UAS", "UTS", "total", "ipk"}
	n.Total = int((n.UTS + n.UAS) / 2)
	switch {
	case n.Total > 80:
		n.Ipk = "A"
	case n.Total > 70:
		n.Ipk = "B"
	case n.Total > 60:
		n.Ipk = "C"
	case n.Total > 50:
		n.Ipk = "D"
	default:
		n.Ipk = "E"
	}
	dst := []interface{}{&n.IdNilai, &n.Kd_mk, &n.NPM, &n.UAS, &n.UTS, &n.Total, &n.Ipk}
	return fields, dst
}

func (n *Nilai) PrimaryKey() ([]string, []interface{}) {
	fields := []string{"IdNilai"}
	dst := []interface{}{&n.IdNilai}
	return fields, dst

}

func (n *Nilai) Structur() lib.Table {
	return &Nilai{}
}

func (n *Nilai) AutoNumber() bool {
	return true
}

func (n *Nilai) Insert(db *sql.DB) error {

	return lib.Insert(db, n)
}

func (n *Nilai) Delete(db *sql.DB) error {
	return lib.Delete(db, n)
}

func (n *Nilai) Get(db *sql.DB) error {
	return lib.Get(db, n)
}

func (n *Nilai) Update(db *sql.DB, data map[string]interface{}) error {
	_, utsOK := data["uts"]
	_, uasOK := data["uas"]
	if utsOK || uasOK {
		if err := n.Get(db); err != nil {
			return err
		}
		uts := int(n.UTS)
		uas := int(n.UAS)
		if utsOK {
			uts = data["uts"].(int)
		}
		if uasOK {
			uas = data["uas"].(int)
		}
		nilai := int((float64(uts) + float64(uas)) / 2)
		var grade string
		switch {
		case nilai > 80:
			grade = "A"
		case nilai > 70:
			grade = "B"
		case nilai > 60:
			grade = "C"
		case nilai > 50:
			grade = "D"
		default:
			grade = "E"
		}
		data["ipk"] = grade
		data["total"] = nilai

	}
	fmt.Println(data["ipk"])
	return lib.Update(db, n, data)
}

func GetAllNilai(db *sql.DB, params ...string) ([]*Nilai, error) {
	m := &Nilai{}
	data, err := lib.Gets(db, m, params...)
	if err != nil {
		return nil, err
	}
	nilai := make([]*Nilai, len(data))
	for index, item := range data {
		nilai[index] = item.(*Nilai)
	}
	return nilai, nil

}
