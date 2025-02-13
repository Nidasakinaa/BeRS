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
	pasienName := "Sari Endah"
	gender := "Perempuan"
	usia := "22 Tahun"
	phonenumber := "0822334455"
	alamat := "Jl.Diponegoro, Kota Jogja 54321"
	doctor := model.Doctor{
		Name:      "Dina",
		Specialty: "Gynecology",
		Contact:   "08217456",
	}
	medicalRecord := model.MedicalRecord{
		VisitDate: "12 Juli 2023",
		Diagnosis: "Pregnancy",
		Treatment: "Prenatal vitamins",
	}
	insertedID, err := module.InsertPasien(module.MongoConn, "DataPasien", pasienName, gender, usia, phonenumber, alamat, doctor, medicalRecord)
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

//FUNCTION USER
//GetUserByID retrieves a user from the database by its ID
func TestGetUserByID(t *testing.T) {
	_id := "67a23885d8d58983179fe315"
	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	menu, err := module.GetUserByID(objectID, module.MongoConn, "User")
	if err != nil {
		t.Fatalf("error calling GetMenuItemByID: %v", err)
	}
	fmt.Println(menu)
}

func TestGetAllUsers(t *testing.T) {
	data, err := module.GetAllUser(module.MongoConn, "User")
	if err != nil {
		t.Fatalf("error calling GetAllUsers: %v", err)
	}
	fmt.Println(data)
}	

func TestInsertUser(t *testing.T) {
	name := "Nida Sakina"
    phone := "083174603834"
    username := "nida"
    password := "Nida150304"
    role := "admin"
    insertedID, err := module.InsertUsers(module.MongoConn, "User", name, phone, username, password, role)
    if err != nil {
        t.Errorf("Error inserting data: %v", err)
    }
    fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestDeleteUserByID(t *testing.T) {
    id := "67a23885d8d58983179fe315"
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        t.Fatalf("error converting id to ObjectID: %v", err)
    }

    err = module.DeleteUserByID(objectID, module.MongoConn, "User")
    if err != nil {
        t.Fatalf("error calling DeleteUserByID: %v", err)
    }

    _, err = module.GetUserByID(objectID, module.MongoConn, "User")
    if err == nil {
        t.Fatalf("expected data to be deleted, but it still exists")
    }
}