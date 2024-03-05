package services

import (
	"encoding/json"
	"fmt"
	"infoBuild/models"
	"infoBuild/pkg/styles"
	"infoBuild/utils"
	"io"
	"net/http"
	"sync"

	"github.com/charmbracelet/lipgloss"
)

var styleCell = lipgloss.NewStyle().
	Inline(true).
	Width(10)

var stylesRow = lipgloss.NewStyle().MarginLeft(2).PaddingBottom(1)

func GetData(packageName string, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://bundlephobia.com/api/size?package=%s&record=true", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var data models.Data
	if err := json.Unmarshal(body, &data); err != nil {
		return
	}

	if data.Error != nil {
		return
	} else {

		Gzip := styleCell.Copy().
			Foreground(styles.Cyan).
			Render(utils.FormatBytes(+data.Gzip))

		Size := styleCell.Copy().
			Foreground(styles.Gray).
			Render(utils.FormatBytes(data.Size))

		Version := styleCell.Copy().
			Foreground(styles.Gray).
			Render("@" + data.Version)

		Name := styleCell.Copy().
			Foreground(styles.Purple).
			Render(data.Name)

		fmt.Println(stylesRow.Render(Gzip, Size, Version, Name))
	}
}
