package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/platform/web/response"
)

// NewHunter returns a new Hunter handler.
func NewHunter(ht hunter.Hunter, pr prey.Prey) *Hunter {
	return &Hunter{ht: ht, pr: pr}
}

// Hunter returns handlers to manage hunting.
type Hunter struct {
	// ht is the Hunter interface that this handler will use
	ht hunter.Hunter
	// pr is the Prey interface that the hunter will hunt
	pr prey.Prey
}

// RequestBodyConfigPrey is an struct to configure the prey for the hunter in JSON format.
type RequestBodyConfigPrey struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigurePrey configures the prey for the hunter.
func (h *Hunter) ConfigurePrey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestBody RequestBodyConfigPrey
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "prey not configured")
			return
		}
		// process
		h.pr.Configure(requestBody.Speed, requestBody.Position)

		// response
		response.JSON(w, http.StatusOK, map[string]string{"message": "prey configured"})
	}
}

// RequestBodyConfigHunter is an struct to configure the hunter in JSON format.
type RequestBodyConfigHunter struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigureHunter configures the hunter.
func (h *Hunter) ConfigureHunter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestBody RequestBodyConfigHunter
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "hunter not configured")
			return
		}

		// process
		h.ht.Configure(requestBody.Speed, requestBody.Position)

		// response
		response.JSON(w, http.StatusOK, map[string]string{"message": "hunter configured"})
	}
}

// Hunt hunts the prey.
func (h *Hunter) Hunt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		duration, err := h.ht.Hunt(h.pr)
		if err != nil {
			if errors.Is(err, hunter.ErrCanNotHunt) {
				response.Error(w, http.StatusInternalServerError, "hunter could not hunt the prey")
				return
			}
			response.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
		// response
		response.JSON(w, http.StatusOK, map[string]any{"message": "hunt done", "duration": duration})
	}
}
