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
