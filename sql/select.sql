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

DROP FUNCTION get_nendoroid_by_id;

-- This didn't really work for scan, so I'll have to fiddle with it later.
-- The joins honestly suck, and a different method would be ideal, like a materialized view or something.
CREATE OR REPLACE FUNCTION get_nendoroids()
RETURNS TABLE (
    item_number int,
    name varchar(255),
    description text,
    item_link text,
    blog_link text,
    details jsonb
)
AS $$ 
BEGIN
    RETURN QUERY 
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
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION get_nendoroids;
