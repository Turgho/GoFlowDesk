package handler

import (
	"net/http"
	"time"

	"github.com/Turgho/GoFlowDesk/internal/handler/render"
)

// APIResponse defines a standard envelope for JSON responses.
//
// Fields:
//   Data    - successful payload (object or array). Omitted when nil.
//   Message - short human‑readable message about the outcome.
//   Errors  - collection of error strings, typically used on failures.
//   Meta    - optional extra metadata (pagination, counts, etc.).
//   Time    - server timestamp, useful for debugging and caching.
//
// The envelope makes it easy for clients to handle both success and
// error conditions without separate schemas.
//
// Example:
//   {"data":{...},"message":"ok","time":"..."}
//   {"errors":["email exists"],"message":"conflict","time":"..."}
//
// Anyone returning JSON from handlers should wrap their payload in this
// struct and then call render.WriteJSON.

type APIResponse struct {
	Data    any       `json:"data,omitempty"`
	Message string    `json:"message,omitempty"`
	Errors  []string  `json:"errors,omitempty"`
	Meta    any       `json:"meta,omitempty"`
	Time    time.Time `json:"time"`
}

// Send writes the envelope to the http.ResponseWriter using the
// provided status code.  It ensures the timestamp is set and hides
// the render.WriteJSON call from handlers.
func (r *APIResponse) Send(w http.ResponseWriter, status int) error {
	if r.Time.IsZero() {
		r.Time = time.Now().UTC()
	}
	return render.WriteJSON(w, status, r, nil)
}
