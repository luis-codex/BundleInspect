package cmd

import (
	"fmt"
	services "infoBuild/pkg/controllers"
	"infoBuild/pkg/header"
	"infoBuild/pkg/styles"
	"infoBuild/utils"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func Footer(startTime time.Time) {
	timeEnd := utils.PrintElapsedTime(startTime)
	fmt.Println(lipgloss.NewStyle().MarginTop(1).MarginBottom(1).MarginLeft(2).
		Render(lipgloss.NewStyle().Italic(true).
			Render(timeEnd), lipgloss.NewStyle().
			Foreground(styles.Gray).Italic(true).
			Render("  https://bundlephobia.com/api")))
}

func AppCli() {
	if len(os.Args) < 2 {
		fmt.Println(styles.ErrorStyle.Render("Package name is required"))
	} else {
		startTime := time.Now()

		var wg sync.WaitGroup
		header.PrintHeader()

		for _, arg := range os.Args[1:] {
			wg.Add(1)
			go services.GetData(arg, &wg)
		}

		wg.Wait()
		Footer(startTime)
	}

}
