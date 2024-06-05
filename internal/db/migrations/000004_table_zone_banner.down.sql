IF EXISTS (SELECT * FROM sysobjects WHERE name='zone_banner' and xtype='U')
    BEGIN
        DROP TABLE cms.zone_banner
    END