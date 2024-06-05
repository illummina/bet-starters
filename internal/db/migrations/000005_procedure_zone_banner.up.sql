CREATE PROCEDURE cms.spGetBanners(
    @DeviceTypeId TINYINT,
    @LanguageCode VARCHAR(30),
    @PageId INT = 1,
    @PageSize INT = 20,
    @SortBy NVARCHAR(10) = 'asc',
    @ErrorMessage VARCHAR(1000) OUTPUT
)
AS
BEGIN
    SET NOCOUNT ON;
    DECLARE @Offset INT = (@PageId - 1) * @PageSize;

    BEGIN TRY
        SELECT
            Z.id AS zoneId,
            Z.width AS width,
            Z.height AS height,
            (
                SELECT
                    B.*
                FROM zone_banner AS ZB
                         INNER JOIN banners AS B ON ZB.banner_id = B.id
                WHERE ZB.zone_id = Z.id
                FOR JSON PATH
            ) AS banners
        FROM
            (
                SELECT
                    id,
                    width,
                    height,
                    ROW_NUMBER() OVER (
                        ORDER BY
                            CASE WHEN @SortBy = 'desc' THEN id END DESC,
                            CASE WHEN @SortBy = 'count_desc' THEN COUNT(ZB.banner_id) END DESC,
                            CASE WHEN @SortBy = 'count_asc' THEN COUNT(ZB.banner_id) END,
                            id
                        ) AS row_number
                FROM
                    zones AS Z
                        LEFT JOIN
                    zone_banner AS ZB ON Z.id = ZB.zone_id
                GROUP BY
                    Z.id, Z.width, Z.height
            ) AS Z
        ORDER BY
            Z.row_number
        OFFSET
            @Offset ROWS
            FETCH NEXT
                @PageSize ROWS ONLY
        FOR JSON PATH;
    END TRY
    BEGIN CATCH
        SET @ErrorMessage = ERROR_MESSAGE();
    END CATCH;
END;