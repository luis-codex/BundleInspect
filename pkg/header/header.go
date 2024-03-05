package header

import (
	"fmt"

	"infoBuild/pkg/styles"

	"github.com/charmbracelet/lipgloss"
)

var stylesHeader = lipgloss.NewStyle().MarginBottom(1)

func PrintHeader() {

	fmt.Println()

	fmt.Println(stylesHeader.Render(
		styles.CreateHeaderCell("gzip", styles.Cyan),
		styles.CreateHeaderCell("size", styles.Gray),
		styles.CreateHeaderCell("version", styles.Gray),
		styles.CreateHeaderCell("name", styles.Purple),
	))
}
