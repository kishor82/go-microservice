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

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)

	if err != data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
