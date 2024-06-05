IF EXISTS (SELECT * FROM sysobjects WHERE name='banners' and xtype='U')
    BEGIN
        DROP TABLE cms.banners
    END
