package sqlc

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

func (p Product) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s,%s,%s,%f,%s,%s,%s,%s)",
		p.ID,
		p.Name,
		p.Description,
		fmt.Sprintf("{%s}", strings.Join(p.Categories, "\\,")),
		p.Price,
		fmt.Sprintf("{%s}", strings.Join(p.Features, "\\,")),
		p.Color,
		p.Material,
		p.Upc,
	), nil
}
