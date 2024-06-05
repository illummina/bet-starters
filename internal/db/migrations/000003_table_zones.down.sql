IF EXISTS (SELECT * FROM sysobjects WHERE name='zones' and xtype='U')
    BEGIN
        DROP TABLE cms.zones
    END
