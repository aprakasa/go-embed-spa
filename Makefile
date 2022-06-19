FRONTEND_DIR=frontend
BUILD_DIR=build
APP_NAME=http
APP_PORT=5050

clean:
	cd $(FRONTEND_DIR); \
	if [ -d $(BUILD_DIR) ]; then rm -rf $(BUILD_DIR); fi

static: clean
	cd $(FRONTEND_DIR); \
	npm install; \
	npm run build

build: clean
	DOCKER_BUILDKIT=1 docker build -t spa:$(APP_NAME) --build-arg APP_NAME=$(APP_NAME) .

run:
	docker run -d -p $(APP_PORT):$(APP_PORT) --name spa-$(APP_NAME) -e APP_PORT=$(APP_PORT) --restart always spa:$(APP_NAME)

.PHONY: clean run build