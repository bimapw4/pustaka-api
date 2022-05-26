package book

import "encoding/json"

type Buku struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}
