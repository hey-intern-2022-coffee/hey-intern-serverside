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

func TestPurchasePost(t *testing.T) {
	log := log.New()
	purchaseCtrl := controller.NewPurchaseController(log)
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	reqBody := `{
		"name": "string",
		"address": "string",
		"mail_address": "string",
		"purchases_products":[{
			"product_id": 1
		}]
	}`
	want := entity.Purchase{
		Name:        "string",
		Address:     "string",
		MailAddress: "string",
		PurchasesProducts: []entity.PurchasesProducts{
			{
				ProductID: 1,
			},
		},
	}
	context.Request = httptest.NewRequest("POST", "/", bytes.NewBufferString(reqBody))
	purchaseCtrl.Post(context, func(p *entity.Purchase) error {
		return nil
	})

	var got entity.Purchase
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Error(err.Error())
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Create (-want +got) =\n%s\n", diff)
	}
}
