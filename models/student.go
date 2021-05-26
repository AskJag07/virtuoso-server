package models

import "time"

type Student struct {
	Full_name  *string
	Email      *string
	Standard   *int
	Created_at time.Time
	Updated_at time.Time
}
