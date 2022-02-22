package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/internal/config"
	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/internal/forms"
	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/internal/models"
	renderer "github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/internal/renderers"
)

//the repository used by handlers
var Repo *Repository

 //repository type
type Repository struct {
	App *config.AppConfig
}

//create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {

	remoteAddr := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteAddr)

	//fmt.Fprintf(rw,"This is home page")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again.."

	renderer.RenderTemplate(rw,r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
//About Page
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	remote_ip := m.App.Session.GetString(r.Context(), "remote_ip")
	fmt.Println(remote_ip)
	stringMap["remote_ip"] = remote_ip
	//fmt.Fprintf(rw,fmt.Sprintf("This is About page and 2+2 is %d",addValues(2,2)))
	renderer.RenderTemplate(rw, r,"about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(rw, r,"generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(rw,r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {	
	renderer.RenderTemplate(rw,r, "reservation.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PostReservation(rw http.ResponseWriter, r *http.Request) {	
	err := r.ParseForm()
	if err!=nil{
		log.Println(err)
		return
	}

}
func (m *Repository) Availability(rw http.ResponseWriter, r *http.Request) {	
	renderer.RenderTemplate(rw,r, "search-availability.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}//Availability

func (m *Repository) PostAvailability(rw http.ResponseWriter, r *http.Request) {	
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	rw.Write([]byte(fmt.Sprintf("start date:%s,end date:%s",start,end)))
	// rw.Write([]byte("Request Posted Successfully..."))
}

type jsonResponse struct{
	OK bool `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(rw http.ResponseWriter, r *http.Request) {	
	resp := jsonResponse{
		OK: true,
		Message: "Available",
	}

	out,err := json.MarshalIndent(resp,"","     ")
	if err!=nil {
		log.Println(err)
	}

 	rw.Header().Set("Content-Type","application/json")
	rw.Write(out)
}


