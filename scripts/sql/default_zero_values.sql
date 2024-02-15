CREATE OR REPLACE FUNCTION report_default_go_sql_zero_values_mismatch()
    RETURNS SETOF information_schema.columns
    AS $$
BEGIN
    RETURN QUERY
    SELECT
        *
    FROM
        information_schema.columns
    WHERE (table_schema = 'public'
        AND column_default IS NOT NULL)
        AND (
            (data_type = 'boolean' AND column_default <> 'false'::boolean)
            OR (data_type IN ('char', 'character', 'varchar', 'character varying', 'text')
                AND column_default NOT LIKE '''%''')
            OR (data_type IN ('smallint', 'integer', 'bigint', 'smallserial', 'serial', 'bigserial')
                AND column_default <> '0' AND column_default NOT LIKE 'nextval(%'::text)
            OR (data_type IN ('decimal', 'numeric', 'real', 'double precision')
                AND column_default <> '0.0')
        );
END
$$
LANGUAGE plpgsql
SECURITY DEFINER;

CREATE OR REPLACE FUNCTION report_mismatched_default_zero_values()
    RETURNS void
    AS $$
DECLARE
    item record;
BEGIN
    FOR item IN SELECT * FROM report_default_go_sql_zero_values_mismatch()
    LOOP
        RAISE NOTICE 'Mismatch in %.% (Type: %): Default value ''%'' might not align with Go zero value.', item.table_name, item.column_name, item.data_type, item.column_default;
    END LOOP;
END;
$$
LANGUAGE plpgsql;

-- Execute the report function to find mismatches
SELECT report_mismatched_default_zero_values();
