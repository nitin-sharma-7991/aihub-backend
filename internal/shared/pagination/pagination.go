package pagination

import "math"

// NewMeta creates pagination metadata.
func NewMeta(
	req Request,
	total int64,
) Meta {

	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	if totalPages == 0 {
		totalPages = 1
	}

	return Meta{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      total,
		TotalPages: totalPages,
	}
}
