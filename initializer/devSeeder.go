package initializer

import "github.com/keijoraamat/mka_register/models"

func SeedFindingsActs() error {

	var acts = make([]models.Act, 3)

	var one = &models.Act{
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
	var two = &models.Act{
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
	var three = &models.Act{
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

	result := DB.Create(&acts)
	return result.Error
}
