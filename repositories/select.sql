CREATE OR REPLACE FUNCTION get_nendoroid_by_id(item_num int)
RETURNS TABLE
(
    item_num INT,
    name VARCHAR(255)
)
LANGUAGE plpgsql AS $$ 
DECLARE _id AS INTEGER;
BEGIN
    SELECT id FROM nendoroid INTO _id;
    
    RETURN QUERY
    SELECT text FROM nendoroid_name WHERE nendoroid_id = _id;
END;
$$ LANGUAGE plpgsql
