FRONTEND_DIR=frontend
BUILD_DIR=build
APP_NAME=http
APP_PORT=5050
NETWORK=bridge

clean:
	cd $(FRONTEND_DIR); \
	if [ -d $(BUILD_DIR) ]; then rm -rf $(BUILD_DIR); fi

static: clean
	cd $(FRONTEND_DIR); \
	pnpm i; \
	pnpm build

build: clean
	DOCKER_BUILDKIT=1 docker build \
	-t spa:$(APP_NAME) \
	--build-arg APP_NAME=$(APP_NAME) \
	.

run:
	docker run -dp $(APP_PORT):$(APP_PORT) \
	--network $(NETWORK) \
	--name $(APP_NAME) \
	--env APP_PORT=$(APP_PORT) \
	--restart unless-stopped \
	spa:$(APP_NAME)

stop:
	docker rm -f $(APP_NAME)

.PHONY: clean run build