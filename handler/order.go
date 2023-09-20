package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/adamelfsborg-code/orders-api-go/model"
	"github.com/adamelfsborg-code/orders-api-go/repository/order"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Order struct {
	Repo *order.RedisRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerID uuid.UUID        `json:"customer_id"`
		LineItems  []model.LineItem `json:"line_items"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
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

	err = o.Repo.Insert(r.Context(), order)
	if err != nil {
		fmt.Println("Failed to Insert: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Failed to encode json: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		fmt.Println("Failed to parse cursor: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const size = 50
	res, err := o.Repo.FindAll(r.Context(), order.FindAllPage{
		Offset: cursor,
		Size:   size,
	})
	if err != nil {
		fmt.Println("Failed to find all: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		Items []model.Order `json:"items"`
		Next  uint64        `json:"next,omitempty"`
	}

	response.Items = res.Orders
	response.Next = res.Cursor

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Failed to marshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64
	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		fmt.Println("Failed to parse idParam to uint: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	row, err := o.Repo.FindByID(r.Context(), orderID)
	if errors.Is(err, order.ErrNotExists) {
		fmt.Println("Order was not found: ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("Failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(row)
	if err != nil {
		fmt.Println("Failed to marshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Status string `json:"status"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")
	const base = 10
	const bitSize = 64
	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		fmt.Println("Failed to parse", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	row, err := o.Repo.FindByID(r.Context(), orderID)
	if errors.Is(err, order.ErrNotExists) {
		fmt.Println("Order was not found: ", err)
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
		if row.ShippedAt != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		row.ShippedAt = &now
	case completedStatus:
		if row.CompletedAt != nil || row.ShippedAt == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		row.CompletedAt = &now
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = o.Repo.UpdateByID(r.Context(), row)
	if err != nil {
		fmt.Println("Failed to update:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(row)
	if err != nil {
		fmt.Println("Failed to marshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an Order by ID")
}
