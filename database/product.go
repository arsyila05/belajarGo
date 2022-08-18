package database

type Product struct {
	Product_ID   int    `json:"product_id" gotm:"primary_key"`
	Product_Name string `json:"product_name"`
	Description  string `json:"description"`
}
