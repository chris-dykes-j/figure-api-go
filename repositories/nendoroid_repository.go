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
        SELECT text from nendoroid_name
        WHERE language_code = 'en';
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var nendoroids []m.Nendoroid
	for rows.Next() {
		var nendo m.Nendoroid
		err = rows.Scan(&nendo.English.Name)
		if err != nil {
			log.Fatal(err)
		}
		nendoroids = append(nendoroids, nendo)
	}

	return nendoroids
}

func (r *NendoroidRepository) GetNendoroidById(id string) m.Nendoroid {
	query, err := r.conn.Query(context.Background(), `
        CREATE OR REPLACE FUNCTION get_nendoroid_by_id(item_num int)
        RETURNS TABLE (name varchar(255)) AS $$ 
        DECLARE _id INTEGER;
        BEGIN
            SELECT id FROM nendoroid WHERE item_number = item_num INTO _id;
            
            RETURN QUERY
            SELECT text FROM nendoroid_name
            WHERE nendoroid_id = _id
            AND language_code = 'en';
        END;
        $$ LANGUAGE plpgsql;

        `)
	if err != nil {
		log.Fatal(err)
	}
    query.Close();

    row, err := r.conn.Query(context.Background(), "SELECT get_nendoroid_by_id($1::int);", id)

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var nendo m.Nendoroid
	for row.Next() {
        nendo.English.ItemNumber = id
		err = row.Scan(&nendo.English.Name)
		if err != nil {
			log.Fatal(err)
		}
        row.Close()
	}

	return nendo
}
