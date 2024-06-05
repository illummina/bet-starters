IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'zone_banner')
    BEGIN
        CREATE TABLE cms.zone_banner (
            zone_id INT,
            banner_id INT,
            PRIMARY KEY (zone_id, banner_id),
            FOREIGN KEY (zone_id) REFERENCES cms.zones (id),
            FOREIGN KEY (banner_id) REFERENCES cms.banners (id)
        );
    END