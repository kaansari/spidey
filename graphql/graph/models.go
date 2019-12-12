package graph

import "time"

type Account struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Orders        []Order    `json:"orders"`
	Email         string     `json:"email"`
	PhotoURL      string     `json:"photoURL"`
	MobilePhone   string     `json:"mobilePhone"`
	Type          string     `json:"type"`
	Locations     []Location `json:"locations"`
	EmailVerified bool       `json:"emailVerified"`
	IsActive      bool       `json:"isActive"`
	CreatedAt     time.Time  `json:"createdAt"`
}

type Order struct {
	ID         string           `json:"id"`
	CreatedAt  time.Time        `json:"createdAt"`
	TotalPrice float64          `json:"totalPrice"`
	Products   []OrderedProduct `json:"products"`
}
