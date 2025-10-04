package tooltip

import (
	"fmt"
)

func Render(text string) string {
	return fmt.Sprintf(`<lens-tooltip role="tooltip">%s</lens-tooltip>`, text)
}
