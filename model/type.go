package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct {
	Name      string `bson:"name,omitempty" json:"name,omitempty"`
	Specialty string `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Contact   string `bson:"contact,omitempty" json:"contact,omitempty"`
}

type MedicalRecord struct {
	ID        primitive.ObjectID `bson:"m_id,omitempty" json:"m_id,omitempty"`
	VisitDate string             `bson:"visitdate,omitempty" json:"visitdate,omitempty"`
	Diagnosis string             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Treatment string             `bson:"treatment,omitempty" json:"treatment,omitempty"`
}

type Biodata struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	PasienName    string             `bson:"pasienName,omitempty" json:"pasienName,omitempty"`
	Gender        string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Usia          string             `bson:"usia,omitempty" json:"usia,omitempty"`
	Phonenumber   string             `bson:"phonenumber,omitempty" json:"phonenumber,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Doctor        Doctor             `bson:"doctor,omitempty" json:"doctor,omitempty"`
	MedicalRecord MedicalRecord      `bson:"medicalRecord,omitempty" json:"medicalRecord,omitempty"`
}

type User struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FullName               string             `bson:"name,omitempty" json:"name,omitempty"`
	Phone                  string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Username               string             `bson:"username,omitempty" json:"username,omitempty"`
	Password               string             `bson:"password,omitempty" json:"password,omitempty"`
}

type Token struct {
	ID        string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Token     string    `bson:"token" json:"token,omitempty"`
	// AdminID   string    `bson:"admin_id" json:"admin_id,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}