IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'zones')
    BEGIN
        CREATE TABLE cms.zones
        (
            id     INT PRIMARY KEY,
            width  INT,
            height INT
        )
    END