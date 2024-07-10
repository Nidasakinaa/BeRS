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
	pasienName := "Mira"
	gender := "Perempuan"
	ttl := "Semarang, 5 November 1992"
	status := "Menikah"
	phonenumber := "0822334455"
	alamat := "Jl.Diponegoro, Kota Semarang 54321"
	doctor := model.Doctor{
		Name:      "Dina",
		Specialty: "Gynecology",
		Contact:   "222-3333-4444",
	}
	medicalRecord := model.MedicalRecord{
		VisitDate:  "12 Juli 2023",
		DoctorName: "Eko",
		Diagnosis:  "Pregnancy",
		Treatment:  "Prenatal vitamins",
		Notes:      "Regular check-ups needed",
	}
	insertedID, err := module.InsertPasien(module.MongoConn, "DataPasien", pasienName, gender, ttl, status, phonenumber, alamat, doctor, medicalRecord)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestDeletePasienByID(t *testing.T) {
	id := "668e2b1540bdb1d47710a316" // ID data yang ingin dihapus
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
