.DEFAULT_GOAL := help

# Declare targets that aren't actual files on disk
.PHONY: run build exec help

help: ## Display this help message
	@powershell -NoProfile -Command "Get-Content '$(MAKEFILE_LIST)' | Where-Object { $$_ -match '^[a-zA-Z_-]+:.*##' } | ForEach-Object { $$parts = $$_ -split ':.*?##'; $$t = $$parts[0].Trim(); $$d = $$parts[1].Trim(); Write-Host ('  {0,-15}{1}' -f $$t, $$d) }"

run: ## Run the application without building it
	go run ./cmd/main.go

build: ## Build the application as an executable file
	go build -o ./cmd/cmd.exe ./cmd/main.go

exec: build ## Execute the application from the automatically built executable file
	./cmd/cmd.exe
