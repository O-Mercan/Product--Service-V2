package models

import (
	"database/sql"
	"fmt"
	"log"

	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"description"`
	Category    string    `json:"category"`
	Summary     string    `json:"summary"`
	Description string    `json:"desription"`
	Price       int       `json:"price"`
	CreatedOn   time.Time `json:"createdOn"`
	ChangedOn   time.Time `json:"changedOn"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sql1234"
	database = "productsdb"
)

var db *sql.DB

func init() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable", host, port, user, password, database)

	var err error

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
}

//ADD A PRODUCT
func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(name, category, summary, description,price ) VALUES($1,$2,$3,$4,$5)", data.Name, data.Category, data.Summary, data.Description, data.Price)
	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Number of affected rows(%d", rowsAffected)
}

// UPDATE A PRODUCT
func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE products SET name=$2, category=$3, summary=$4, description=$5, price=$6 WHERE id=$1", data.Id, data.Name, data.Category, data.Summary, data.Description, data.Price)
	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Number of affected rows(%d", rowsAffected)
}

//GET ALL PRODUCTS
func GetProducts() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No records found")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		produc := &Product{}
		err := rows.Scan(&produc.Id, &produc.Name, &produc.Category, &produc.Summary, &produc.Description, &produc.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, produc)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	for _, produc := range products {
		fmt.Printf("%d,%s,%s,%s,%s,%d\n", produc.Id, produc.Name, produc.Category, produc.Summary, produc.Description, produc.Price)
	}
}

//GET A PRODUCT BY ID
func GetProductByID(id int) {
	var product string
	err := db.QueryRow("SELECT name FROM products WHERE id =$1", id).Scan(&product)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product is %s\n", product)
	}

}

//DELETE A PRODUCT BY ID
func DeleteProductByID(id int) {

	result, err := db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Number of affected rows(%d", rowsAffected)
	if rowsAffected == 0 {
		fmt.Println("An error has occured while trying to delete the product")
	}

}
