package models

import "github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form 	  *forms.Form
}
