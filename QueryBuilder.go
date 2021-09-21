package Monoql

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)
type Monoql struct {
Client *mongo.Client
DB *mongo.Database
Coll *mongo.Collection
query Query
}

type Query struct {
	DBName string
	CollectionName string
}

func (m Monoql) Connect(options *options.ClientOptions) (*Monoql,error){
	client, err := mongo.Connect(m.NewCTX(), options)
	m.Client = client
	return &m,err
}
func (m Monoql) Database (name string) *Monoql {
	m.query.DBName = name
	m.DB = m.Client.Database(name)
	return &m
}
func (m Monoql) Collection(name string) *Monoql{
	m.query.CollectionName = name
	m.Coll = m.DB.Collection(name)
	return &m
}
func (m Monoql) FindOne (filter interface{}) *mongo.SingleResult {
	return m.Coll.FindOne(m.NewCTX(),filter)
}
func (m Monoql) Find (filter interface{}) (*mongo.Cursor, error){
	return m.Coll.Find(m.NewCTX(),filter)
}
func (m Monoql) InsertOne(document interface{}) (*mongo.InsertOneResult, error){
	return m.Coll.InsertOne(m.NewCTX(),document)
}

func (m Monoql) InsertMany(document []interface{}) (*mongo.InsertManyResult, error){
	return m.Coll.InsertMany(m.NewCTX(),document)
}
func (m Monoql) UpdateOne(filter interface{},document interface{})(*mongo.UpdateResult, error) {
	return m.Coll.UpdateOne(m.NewCTX(),filter,document)
}
func (m Monoql) UpdateByID(id interface{},document interface{})(*mongo.UpdateResult, error) {
	return m.Coll.UpdateByID(m.NewCTX(),id,document)
}
func (m Monoql) UpdateMany(filter interface{},document []interface{})(*mongo.UpdateResult, error) {
	return m.Coll.UpdateMany(m.NewCTX(),filter,document)
}
func (m Monoql) NewCTX() context.Context{
	ctx,_:= context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func (m Monoql) ConvertObjectIDToString(InsertedID interface{}) (string,bool){
	if oid, ok := InsertedID.(primitive.ObjectID); ok {
		return oid.String(),true
	}
	return "",false
}