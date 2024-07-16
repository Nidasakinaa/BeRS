package bers_test

import (
	"fmt"
	"testing"

	model "github.com/Nidasakinaa/BeRS/model"
	module "github.com/Nidasakinaa/BeRS/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetPasienByID(t *testing.T) {
	_id := "669534d2af52bee3d2606c34"
	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPasienByID(objectID, module.MongoConn, "DataPasien")
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
	pasienName := "Sari Purnamasari"
	gender := "Perempuan"
	ttl := "Semarang, 7 November 1962"
	phonenumber := "0822334455"
	alamat := "Jl.Diponegoro, Kota Jogja 54321"
	doctor := model.Doctor{
		Name:      "Dina",
		Specialty: "Gynecology",
		Contact:   "222-3333-4444",
	}
	medicalRecord := model.MedicalRecord{
		VisitDate: "12 Juli 2023",
		Diagnosis: "Pregnancy",
		Treatment: "Prenatal vitamins",
	}
	insertedID, err := module.InsertPasien(module.MongoConn, "DataPasien", pasienName, gender, ttl, phonenumber, alamat, doctor, medicalRecord)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestDeletePasienByID(t *testing.T) {
	id := "668e2b1540bdb1d47710a316"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePasienByID(objectID, module.MongoConn, "DataPasien")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	_, err = module.GetPasienByID(objectID, module.MongoConn, "DataPasien")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
