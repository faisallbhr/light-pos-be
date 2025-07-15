# =========================
# ⚙️ ENVIRONMENT HANDLING
# =========================
ENV_FILE ?= .env.development

# Load env file
include $(ENV_FILE)
export $(shell sed -n 's/^\([^#][^=]*\)=.*/\1/p' $(ENV_FILE))

# =========================
# 📂 CONFIG
# =========================
APP_NAME     := light-pos-be
CMD_DIR      := ./cmd/server
BUILD_DIR    := ./bin

# =========================
# 🛠️ BUILD & RUN
# =========================
run:
	go run $(CMD_DIR)/main.go

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

clean:
	rm -rf $(BUILD_DIR)

# =========================
# 🗃️ GOOSE MIGRATION
# =========================
migrate-up:
	goose -dir $(GOOSE_MIGRATION_DIR) -table $(GOOSE_TABLE) \
		$(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" \
		$(if $(version),up-to $(version),up)

migrate-down:
	goose -dir $(GOOSE_MIGRATION_DIR) -table $(GOOSE_TABLE) \
		$(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" \
		$(if $(version),down-to $(version),down)

migrate-create:
	goose -dir $(GOOSE_MIGRATION_DIR) create $(name) sql

migrate-status:
	goose -dir $(GOOSE_MIGRATION_DIR) -table $(GOOSE_TABLE) \
		$(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" status

migrate-version:
	goose -dir $(GOOSE_MIGRATION_DIR) -table $(GOOSE_TABLE) \
		$(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" version

migrate-redo:
	goose -dir $(GOOSE_MIGRATION_DIR) -table $(GOOSE_TABLE) \
		$(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" redo

migrate-reset:
ifeq ($(ENV_FILE),.env.production)
	$(error ❌ ERROR: migrate-reset should not be run in production!)
endif
	goose -dir $(GOOSE_MIGRATION_DIR) -table $(GOOSE_TABLE) \
		$(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" reset

migrate-fresh:
ifeq ($(ENV_FILE),.env.production)
	$(error ❌ ERROR: migrate-fresh should not be run in production!)
endif
	goose -dir $(GOOSE_MIGRATION_DIR) -table $(GOOSE_TABLE) \
		$(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" reset && \
	goose -dir $(GOOSE_MIGRATION_DIR) -table $(GOOSE_TABLE) \
		$(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" up

# =========================
# 🌱 SEEDER
# =========================
seed:
	go run ./cmd/seed/main.go $(if $(only),--only=$(only))

# =========================
# 🔧 UTILITIES
# =========================
print-env:
	@echo "Using $(ENV_FILE)"
	@echo "GOOSE_DRIVER=$(GOOSE_DRIVER)"
	@echo "GOOSE_DBSTRING=$(GOOSE_DBSTRING)"
	@echo "GOOSE_TABLE=$(GOOSE_TABLE)"
	@echo "GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR)"

# =========================
# 📌 PHONY
# =========================
.PHONY: run build clean \
	migrate-up migrate-down migrate-create migrate-status migrate-version \
	migrate-redo migrate-reset migrate-fresh seed print-env
