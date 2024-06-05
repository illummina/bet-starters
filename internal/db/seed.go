package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func Seed(db *sqlx.DB) error {
	_, err := db.Exec(`
		INSERT INTO cms.zones (id, width, height) VALUES (1, 1920, 400), (2, 1920, 400);
		
		INSERT INTO cms.banners (id, date, background_color, background_image, button_text, button_color, button_background,
							title, description, text_align, link, link_is_external)
		VALUES (1, '2019-09-13T08:00:00Z', '#333333', 'https://cdn.betstarters.com/banners/banner-1920x450.jpg', 'Play Now!',
				'#FFFFFF', '#CDCDCD', 'Play Now!', 'Play Now!', 'left',
				'https://online.sbbet.me/banner-sport/baner-kosarka-polufinale-1920x450/', 1),
			   (2, '2019-09-13T08:00:00Z', '#333333', 'https://cdn.betstarters.com/banners/banner-1920x450.jpg', 'Play Now!',
				'#FFFFFF', '#CDCDCD', 'Play Now!', 'Play Now!', 'left',
				'https://online.sbbet.me/banner-sport/baner-kosarka-polufinale-1920x450/', 1),
			   (3, '2019-09-13T08:00:00Z', '#333333', 'https://cdn.betstarters.com/banners/banner-1920x450.jpg', 'Play Now!',
				'#FFFFFF', '#CDCDCD', 'Play Now!', 'Play Now!', 'left',
				'https://online.sbbet.me/banner-sport/baner-kosarka-polufinale-1920x450/', 1);
		
		INSERT INTO cms.zone_banner (zone_id, banner_id) VALUES (1, 1);
		INSERT INTO cms.zone_banner (zone_id, banner_id) VALUES (2, 1);
		INSERT INTO cms.zone_banner (zone_id, banner_id) VALUES (2, 2);
		INSERT INTO cms.zone_banner (zone_id, banner_id) VALUES (2, 3);
    `)
	if err != nil {
		return fmt.Errorf("failed to seed data: %w", err)
	}

	log.Println("Seed data inserted successfully")
	return nil
}
