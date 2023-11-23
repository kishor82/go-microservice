.PHONY: setup-go-work

setup-go-work: ## Sets up your go.work file
	@echo "Creating a go.work file"
	rm -f go.work
	go work init
	go work use ./currency
	go work use ./product-api
	go work user ./product-images
