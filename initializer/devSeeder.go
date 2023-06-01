package initializer

import (
	"errors"
	"log"
	"time"

	"github.com/keijoraamat/mka_register/models"
	"gorm.io/gorm"
)

func SeedDatabase() {
	seedFindingsActs()
}

func seedFindingsActs() {

	var acts = make([]models.FindingAct, 3)
	now := time.Now().AddDate(0, 0, -9)
	lastMonth := now.AddDate(0, -1, 11)
	twoMonthsAgo := now.AddDate(0, -5, -6)

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
		TransferDate:     now,
	}
	var two = &models.FindingAct{
		FinderName:       "Mati Kurikas",
		FinderIdNumber:   "EE2345678900",
		RecieverName:     "Tuule Kallermaa",
		FindingType:      "detektor",
		FindersFee:       true,
		ResiginOwnership: true,
		RemainAnonymous:  false,
		TransferLocation: "Dorpat",
		WDActNumber:      "33.234523",
		Status:           "arhiiv",
		TransferDate:     lastMonth,
	}
	var three = &models.FindingAct{
		FinderName:       "Mari Maasikas",
		FinderIdNumber:   "EE12345667890",
		RecieverName:     "Tuuve Mallerma",
		FindingType:      "detektor",
		FindersFee:       false,
		ResiginOwnership: true,
		RemainAnonymous:  false,
		TransferLocation: "Lindanise",
		WDActNumber:      "44.66453",
		Status:           "töös",
		TransferDate:     twoMonthsAgo,
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
