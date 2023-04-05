package renderer

import (
	"fmt"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func (pR *pdfRenderer) Render(filepath string, render func(timeZone *time.Location, timeFormat string, price float32) error) error {
	pdf := pdf.NewMaroto(consts.Portrait, consts.A4)
	pdf.SetPageMargins(10, 15, 10)

	pdf.RegisterHeader(func() {
		pdf.SetBackgroundColor(color.Color{
			Red:   84,
			Green: 84,
			Blue:  84,
		})
		pdf.Row(20, func() {
			pdf.Col(12, func() {
				if err := pdf.FileImage("/workspaces/warp/pkg/internal/renderer/logo.png", props.Rect{
					Center:  true,
					Percent: 80,
				}); err != nil {
					fmt.Println("PDF Header Image not found")
				}
			})
		})
		pdf.SetBackgroundColor(color.NewWhite())

		pdf.Line(10)
	})

	pdf.Row(5, func() {
		pdf.Col(4, func() {
			pdf.Text("Vorname Nachname", props.Text{
				Align: consts.Left,
			})
		})

		pdf.ColSpace(3)

		pdf.Col(5, func() {
			pdf.Text("Wallbox: Warp-Charger", props.Text{
				Align: consts.Left,
			})
		})
	})

	pdf.Row(5, func() {
		pdf.Col(4, func() {
			pdf.Text("Straße Nr", props.Text{
				Align: consts.Left,
			})
		})

		pdf.ColSpace(3)

		pdf.Col(5, func() {
			pdf.Text("PDF erstellt am: <DATE>", props.Text{
				Align: consts.Left,
			})
		})
	})

	pdf.Row(5, func() {
		pdf.Col(4, func() {
			pdf.Text("0000 Ort", props.Text{
				Align: consts.Left,
			})
		})

		pdf.ColSpace(3)

		pdf.Col(5, func() {
			pdf.Text("Exportierte Benutzer: <LIST USERS>", props.Text{
				Align: consts.Left,
			})
		})
	})

	pdf.Row(5, func() {
		pdf.ColSpace(7)

		pdf.Col(5, func() {
			pdf.Text("Exportierte Zeitraum: <EXPORT TIME RNAGe>", props.Text{
				Align: consts.Left,
			})
		})
	})

	pdf.Row(5, func() {
		pdf.ColSpace(7)

		pdf.Col(5, func() {
			pdf.Text("Gesammt kWH: <SUM kWH>", props.Text{
				Align: consts.Left,
			})
		})
	})

	pdf.Row(5, func() {
		pdf.ColSpace(7)

		pdf.Col(5, func() {
			pdf.Text("Kosten: € (35,56 ct/kWh)", props.Text{
				Align: consts.Left,
			})
		})
	})

	pdf.Line(10)

	headers := []string{"Time", "User", "Power meter start", "Power meter end", "Charge (kWh)", "Duration (hh:mm:ss)", "Cost (€)"}
	contents := [][]string{
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
		{"Content1", "Content2", "Content3", "Content4", "Content5", "Content6", "Content7"},
	}

	pdf.TableList(headers, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Normal,
			GridSizes: []uint{2, 1, 2, 2, 2, 2, 1},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			GridSizes: []uint{2, 1, 2, 2, 2, 2, 1},
		},
		Align: consts.Left,
		AlternatedBackground: &color.Color{
			Red:   240,
			Green: 240,
			Blue:  240,
		},
		HeaderContentSpace:     2.0,
		Line:                   false,
		VerticalContentPadding: 4.0,
	})

	return pdf.OutputFileAndClose(filepath)
}
