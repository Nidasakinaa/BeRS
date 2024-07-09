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
	collection := db.Collection("DataPasien")
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

func UpdatePasien(ctx context.Context, db *mongo.Database, col string, _id primitive.ObjectID, pasienName string, gender string, ttl string, status string, phonenumber string, alamat string, doctor model.Doctor, medicalRecord model.MedicalRecord) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"pasienName":    pasienName,
			"gender":        gender,
			"ttl":           ttl,
			"status":        status,
			"phonenumber":   phonenumber,
			"alamat":        alamat,
			"doctor":        doctor,
			"medicalRecord": medicalRecord,
		},
	}
	result, err := db.Collection(col).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdatePasien: gagal memperbarui pasien: %w", err)
	}
	if result.MatchedCount == 0 {
		return errors.New("UpdatePasien: tidak ada data yang diubah dengan ID yang ditentukan")
	}
	return nil
}

func DeletePasienByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	pasien := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := pasien.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
