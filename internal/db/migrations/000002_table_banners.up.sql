IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'banners')
    BEGIN
        CREATE TABLE cms.banners
        (
            id                INT PRIMARY KEY,
            date              DATETIME,
            background_color  VARCHAR(20),
            background_image  VARCHAR(255),
            button_text       VARCHAR(50),
            button_color      VARCHAR(20),
            button_background VARCHAR(20),
            title             VARCHAR(100),
            description       VARCHAR(255),
            text_align        VARCHAR(20),
            link              VARCHAR(255),
            link_is_external  BIT
        )
    END