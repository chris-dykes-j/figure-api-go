package repositories

import (
	"context"
	"errors"
	m "figures/models"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type NendoroidRepository struct {
	conn *pgx.Conn
}

func Init() *NendoroidRepository {
	// TODO update connection string
	conn, err := pgx.Connect(context.Background(), "postgres://chris:@localhost:5432/figures")
	if err != nil {
		log.Fatal(err)
	}

	return &NendoroidRepository{
		conn: conn,
	}
}

func (r *NendoroidRepository) GetAllNendoroids(lang string, character string) []m.Nendoroid {
	rows, err := r.conn.Query(context.Background(),
		`
            SELECT
            n.item_number,
            nd.name,
            nd.description,
            nd.item_link,
            nd.blog_link,
            nd.details
            FROM nendoroid AS n
            LEFT JOIN nendoroid_data AS nd ON n.id = nd.nendoroid_id AND nd.language_code = $1
            WHERE nd.name ILIKE '%' || $2 || '%';
        `, // TODO Fix this query to be sargeable.
		lang, character)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// TODO replace append with something efficient. After paging and language code impl.
	var nendoroids []m.Nendoroid
	for rows.Next() {
		var nendo m.Nendoroid
		err = rows.Scan(
			&nendo.ItemNumber, &nendo.Name, &nendo.Description,
			&nendo.ItemLink, &nendo.BlogLink, &nendo.Details)
		if err != nil {
			fmt.Println(err)
		}
		nendoroids = append(nendoroids, nendo)
	}

	return nendoroids
}

func (r *NendoroidRepository) GetNendoroidById(id int, lang string) (m.Nendoroid, error) {
	row, err := r.conn.Query(context.Background(), `

        SELECT
        n.item_number,
        nd.name,
        nd.description,
        nd.item_link,
        nd.blog_link,
        nd.details
        FROM nendoroid AS n
        LEFT JOIN nendoroid_data AS nd ON n.id = nd.nendoroid_id AND nd.language_code = $1
        WHERE n.item_number = $2;

    `, lang, id)
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	var nendo m.Nendoroid
	for row.Next() {
		err = row.Scan(
			&nendo.ItemNumber, &nendo.Name, &nendo.Description,
			&nendo.ItemLink, &nendo.BlogLink, &nendo.Details)
		if err != nil {
			fmt.Println(err)
		}
		row.Close()
	}

	if nendo.ItemNumber == "" {
		message := fmt.Sprintf("Nendoroid does not exist. Given itemNumber %d", id)
		return nendo, errors.New(message)
	}

	return nendo, nil
}
