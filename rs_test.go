package bers_test

import (
	"fmt"
	"testing"

	model "github.com/Nidasakinaa/BeRS/model"
	module "github.com/Nidasakinaa/BeRS/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetPasienByID(t *testing.T) {
	_id := "668d1f52ff875d6b3b3ed63e"
	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPasienByID(objectID, module.MongoConn, "pasien")
	if err != nil {
		t.Fatalf("error calling GetPasienFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllPasien(module.MongoConn, "DataPasien")
	fmt.Println(data)
}

func TestInsertPasien(t *testing.T) {
	pasienName := "Naya Kania"
	gender := "Perempuan"
	ttl := "Bandung, 10 Mei 2019"
	status := "Belum Menikah"
	phonenumber := "13456"
	alamat := "Jl.Bandung Kota Maluku 71903"
	doctor := model.Doctor{
		Name:      "Gita",
		Specialty: "Oncology",
		Contact:   "123-456-7890",
	}
	medicalRecord := model.MedicalRecord{
		VisitDate:  "12 Juli 2023",
		DoctorName: "Andre",
		Diagnosis:  "Cancer",
		Treatment:  "Kemo",
		Notes:      "-",
	}
	insertedID, err := module.InsertPasien(module.MongoConn, "DataPasien", pasienName, gender, ttl, status, phonenumber, alamat, doctor, medicalRecord)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestDeletePresensiByID(t *testing.T) {
	id := "668d22c2ecd9a334601ece41" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePasienByID(objectID, module.MongoConn, "DataPasien")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetPasienByID(objectID, module.MongoConn, "DataPasien")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
