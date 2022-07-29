package models

import (
	//"database/sql"
	//"fmt"
	//"log"

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


