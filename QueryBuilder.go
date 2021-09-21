package monoql

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)
type monoql struct {
Client *mongo.Client
DB *mongo.Database
Coll *mongo.Collection
ctx context.Context
query Query
}

type Query struct {
	DBName string
	CollectionName string
}
func (m monoql) Connect(options *options.ClientOptions) *monoql{
	m.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(m.ctx, options)
	if err != nil{
		fmt.Printf("%v \n",err.Error())
	}
	m.Client = client
	return &m
}
func (m monoql) Database (name string) *monoql {
	m.query.DBName = name
	m.DB = m.Client.Database(name)
	return &m
}
func (m monoql) Collection(name string) *monoql{
	m.query.CollectionName = name
	m.Coll = m.DB.Collection(name)
	return &m
}
func (m monoql) FindOne (filter interface{}) *mongo.SingleResult {
	return m.Coll.FindOne(m.ctx,filter)
}
func (m monoql) Find (filter interface{}) (*mongo.Cursor, error){
	return m.Coll.Find(m.ctx,filter)
}
func (m monoql) InsertOne(document interface{}) (*mongo.InsertOneResult, error){
	return m.Coll.InsertOne(m.ctx,document)
}

func (m monoql) InsertMany(document []interface{}) (*mongo.InsertManyResult, error){
	return m.Coll.InsertMany(m.ctx,document)
}
func (m monoql) UpdateOne(filter interface{},document interface{})(*mongo.UpdateResult, error) {
	return m.Coll.UpdateOne(m.ctx,filter,document)
}
func (m monoql) UpdateByID(id interface{},document interface{})(*mongo.UpdateResult, error) {
	return m.Coll.UpdateByID(m.ctx,id,document)
}
func (m monoql) UpdateMany(filter interface{},document []interface{})(*mongo.UpdateResult, error) {
	return m.Coll.UpdateMany(m.ctx,filter,document)
}