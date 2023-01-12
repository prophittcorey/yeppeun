package web

import (
	"encoding/json"
	"log"
	"net/http"
)

func init() {
	// The index page handler. This is a bit special because it handles the index
	// page ("/") and any pages that don't match a registered route (serves as the
	// catch all handler).
	routes.register(route{
		Path: "/",
		Handler: logger(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":
				var data any
				var cleaned []byte

				if err := json.Unmarshal([]byte(r.FormValue("dirty-json")), &data); err == nil {
					if bs, err := json.MarshalIndent(data, "", "  "); err == nil {
						cleaned = bs
					}
				}

				err := templates.ExecuteTemplate(w, "pages/index.tmpl", map[string]interface{}{
					"Ugly":   r.FormValue("dirty-json"),
					"Pretty": string(cleaned),
				})

				if err != nil {
					log.Printf("web: error rendering pretty page; %s", err)
				}
			case "GET":
				if r.URL.Path != "/" {
					// NOTE: This is the catch all code path. We could do things here like
					// redirect old broken links, render a 404 page, etc.

					http.NotFound(w, r)
				} else {
					if err := templates.ExecuteTemplate(w, "pages/index.tmpl", nil); err != nil {
						log.Printf("web: error rendering index page; %s", err)
					}
				}
			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}
		}),
	})
}
