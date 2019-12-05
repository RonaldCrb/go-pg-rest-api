package offers

import (
	"errors"
	"log"
	"time"

	"github.com/RonaldCrb/go-pg-rest-api/config"
)

// Offer represents a post from localbtc
type Offer struct {
	ID         int
	Title      string
	Trader     string
	Reputation int
	Price      float32
	Min        float32
	Max        float32
	Index      int
	Bank       string
	Currency   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// CreateOffer creates a new Offer instance
func (o Offer) CreateOffer() error {
	ofr := `INSERT INTO 
						btcoffers (Title,
											 Trader,
											 Reputation,
											 Price,
											 Min,
											 Max,
											 Index,
											 Bank,
											 Currency
											)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	stmt, err := config.DB.Prepare(ofr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(o.Title, o.Trader, o.Reputation, o.Price, o.Min, o.Max, o.Index, o.Bank, o.Currency)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := rows.RowsAffected()

	if aff != 1 {
		err = errors.New("[ERROR - OFFERS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - OFFERS - MODEL] => %v", err)
		return err
	}

	return nil
}

// AllOffers returns a slice of Offer (all Offers in Offers table)
func AllOffers() ([]Offer, error) {
	q := "SELECT * FROM btcoffers"

	rows, err := config.DB.Query(q)

	if err != nil {
		log.Printf("[ERROR - OFFERS - MODEL] => %v", err)
		return nil, err
	}
	defer rows.Close()

	ofrs := make([]Offer, 0)

	for rows.Next() {
		ofr := Offer{}
		err := rows.Scan(&ofr.ID, &ofr.Title, &ofr.Trader, &ofr.Bank, &ofr.Currency, &ofr.Reputation, &ofr.Price, &ofr.Min, &ofr.Max, &ofr.Index, &ofr.CreatedAt, &ofr.UpdatedAt)

		if err != nil {
			log.Printf("[ERROR - OFFERS - MODEL] => %v", err)
			return nil, err
		}

		ofrs = append(ofrs, ofr)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[ERROR - OFFERS - MODEL] => %v", err)
		return nil, err
	}
	return ofrs, nil
}

// FindOffer returns a Offer instance from the database
func (o Offer) FindOffer() (Offer, error) {
	q := "SELECT * FROM btcoffers WHERE id = $1"

	row := config.DB.QueryRow(q, o.ID)

	ofr := Offer{}

	err := row.Scan(&ofr.ID, &ofr.Title, &ofr.Trader, &ofr.Bank, &ofr.Currency, &ofr.Reputation, &ofr.Price, &ofr.Min, &ofr.Max, &ofr.Index, &ofr.CreatedAt, &ofr.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR - OFFERS - MODEL] => %v", err)
		return Offer{}, err
	}

	return ofr, nil
}

// UpdateOffer updates the data for a Offer instance in the database
func (o Offer) UpdateOffer() error {
	q := "UPDATE btcoffers SET Title=$1, Trader=$2, Reputation=$3, Price=$4, Min=$5, Max=$6, Index=$7, Bank=$8, Currency=$9, UpdatedAt=now() WHERE ID = $10"

	if o.Index <= 0 || o.Trader == "" || o.Title == "" || o.Currency == "" || o.Bank == "" {
		err := errors.New("[ERROR - OFFERS - MODEL] => Index, Trader, Title, Currency and Bank fields are Mandatory")
		return err
	}

	stmt, err := config.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(&o.Title, &o.Trader, &o.Reputation, &o.Price, &o.Min, &o.Max, &o.Index, &o.Bank, &o.Currency, &o.ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := row.RowsAffected()
	if aff != 1 {
		err = errors.New("[ERROR - OFFERS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - OFFERS - MODEL] => %v", err)
		return err
	}
	return nil
}

// DeleteOffer deletes a Offer instance from the database
func (o Offer) DeleteOffer() error {
	q := `DELETE FROM btcoffers WHERE ID=$1`

	_, err := config.DB.Exec(q, o.ID)
	if err != nil {
		log.Printf("[ERROR - OFFERS - MODEL] => %v", err)
		return err
	}

	return nil
}

func CreateOffersTable() error {

	// create required tables
	offersTable := `
		CREATE TABLE btcoffers (
		  ID 					SERIAL PRIMARY KEY NOT NULL,
		  Title 			VARCHAR(255) NOT NULL,
			Trader 			VARCHAR(255) NOT NULL,
			Bank 				VARCHAR(255) NOT NULL,
			Currency 		VARCHAR(255) NOT NULL,
			Reputation	SMALLINT,
			Price 			FLOAT,
			Min 				FLOAT,
			Max 				FLOAT,
			Index 			SMALLINT NOT NULL,
		  CreatedAt   TIMESTAMP NOT NULL DEFAULT NOW(),
		  UpdatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`

	_, err := config.DB.Exec(offersTable)
	if err != nil {
		log.Printf("[WARNING - CONFIG - DB] => %v", err)
	}

	return nil
}
