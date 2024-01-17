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
    rows, err := r.conn.Query(context.Background(), //"SELECT get_nendoroids()")
    `
        SELECT 
        n.item_number,
        nn.text AS name,
        nd.text AS description,
        nl.text AS item_link,
        nbl.text AS blog_link,
        ndet.details
        FROM
        nendoroid AS n
        LEFT JOIN nendoroid_name as nn on n.id = nn.nendoroid_id and nn.language_code = 'en'
        LEFT JOIN nendoroid_description as nd on n.id = nd.nendoroid_id and nd.language_code = 'en'
        LEFT JOIN nendoroid_link as nl on n.id = nl.nendoroid_id and nl.language_code = 'en'
        LEFT JOIN nendoroid_blog_link as nbl on n.id = nbl.nendoroid_id and nbl.language_code = 'en'
        LEFT JOIN nendoroid_details as ndet on n.id = ndet.nendoroid_id and ndet.language_code = 'en';
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

func (r *NendoroidRepository) GetNendoroidById(id string) m.Nendoroid {
	row, err := r.conn.Query(context.Background(), "SELECT get_nendoroid_by_id($1::int);", id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var nendo m.Nendoroid
	for row.Next() {
		nendo.ItemNumber = id
		err = row.Scan(&nendo.Name)
		if err != nil {
			log.Fatal(err)
		}
		row.Close()
	}

	return nendo
}
