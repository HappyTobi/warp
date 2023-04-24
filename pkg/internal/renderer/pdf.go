package renderer

import (
	"fmt"
	"strings"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func (pR *pdfRenderer) Render(filepath string, render func(timeZone *time.Location, timeFormat string, price float32) ([]string, float32, [][]string, error)) error {
	location, _ := time.LoadLocation(pR.settings.TimeZone)
	timeFormat := pR.settings.TimeFormat

	headers := []string{"Datetime", "User", "Power start", "Power end", "Charge (kWh)", "Duration (hh:mm:ss)", "Cost €"}
	users, totalEnergy, content, err := render(location, timeFormat, pR.settings.Price)
	if err != nil {
		return err
	}

	sumCosts := totalEnergy * pR.settings.Price

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
				if err := pdf.FileImage(pR.settings.LogoHeader, props.Rect{
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

	if pR.settings.PrintHeader {
		pdf.Row(5, func() {
			pdf.Col(4, func() {
				pdf.Text(fmt.Sprintf("%s %s", pR.settings.Settings.Firstname, pR.settings.Settings.Lastname), props.Text{
					Align: consts.Left,
				})
			})

			pdf.ColSpace(3)

			pdf.Col(5, func() {
				pdf.Text("Charger: warp-charger", props.Text{
					Align: consts.Left,
				})
			})
		})

		pdf.Row(5, func() {
			pdf.Col(4, func() {
				pdf.Text(pR.settings.Settings.Street, props.Text{
					Align: consts.Left,
				})
			})

			pdf.ColSpace(3)

			pdf.Col(5, func() {
				pdf.Text(fmt.Sprintf("Exported on: %s", time.Now().In(location).Format(timeFormat)), props.Text{
					Align: consts.Left,
				})
			})
		})

		pdf.Row(5, func() {
			pdf.Col(4, func() {
				pdf.Text(fmt.Sprintf("%s %s", pR.settings.Settings.Postcode, pR.settings.Settings.Street), props.Text{
					Align: consts.Left,
				})
			})

			pdf.ColSpace(3)

			pdf.Col(5, func() {
				pdf.Text(fmt.Sprintf("Exported users: %s", strings.Join(users, ",")), props.Text{
					Align: consts.Left,
				})
			})
		})

		/*pdf.Row(5, func() {
			pdf.ColSpace(7)

			pdf.Col(5, func() {
				pdf.Text("Exported period: <EXPORT TIME RNAGe>", props.Text{
					Align: consts.Left,
				})
			})
		})*/

		pdf.Row(5, func() {
			pdf.ColSpace(7)

			pdf.Col(5, func() {
				pdf.Text(fmt.Sprintf("Total energy kWh: %.2f", float32(totalEnergy)), props.Text{
					Align: consts.Left,
				})
			})
		})

		pdf.Row(5, func() {
			pdf.ColSpace(7)

			pdf.Col(5, func() {
				pdf.Text(fmt.Sprintf("Total costs: %.2f€ (%.2f ct/kWh)", sumCosts, pR.settings.Price), props.Text{
					Align: consts.Left,
				})
			})
		})

		pdf.Line(10)
	}

	pdf.TableList(headers, content, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.Arial,
			Style:     consts.Normal,
			GridSizes: []uint{2, 2, 2, 2, 1, 2, 1},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Normal,
			GridSizes: []uint{2, 2, 2, 2, 1, 2, 1},
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
