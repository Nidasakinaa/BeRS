package module

import (
	"context"
	"errors"
	"fmt"

	model "github.com/Nidasakinaa/BeRS/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetPasienByID(_id primitive.ObjectID, db *mongo.Database, col string) (model.Biodata, error) {
	var pasien model.Biodata
	collection := db.Collection("pasien")
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&pasien)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return pasien, fmt.Errorf("GetPasienByID: pasien dengan ID %s tidak ditemukan", _id.Hex())
		}
		return pasien, fmt.Errorf("GetPasienByID: gagal mendapatkan pasien: %w", err)
	}
	return pasien, nil
}

func GetAllPasien(db *mongo.Database, col string) (data []model.Biodata) {
	pasien := db.Collection(col)
	filter := bson.M{}
	cursor, err := pasien.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPasien :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertPasien(db *mongo.Database, col string, pasienName string, gender string, ttl string, status string, phonenumber string, alamat string, doctor model.Doctor, medicalRecord model.MedicalRecord) (insertedID primitive.ObjectID, err error) {
	pasien := bson.M{
		"pasienName":    pasienName,
		"gender":        gender,
		"ttl":           ttl,
		"status":        status,
		"phonenumber":   phonenumber,
		"alamat":        alamat,
		"doctor":        doctor,
		"medicalRecord": medicalRecord,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), pasien)
	if err != nil {
		fmt.Printf("InsertPasien: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}
