package handlers

import (
	"fmt"
	"net/http"
	"time"

	"fwdlio/views"

	"github.com/a-h/templ"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	views.Index().Render(r.Context(), w)
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		views.Error("Failed to parse form data").Render(r.Context(), w)
		return
	}

	name := r.FormValue("name")
	message := r.FormValue("message")

	// Simulate backend call
	err = callBackend(name, message)

	if err != nil {
		views.Error(fmt.Sprintf("Backend call failed: %v", err)).Render(r.Context(), w)
		return
	}

	views.Success(fmt.Sprintf("Message sent successfully, %s!", name)).Render(r.Context(), w)
}

func callBackend(name, message string) error {
	// Simulate latency
	time.Sleep(500 * time.Millisecond)

	// Simulate random failure (10% chance)
	// if time.Now().UnixNano()%10 == 0 {
	//     return fmt.Errorf("service unavailable")
	// }

	// In a real app, this would be an HTTP POST to the backend
	fmt.Printf("Backend called with Name: %s, Message: %s\n", name, message)
	return nil
}

// Wrapper to convert templ component to http.Handler
func TemplHandler(c templ.Component) http.Handler {
	return templ.Handler(c)
}
