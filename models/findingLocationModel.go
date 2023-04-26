package models

import (
	"time"

	"gorm.io/gorm"
)

// Creates SQL:
//
//	CREATE TABLE public.finding_locations (
//		location_id int8 NOT NULL,
//		finding_act_id int8 NOT NULL,
//		created_at timestamptz NULL,
//		updated_at timestamptz NULL,
//		deleted_at timestamptz NULL,
//		"from" timestamptz NULL,
//		"until" timestamptz NULL,
//		id bigserial NOT NULL,
//		CONSTRAINT finding_locations_pkey PRIMARY KEY (location_id, finding_act_id),
//		CONSTRAINT fk_finding_locations_finding_act FOREIGN KEY (finding_act_id) REFERENCES public.finding_acts(id),
//		CONSTRAINT fk_finding_locations_location FOREIGN KEY (location_id) REFERENCES public.locations(id)
//	);
type FindingLocation struct {
	gorm.Model
	FindingActID uint `gorm:"primaryKey"`
	FindingAct   FindingAct
	LocationID   uint `gorm:"primaryKey"`
	Location     Location
	From         time.Time
	Until        time.Time
}
