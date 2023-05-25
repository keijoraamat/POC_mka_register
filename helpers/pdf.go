package helpers

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func CreateFile() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.Row(10, func() {
		m.Text("just things")
	})
	m.Row(10, func() {
		m.Text("right things", props.Text{
			Align: consts.Right,
			Style: consts.BoldItalic,
		})
	})
	m.OutputFileAndClose("test.pdf")
}
