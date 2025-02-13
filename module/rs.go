package module

import (
	"context"
	"errors"
	"fmt"
	"time"

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

func InsertPasien(db *mongo.Database, col string, pasienName string, gender string, usia string, phonenumber string, alamat string, doctor model.Doctor, medicalRecord model.MedicalRecord) (insertedID primitive.ObjectID, err error) {
	pasien := bson.M{
		"pasienName":    pasienName,
		"gender":        gender,
		"usia":          usia,
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

func UpdatePasien(ctx context.Context, db *mongo.Database, col string, _id primitive.ObjectID, pasienName string, gender string, usia string, phonenumber string, alamat string, doctor model.Doctor, medicalRecord model.MedicalRecord) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"pasienName":    pasienName,
			"gender":        gender,
			"usia":          usia,
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

// FUNCTION USER
// GetUserByID retrieves a user from the database by its ID
func GetUserByID(_id primitive.ObjectID, db *mongo.Database, col string) (model.User, error) {
	var user model.User
	collection := db.Collection("User")
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, fmt.Errorf("GetUserByID: user dengan ID %s tidak ditemukan", _id.Hex())
		}
		return user, fmt.Errorf("GetUserByID: gagal mendapatkan data user: %w", err)
	}
	return user, nil
}

func InsertUsers(db *mongo.Database, col string, fullname string, phonenumber string, username string, password string, role string) (insertedID primitive.ObjectID, err error) {
	users := bson.M{
		"fullname": fullname,
		"phone":    phonenumber,
		"username": username,
		"password": password,
		"role":     role,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), users)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetByUsername(db *mongo.Database, col string, username string) (*model.User, error) {
	var admin model.User
	err := db.Collection(col).FindOne(context.Background(), bson.M{"username": username}).Decode(&admin)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func DeleteTokenFromMongoDB(db *mongo.Database, col string, token string) error {
	collection := db.Collection(col)
	filter := bson.M{"token": token}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUser retrieves all users from the database
func GetAllUser(db *mongo.Database, col string) ([]model.User, error) {
	var data []model.User
	user := db.Collection(col)

	cursor, err := user.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("GetAllUser error:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO()) // Selalu tutup cursor

	if err := cursor.All(context.TODO(), &data); err != nil {
		fmt.Println("Error decoding users:", err)
		return nil, err
	}

	return data, nil
}

func SaveTokenToDatabase(db *mongo.Database, col string, adminID string, token string) error {
	collection := db.Collection(col)
	filter := bson.M{"admin_id": adminID}
	update := bson.M{
		"$set": bson.M{
			"token":      token,
			"updated_at": time.Now(),
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(ctx context.Context, db *mongo.Database, col string, _id primitive.ObjectID, name string, phone string, username string, password string, role string) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"name":     name,
			"phone":    phone,
			"username": username,
			"password": password,
			"role":     role,
		},
	}
	result, err := db.Collection(col).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateUser: gagal memperbarui User: %w", err)
	}
	if result.MatchedCount == 0 {
		return errors.New("UpdateUser: tidak ada data yang diubah dengan ID yang ditentukan")
	}
	return nil
}

// DeleteUserByID deletes a menu item from the database by its ID
func DeleteUserByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	user := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}