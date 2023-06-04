package helpers

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type FindingActPdf struct {
	ID               uint
	FinderName       string
	FinderIdNumber   string
	RecieverName     string
	FindingType      string
	FindersFee       bool
	ResiginOwnership bool
	RemainAnonymous  bool
	TransferLocation string
	TransferDate     string
	WDActNumber      string
	Artefacts        string
}

func CreateFile(findingAct *FindingActPdf) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(30, 15, 15)

	m.Row(10, func() {
		m.Text(findingAct.TransferLocation, props.Text{
			Align: consts.Right,
			Style: consts.BoldItalic,
		})
	})
	m.Row(10, func() {
		m.Text(findingAct.TransferDate, props.Text{
			Align: consts.Right,
			Style: consts.BoldItalic,
		})
	})

	m.Row(12, func() {
		m.Text("Arheoloogiliste leidude", props.Text{
			Align: consts.Center,
			Style: consts.Bold,
		})
	})
	m.Row(12, func() {
		m.Text("üleandmise-vastuvõtmise akt nr "+findingAct.WDActNumber, props.Text{
			Align: consts.Center,
			Style: consts.Bold,
		})
	})

	m.Row(10, func() {
		m.Text("Tulenevalt Muinsuskaitseseadusest § 24, 27–28 Muinsuskaitseamet (edaspidi vastuvõtja või MKA) registrikood 70000958, keda esindab arheoloogia nõunik " + findingAct.RecieverName + " ja " + findingAct.FinderName + " isikukood: " + findingAct.FinderIdNumber + ", telefoni nr …., meiliaadress …. (edaspidi üleandja) koostasid kultuuriväärtusega arheoloogiliste leidude üleandmise-vastuvõtmise akti alljärgnevas:")
	})

	m.Row(10, func() {
		m.Text("\n\n\n\n\n")
	})

	m.Row(10, func() {
		m.Text("      \t1. Üleandja annab üle ja vastuvõtja võtab vastu järgmised arheoloogilised leiud: kokku " + findingAct.Artefacts + " leidu")
	})

	m.Row(10, func() {
		m.Text("      \t2. Leiukoht ja leidude leidmise asjaolud: Leitud otsinguvahendiga, luba olemas, otsinguteavitused ja otsinguaruanded esitatud.")
	})

	m.Row(10, func() {
		m.Text("      \t3. Leidude üleandmine Muinsuskaitseametile: mujal. Ameti kontoris " + findingAct.TransferLocation + " " + findingAct.TransferDate + ".")
	})

	m.Row(10, func() {
		m.Text("      \t4. Üleandja " + parseRequest(findingAct.FindersFee) + " leiuautasu")
	})

	m.Row(10, func() {
		m.Text("      \t5. Üleandja " + parseWish(findingAct.ResiginOwnership) + " riigi loobumist omandiõigusest")
	})

	m.Row(10, func() {
		m.Text("      \t6. Üleandja " + parseWish(findingAct.RemainAnonymous) + " jääda anonüümseks")
	})

	m.Row(10, func() {
		m.Text("\n\n\n\n\n")
	})

	m.Row(10, func() {
		m.Text("Üleandja:                                                                               \t\t\t\t\tVasuvõtja:", props.Text{
			Align: consts.Center,
			Style: consts.Bold,
		})
	})

	m.Row(10, func() {
		m.Text(findingAct.FinderName+"                                                                   \t\t\t\t\t"+findingAct.RecieverName, props.Text{
			Align: consts.Center,
		})
	})

	m.Row(10, func() {
		m.Text("/allkijastatud digitaalselselt/                                                             \t\t\t\t\t/allkijastatud digitaalselselt/", props.Text{
			Align: consts.Center,
		})
	})

	m.OutputFileAndClose("test.pdf")
}

func parseWish(wish bool) (strWish string) {
	if wish {
		strWish = "soovib"
		return
	}
	strWish = "ei soovi"
	return
}

func parseRequest(req bool) (strReq string) {
	if req {
		strReq = "taotleb"
		return
	}
	strReq = "ei taotle"
	return
}
