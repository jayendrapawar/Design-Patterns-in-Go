package main

import "fmt"

// Strategy Pattern / Policy Pattern
// What ?
// Why ?
// How ? (Interfaces, structs and Reciver Functions)
// Pros and Cons

// interfaces can accepts anything - method, value, type
// only constraint it accepts one value at a time
type IDBconnection interface{
	Connect() // accepts anything
}

type DBConnection struct{
	Db IDBconnection // Compatible to accept any type in runtime
}

func (con DBConnection) DBConnect() { // Reciver function for Struct DBConnection
	con.Db.Connect()
}

// Mysql
type MysqlConnection struct{
	ConnectionString string
}

func (con MysqlConnection) Connect() {
	fmt.Println(("Mysql :" + con.ConnectionString))
}


// Postgres
type PostgresConnection struct{
	ConnectionString string
}

func (con PostgresConnection) Connect() {
	fmt.Println(("Postgres :" + con.ConnectionString))
}


// MongoDB
type MongoDBConnection struct{
	ConnectionString string
}

func (con MongoDBConnection) Connect() {
	fmt.Println(("MongoDB :" + con.ConnectionString))
}



func main(){
	MysqlConnection := MysqlConnection{ConnectionString: "Mysql DB is Conncected "}
	con1 := DBConnection{Db: MysqlConnection}
	con1.DBConnect()

	PostgresConnection := PostgresConnection{ConnectionString: "Postgres DB is Conncected "}
	con2 := DBConnection{Db: PostgresConnection}
	con2.DBConnect()

	MongoDBConnection := MongoDBConnection{ConnectionString: "Mongo DB is Conncected "}
	con3 := DBConnection{Db: MongoDBConnection}
	con3.DBConnect()

}