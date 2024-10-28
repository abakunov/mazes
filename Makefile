# Установка переменных
BINARY_NAME=labyrinths
BUILD_DIR=bin
MAIN_PATH=./cmd/run/main.go

# Целевая сборка проекта
.PHONY: build
build:
	@echo "===> Компиляция проекта..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "===> Проект скомпилирован в $(BUILD_DIR)/$(BINARY_NAME)"

# Целевая команда для запуска проекта
.PHONY: run
run: build
	@echo "===> Запуск проекта..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Линтеры (golangci-lint)
.PHONY: lint
lint:
	@echo "===> Запуск линтеров..."
	golangci-lint run

# Удаление собранных файлов
.PHONY: clean
clean:
	@echo "===> Очистка..."
	rm -rf $(BUILD_DIR)
	@echo "===> Очистка завершена"

# Покрытие тестов
.PHONY: test
test:
	@echo "===> Запуск тестов..."
	go test ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out
	@echo "===> Тестирование завершено"

