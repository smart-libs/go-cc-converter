
GO=go
GO_SRC_FOLDER=pkg

#lint: @ run golint to check code style
lint:
	@echo ====[GOLINT]=============================================================================
	@echo ====[Installing staticcheck@latest]======================================================
	@$(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	@echo ====[Performing static checks]===========================================================
	@staticcheck ./...

#run-tests: @ convenient task to run tests without generate code or start/stop db
run-tests:
	@echo "[VET]========================================="
	@$(GO) vet ./$(GO_SRC_FOLDER)/...
	@mkdir -p ./coverage
	@echo "[Test]========================================"
	@APP_ENV=dev $(GO) test -v ./$(GO_SRC_FOLDER)/... --cover -coverpkg ./pkg/... -coverprofile coverage/coverage.fmt fmt
	@APP_ENV=dev $(GO) tool cover -html=coverage/coverage.fmt -o coverage/coverage.html

#test: @ convenient task to generate code, mocks and run tests
test: run-tests lint
