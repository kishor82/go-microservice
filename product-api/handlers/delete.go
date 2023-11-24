package handlers

import (
	"net/http"

	"github.com/kishor82/go-microservices/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product
// responses:
//  201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// DeleteProduct deletes a product from the database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	// this will always convert because of the router
	id := getProductID(r)

	p.l.Debug("deleting record", "id", id)

	err := p.productDB.DeleteProduct(id)

	if err != data.ErrProductNotFound {
		p.l.Error("deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Error("Unable to delete record", "error", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
