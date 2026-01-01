include .env

LOCAL_BIN := $(CURDIR)/bin

# ========================
# DEPS
# ========================

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

# ========================
# MIGRATIONS (MAIN DB)
# ========================

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} down -v

# ========================
# MIGRATIONS (ORDERS DB)
# ========================

local-migration-orders-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_ORDERS_DIR} postgres ${PG_ORDERS_DSN} status -v

local-migration-orders-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_ORDERS_DIR} postgres ${PG_ORDERS_DSN} up -v

local-migration-orders-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_ORDERS_DIR} postgres ${PG_ORDERS_DSN} down -v
