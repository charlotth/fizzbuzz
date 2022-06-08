package server

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/tcharlot-datasweet/fizzbuzz/pkg/fizzbuzz"
)

// FizzbuzzRequest is a request to /api/fizzbuzz
type FizzbuzzRequest struct {
	Limit        int    `json:"limit,omitempty"`
	FizzMultiple int    `json:"int1,omitempty"`
	FizzString   string `json:"str1,omitempy"`
	BuzzMultiple int    `json:"int2,omitempty"`
	BuzzString   string `json:"str2,omitempy"`
}

// Bind is from go-chi
func (req *FizzbuzzRequest) Bind(r *http.Request) error {
	return nil
}

func (s *server) postFizzbuzz(w http.ResponseWriter, r *http.Request) {
	var req FizzbuzzRequest

	if err := render.Bind(r, &req); err != nil && !errors.Is(err, io.EOF) {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// Defaulting
	if req.Limit < 1 {
		req.Limit = 100
	}
	if req.FizzMultiple <= 0 {
		req.FizzMultiple = 3
	}
	if len(req.FizzString) == 0 {
		req.FizzString = "fizz"
	}
	if req.BuzzMultiple <= 0 {
		req.BuzzMultiple = 5
	}
	if len(req.BuzzString) == 0 {
		req.BuzzString = "buzz"
	}

	// call fizzbuzz
	str := fizzbuzz.String(
		fizzbuzz.To(req.Limit),
		fizzbuzz.Fizz(req.FizzMultiple, req.FizzString),
		fizzbuzz.Buzz(req.BuzzMultiple, req.BuzzString),
	)

	render.PlainText(w, r, str)
}

func (s *server) getFizzbuzz(w http.ResponseWriter, r *http.Request) {
	req := FizzbuzzRequest{
		Limit:        QueryInt(r, "limit", 100),
		FizzMultiple: QueryInt(r, "int1", 3),
		FizzString:   QueryString(r, "str1", "fizz"),
		BuzzMultiple: QueryInt(r, "int2", 5),
		BuzzString:   QueryString(r, "str2", "buzz"),
	}

	// call fizzbuzz
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusAccepted)
	fizzbuzz.Write(w,
		fizzbuzz.To(req.Limit),
		fizzbuzz.Fizz(req.FizzMultiple, req.FizzString),
		fizzbuzz.Buzz(req.BuzzMultiple, req.BuzzString),
	)
}
