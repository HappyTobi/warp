package chargeTracker

import (
	"encoding/csv"
	"fmt"
	"strings"
	"time"

	"github.com/HappyTobi/warp/pkg/cmd/tools"
	"github.com/HappyTobi/warp/pkg/internal/chargeTracker"
	"github.com/HappyTobi/warp/pkg/internal/renderer"
	"github.com/HappyTobi/warp/pkg/internal/users"
	"github.com/HappyTobi/warp/pkg/internal/warp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ChargeLog(cmd *cobra.Command, args []string) error {
	requests := make([]*warp.Request, 0, 2)

	request := &warp.Request{
		Path:        "charge_tracker/charge_log",
		ContentType: warp.JSON,
	}

	userRequest := &warp.Request{
		Path:        "users/all_usernames",
		ContentType: warp.JSON,
	}

	requests = append(requests, request, userRequest)

	/*
		get filter params
	*/
	userFilter, _ := cmd.Flags().GetString("user")
	monthFilter, _ := cmd.Flags().GetInt("month")
	yearFilter, _ := cmd.Flags().GetInt("year")

	mFilter := chargeTracker.NewFilter("month", "", monthFilter)
	yFilter := chargeTracker.NewFilter("year", "", yearFilter)
	uFilter := chargeTracker.NewFilter("user", userFilter, -1)

	filters := chargeTracker.Filters(mFilter, yFilter, uFilter)

	outputFlag, _ := cmd.Flags().GetString("output")

	if err := tools.LoadGlobalParams(cmd, func(charger, username, password, output string) {
		for _, req := range requests {
			req.Warp = charger

			if len(username) > 0 && len(password) > 0 {
				req.Username = username
				req.Password = password
			}

			if renderer.IsRenderer(outputFlag) {
				req.OutputRenderer = renderer.NewRenderer(outputFlag)
			}
		}
	}); err != nil {
		return err
	}

	chargeLog := chargeTracker.NewChargeLog(request)
	user := users.NewUsersList(userRequest)
	users, _ := user.Load()

	charges, err := chargeLog.Load(users, filters)
	if err != nil {
		return err
	}

	if requests[0].OutputRenderer != nil {
		fmt.Print(requests[0].OutputRenderer.Render(charges))
		return nil
	}

	filePath, _ := cmd.Flags().GetString("file")
	switch outputFlag {
	case "csv":
		if err = RenderCsv(filePath, charges); err != nil {
			return err
		}
	case "pdf":
		if err = RenderPdf(filePath, charges); err != nil {
			return err
		}
	}
	fmt.Printf("%s file written to: %s \n", strings.ToUpper(outputFlag), filePath)

	return nil
}

func RenderPdf(filePath string, charges *chargeTracker.Charges) error {
	price := (viper.GetFloat64("settings.power.price") / 100)

	pdfSettings := &renderer.PdfSettings{
		Price:       float32(price),
		TimeFormat:  viper.GetString("settings.date_time.time_format"),
		TimeZone:    viper.GetString("settings.date_time.time_zone"),
		PrintHeader: viper.GetBool("pdf.print_header"),
		LogoHeader:  viper.GetString("pdf.image_path"),
		Settings: renderer.GlobalSettings{
			Firstname: viper.GetString("settings.user.firstname"),
			Lastname:  viper.GetString("settings.user.lastname"),
			Street:    viper.GetString("settings.user.street"),
			Postcode:  viper.GetString("settings.user.postcode"),
			City:      viper.GetString("settings.user.city"),
		},
	}

	pdfRenderer := renderer.NewPdfRenderer(pdfSettings)

	return pdfRenderer.Render(filePath, func(timeZone *time.Location, timeFormat string, price float32) ([]string, float32, [][]string, error) {
		content := make([][]string, len(charges.Charges)+2)
		sumPerUser := make(map[string]float32)
		totalEnergy := float32(0)
		users := make([]string, 0)

		line := 0
		for _, charge := range charges.Charges {
			charged := charge.PowerMeterEnd - charge.PowerMeterStart
			paid := charged * price
			sumPerUser[charge.User] += paid
			row := []string{
				charge.Time.In(timeZone).Format(timeFormat),
				charge.User,
				fmt.Sprintf("%.2f", charge.PowerMeterStart),
				fmt.Sprintf("%.2f", charge.PowerMeterEnd),
				fmt.Sprintf("%.2f", charged),
				charge.Duration,
				fmt.Sprintf("%.2f", paid),
			}

			totalEnergy += charged
			content[line] = row
			line++
		}

		for user, sum := range sumPerUser {
			content[line] = []string{"", user, "", "", "", "", fmt.Sprintf("%.2f", sum)}
			users = append(users, user)
			line++
		}

		return users, totalEnergy, content, nil
	})
}

func RenderCsv(filePath string, charges *chargeTracker.Charges) error {
	price := (viper.GetFloat64("settings.power.price") / 100)
	csvSettings := &renderer.CsvSettings{
		Price:         float32(price),
		TimeFormat:    viper.GetString("settings.date_time.time_format"),
		TimeZone:      viper.GetString("settings.date_time.time_zone"),
		Comma:         viper.GetString("csv.comma"),
		HeaderEnabled: viper.GetBool("csv.header"),
	}

	csvRenderer := renderer.NewCsvRenderer(csvSettings)
	return csvRenderer.Render(filePath, func(writer *csv.Writer, timeZone *time.Location, timeFormat string, price float32) error {
		sumPerUser := make(map[string]float32)
		for _, charge := range charges.Charges {
			charged := charge.PowerMeterEnd - charge.PowerMeterStart
			paid := charged * price
			sumPerUser[charge.User] += paid
			row := []string{
				charge.Time.In(timeZone).Format(timeFormat),
				charge.User,
				fmt.Sprintf("%.2f", charge.PowerMeterStart),
				fmt.Sprintf("%.2f", charge.PowerMeterEnd),
				fmt.Sprintf("%.2f", charged),
				charge.Duration,
				fmt.Sprintf("%.2f", paid),
			}
			if err := writer.Write(row); err != nil {
				return err
			}
		}

		for u, p := range sumPerUser {
			sumPaiment := []string{"", u, "", "", "", "", fmt.Sprintf("%.2f", p)}
			if err := writer.Write(sumPaiment); err != nil {
				return err
			}
		}

		return nil

	})
}
