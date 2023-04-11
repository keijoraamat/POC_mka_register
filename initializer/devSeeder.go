package initializer

import (
	"errors"
	"log"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/gorm"
)

func SeedDatabase() {
	seedFindingsActs()
}

func seedFindingsActs() {

	var acts = make([]models.FindingAct, 3)

	var one = &models.FindingAct{
		FinderName:       "Mari Maasikas",
		FinderIdNumber:   "EE12345667890",
		RecieverName:     "Tuule Tallerma",
		FindingType:      "juhu",
		FindersFee:       false,
		ResiginOwnership: true,
		RemainAnonymous:  false,
		TransferLocation: "Fellin",
		WDActNumber:      "33.66453",
		Status:           "töös",
	}
	var two = &models.FindingAct{
		FinderName:       "Mati Kurikas",
		FinderIdNumber:   "EE2345678900",
		RecieverName:     "Tuule Tallerma",
		FindingType:      "detektor",
		FindersFee:       true,
		ResiginOwnership: true,
		RemainAnonymous:  false,
		TransferLocation: "Dorpat",
		WDActNumber:      "33.234523",
		Status:           "arhiiv",
	}
	var three = &models.FindingAct{
		FinderName:       "Mari Maasikas",
		FinderIdNumber:   "EE12345667890",
		RecieverName:     "Tuuve Tallerma",
		FindingType:      "detektor",
		FindersFee:       false,
		ResiginOwnership: true,
		RemainAnonymous:  false,
		TransferLocation: "Lindanise",
		WDActNumber:      "44.66453",
		Status:           "töös",
	}

	acts[0] = *one
	acts[1] = *two
	acts[2] = *three

	check := DB.First(&one)
	log.Println("Dev db seeding result:", check.RowsAffected)
	log.Println("Dev db seeding error:", check.Error)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		DB.Create(&acts)
	}

}
