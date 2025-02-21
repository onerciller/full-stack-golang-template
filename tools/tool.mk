#swagger tool
SWAG := $(LOCAL_BIN)/swag
SWAG_VERSION ?= v1.8.12

$(SWAG):
	@echo "$(BLUE)Installing swag-$(SWAG_VERSION)$(CNone)"
	@GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)



