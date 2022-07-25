package models

import "time"

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"description"`
	Category    string    `json:"category"`
	Summary     string    `json:summary`
	Description string    `json:description`
	Price       int       `json:price`
	CreatedOn   time.Time `json:createdon`
	ChangedOn   time.Time `json:changedon`
}
