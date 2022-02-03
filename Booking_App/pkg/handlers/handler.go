package handler

import (
	"fmt"
	"net/http"
	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/pkg/config"
	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/pkg/models"
	renderer "github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/pkg/renderers"
)

//the repository used by handlers
var Repo *Repository
//repository type
type Repository struct{
	App *config.AppConfig
}
//create a new repository
func NewRepo(a *config.AppConfig)*Repository{
	return &Repository{
		App: a,
	}
}
//sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {

	remoteAddr := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteAddr)

	//fmt.Fprintf(rw,"This is home page")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again.."


	renderer.RenderTemplate(rw, "home.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
}

//About Page
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	remote_ip := m.App.Session.GetString(r.Context(),"remote_ip")
	fmt.Println(remote_ip)
	stringMap["remote_ip"] = remote_ip
	//fmt.Fprintf(rw,fmt.Sprintf("This is About page and 2+2 is %d",addValues(2,2)))
	renderer.RenderTemplate(rw, "about.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
}
