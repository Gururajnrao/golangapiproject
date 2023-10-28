package main

import (
	"database/sql"
	"fmt"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]product, error) {
	query := "SELECT id , name, quantity, price FROM products"
	rows, err := db.Query(query)
	// fmt.Println("Rows:", rows)

	if err != nil {
		return nil, err
	}
	products := []product{}
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil

}

func (p *product) getProduct(db *sql.DB) error {
	query := fmt.Sprintf("SELECT id , name, quantity, price FROM products where id=%v", p.ID)
	rows := db.QueryRow(query)
	err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
	if err != nil {
		return err
	}
	return nil

}

func (p *product) createProduct(db *sql.DB) error {
	query := fmt.Sprintf("insert into products(name,quantity,price) values ('%v', %v, %v) ", p.Name, p.Quantity, p.Price)
	result, err := db.Exec(query)
	fmt.Println("Result, err", result, err)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("LastinsertID:", id)
	p.ID = int(id)
	return nil

}

func (p *product) updateProduct(db *sql.DB) error {
	query := fmt.Sprintf("update products set quantity=%v, name='%v', price=%v where id=%v", p.Quantity, p.Name, p.Price, p.ID)
	result, err := db.Exec(query)
	fmt.Println("Result, err", result, err)
	if err != nil {
		return err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("RowsAffected:", id)
	//p.ID = int(id)
	return nil

}

func (p *product) deleteProduct(db *sql.DB) error {
	query := fmt.Sprintf("delete from products where id=%v", p.ID)
	result, err := db.Exec(query)
	fmt.Println("Result, err", result, err)
	if err != nil {
		return err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("RowsAffected:", id)
	//p.ID = int(id)
	return nil

}
