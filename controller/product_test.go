package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/controller"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/log"
	"net/http/httptest"
	"testing"
)

func TestProductPost(t *testing.T) {
	log := log.New()
	productCtrl := controller.NewProductController(log)
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	reqBody := `{"name":"string","price":0,"image_url": "string","online_stock": {"stock_quantity": 0}}`
	want := entity.Product{
		Name:     "string",
		Price:    0,
		ImageURL: "string",
		OnlineStock: entity.OnlineStock{
			StockQuantity: 0,
		},
	}

	context.Request = httptest.NewRequest("POST", "/", bytes.NewBufferString(reqBody))
	productCtrl.Post(context, func(p *entity.Product) error {
		return nil
	})

	var got entity.Product
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Error(err.Error())
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}

func TestProductGetAll(t *testing.T) {
	log := log.New()
	productCtrl := controller.NewProductController(log)
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = httptest.NewRequest("GET", "/", nil)

	productCtrl.GetAll(context, func() ([]entity.Product, error) {
		return []entity.Product{}, nil
	})

	var got []entity.Product
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Error(err.Error())
	}

	if diff := cmp.Diff(got, []entity.Product{}); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}

func TestProductPatchPurchase(t *testing.T) {
	log := log.New()
	productCtrl := controller.NewProductController(log)
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	reqBody := `{
		"id":1
	}`
	context.Request = httptest.NewRequest("POST", "/", bytes.NewBufferString(reqBody))

	productCtrl.PatchPurchase(context, func(i int) ([]entity.PurchasesProducts, error) {
		return []entity.PurchasesProducts{}, nil
	})

	var got []entity.PurchasesProducts
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Error(err.Error())
	}

	if diff := cmp.Diff(got, []entity.PurchasesProducts{}); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}
