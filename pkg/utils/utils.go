package utils

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"rent-a-car/ent"
)

func CarsData(db *ent.Client) error {
	fd, err := os.Open("cars.csv")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("cars data / open:%w", err)
	}

	log.Println("Successfully opened the CSV file")
	defer fd.Close()

	fileReader := csv.NewReader(fd)

	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("cars data / read all :%w", err)
	}

	for _, row := range records[1:] {
		_, err := db.Car.Create().
			SetCarName(row[0]).
			SetMilesPerGallon(row[1]).
			SetCylinders(row[2]).
			SetPower(row[4]).
			SetYear(row[7]).
			Save(context.Background())

		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("cars data :%w", err)
		}

		log.Println("cars data created")
	}

	return nil
}
