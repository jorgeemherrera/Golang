package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"microservice-go-redis/model"
	"microservice-go-redis/repository/order"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Order struct {
	Repo *order.RedisRepo
}

func (h *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerID uuid.UUID        `json:"customer_id"`
		LineItems  []model.LineItem `json:"line_items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	order := model.Order{
		OrderID:    rand.Uint64(),
		CustomerID: body.CustomerID,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}

	err := h.Repo.Insert(r.Context(), order)
	if err != nil {
		fmt.Println("Failed to insert:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusCreated)
}

func (h *Order) List(w http.ResponseWriter, r *http.Request) {
	cursorString := r.URL.Query().Get("cursor")
	if cursorString == "" {
		cursorString = "0"
	}

	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorString, decimal, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const size = 50
	resp, err := h.Repo.FindAll(r.Context(), order.FindAllPage{
		Offset: uint(cursor),
		Size:   size,
	})
	if err != nil {
		fmt.Println("Failed to find all:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		Items []model.Order `json:"items"`
		Next  uint64        `json:next,omitempty`
	}
	response.Items = resp.Orders
	response.Next = resp.Cursor

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	const decimal = 10
	const bitSize = 64
	orderID, err := strconv.ParseUint(idParam, decimal, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	o, err := h.Repo.FindByID(r.Context(), orderID)
	if errors.Is(err, order.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("Failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(o); err != nil {
		fmt.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	theOrder, err := h.Repo.FindByID(r.Context(), orderID)
	if errors.Is(err, order.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("Failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	const completedStatus = "completed"
	const shippedStatus = "shipped"
	now := time.Now().UTC()

	switch body.Status {
	case shippedStatus:
		if theOrder.ShippedAt != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		theOrder.ShippedAt = &now
	case completedStatus:
		if theOrder.CompletedAt != nil || theOrder.ShippedAt == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		theOrder.CompletedAt = &now
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.Repo.Update(r.Context(), theOrder)
	if err != nil {
		fmt.Println("Failed to update:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(theOrder); err != nil {
		fmt.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64
	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		fmt.Println("Failed to parse id:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Repo.DeleteByID(r.Context(), orderID)
	if errors.Is(err, order.ErrNotExist) {
		fmt.Println("Failed to delete:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("Failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
