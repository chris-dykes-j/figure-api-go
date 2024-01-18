package repositories

import (
	"context"
	m "figures/models"
	"log"

	"github.com/jackc/pgx/v5"
)

type NendoroidRepository struct {
	conn *pgx.Conn
}

func Init() *NendoroidRepository {
	// Fix later
	conn, err := pgx.Connect(context.Background(), "postgres://chris:@localhost:5432/figures")
	if err != nil {
		log.Fatal(err)
	}

	return &NendoroidRepository{
		conn: conn,
	}
}

func (r *NendoroidRepository) GetAllNendoroids() []m.Nendoroid {
    rows, err := r.conn.Query(context.Background(), `

        SELECT
        n.item_number,
        nd.name,
        nd.description,
        nd.item_link,
        nd.blog_link,
        nd.details
        FROM nendoroid AS n
        LEFT JOIN nendoroid_data AS nd ON n.id = nd.nendoroid_id AND nd.language_code = 'en';

    `) 
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

	var nendoroids []m.Nendoroid
	for rows.Next() {
		var nendo m.Nendoroid
		err = rows.Scan(
            &nendo.ItemNumber, &nendo.Name, &nendo.Description,
            &nendo.ItemLink, &nendo.BlogLink, &nendo.Details)
		if err != nil {
			log.Fatal(err)
		}
		nendoroids = append(nendoroids, nendo)
	}

	return nendoroids
}

func (r *NendoroidRepository) GetNendoroidById(id int) m.Nendoroid {
	row, err := r.conn.Query(context.Background(),`

        SELECT
        n.item_number,
        nd.name,
        nd.description,
        nd.item_link,
        nd.blog_link,
        nd.details
        FROM nendoroid AS n
        LEFT JOIN nendoroid_data AS nd ON n.id = nd.nendoroid_id AND nd.language_code = 'en'
        WHERE n.item_number = $1;

    `, id)
    if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var nendo m.Nendoroid
	for row.Next() {
		err = row.Scan(
            &nendo.ItemNumber, &nendo.Name, &nendo.Description,
            &nendo.ItemLink, &nendo.BlogLink, &nendo.Details)
		if err != nil {
			log.Fatal(err)
		}
		row.Close()
	}

	return nendo
}
