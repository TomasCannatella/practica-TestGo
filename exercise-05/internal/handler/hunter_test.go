package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testdoubles/internal/handler"
	"testdoubles/internal/hunter"
	"testdoubles/internal/prey"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestHunter_ConfigurePrey(t *testing.T) {
	t.Run("case 01 - prey is configured successfully", func(t *testing.T) {
		// arrange
		// hunter
		ht := hunter.NewHunterMock()
		// prey
		pr := prey.NewPreyStub()

		// handler
		hd := handler.NewHunter(ht, pr)

		// chi router
		r := chi.NewRouter()
		r.Post("/configure-prey", hd.ConfigurePrey())

		req := httptest.NewRequest(http.MethodPost, "/configure-prey", strings.NewReader(`{"speed": 1.0, "position": {"x": 1.0, "y": 1.0, "z": 1.0}}`))
		res := httptest.NewRecorder()

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedStatus := http.StatusOK
		expectedBody := `{"message": "prey configured"}`

		require.Equal(t, expectedStatus, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestHunter_ConfigureHunter(t *testing.T) {
	t.Run("case 01 - hunter is not configured successfully", func(t *testing.T) {
		// arrange
		// hunter
		ht := hunter.NewHunterMock()
		// prey
		pr := prey.NewPreyStub()
		// handler
		hd := handler.NewHunter(ht, pr)

		// chi router
		r := chi.NewRouter()
		r.Post("/configure-hunter", hd.ConfigureHunter())

		req := httptest.NewRequest(http.MethodPost, "/configure-hunter", strings.NewReader(`{"speed": "1.0", "position": {"x": 1.0, "y": 1.0, "z": 1.0}}`))
		res := httptest.NewRecorder()

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message": "hunter not configured", "status": "Bad Request"}`

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestHunter_Hunt(t *testing.T) {
	t.Run("case 01 - hunter hunts the prey successfully", func(t *testing.T) {
		// arrange
		// hunter
		ht := hunter.NewHunterMock()
		// prey
		pr := prey.NewPreyStub()
		// handler
		hd := handler.NewHunter(ht, pr)
		r := chi.NewRouter()
		r.Post("/hunt", hd.Hunt())

		req := httptest.NewRequest(http.MethodPost, "/hunt", nil)
		res := httptest.NewRecorder()

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message":"hunt done","duration": 0.0}`

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())

	})

	t.Run("case 02 - hunter can not hunt the prey", func(t *testing.T) {
		// arrange
		// hunter
		ht := hunter.NewHunterMock()
		ht.HuntFunc = func(p prey.Prey) (float64, error) {
			return 0.0, hunter.ErrCanNotHunt
		}
		// prey
		pr := prey.NewPreyStub()
		// handler
		hd := handler.NewHunter(ht, pr)
		r := chi.NewRouter()
		r.Post("/hunt", hd.Hunt())

		req := httptest.NewRequest(http.MethodPost, "/hunt", nil)
		res := httptest.NewRecorder()

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := `{"message":"hunter could not hunt the prey","status":"Internal Server Error"}`

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 03 - internal server error", func(t *testing.T) {
		// arrange
		// hunter
		ht := hunter.NewHunterMock()
		ht.HuntFunc = func(p prey.Prey) (float64, error) {
			return 0.0, errors.New("unknown error")
		}
		// prey
		pr := prey.NewPreyStub()
		// handler
		hd := handler.NewHunter(ht, pr)
		r := chi.NewRouter()
		r.Post("/hunt", hd.Hunt())

		req := httptest.NewRequest(http.MethodPost, "/hunt", nil)
		res := httptest.NewRecorder()

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := `{"message":"internal server error","status":"Internal Server Error"}`

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}
