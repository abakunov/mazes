# README

## Описание проекта

Этот проект представляет собой приложение для генерации и решения лабиринтов. Программа позволяет пользователю выбрать алгоритм генерации лабиринта, задать начальные и конечные точки, а также выбрать алгоритм поиска пути. После этого лабиринт генерируется и отображается в консоли, а также показывается найденный путь.

## Структура проекта

- **cmd/run/main.go**: Главный файл, который запускает приложение. Он обрабатывает ввод пользователя, инициализирует лабиринт, выбирает алгоритмы генерации и поиска пути, а также отображает результаты.
- **internal/application**: Содержит реализацию алгоритмов генерации и поиска пути.
    - `dfs_generator.go`: Реализация генерации лабиринта с использованием алгоритма поиска в глубину (DFS).
    - `kruskal_generator.go`: Реализация генерации лабиринта с использованием алгоритма Крускала.
    - `bfs_solver.go`: Реализация поиска пути с использованием алгоритма поиска в ширину (BFS).
    - `astar_solver.go`: Реализация поиска пути с использованием алгоритма A*.
- **internal/domain**: Содержит основные интерфейсы и модели данных.
    - `interfaces.go`: Интерфейсы для генерации и поиска пути.
    - `models.go`: Модели данных для представления точек, ячеек и лабиринта.
- **internal/infrastructure**: Содержит вспомогательные функции для ввода данных и отображения лабиринта.
    - `input_parser.go`: Функции для получения ввода от пользователя.
    - `console_renderer.go`: Функции для отображения лабиринта в консоли.

## Алгоритмы генерации лабиринтов

1. **DFS (поиск в глубину)**
2. **Алгоритм Крускала**

## Алгоритмы поиска пути

1. **BFS**
2. **A***

## Запуск кода

Для запуска приложения выполните следующую команду в терминале:

```bash
go run cmd/run/main.go
```

Или запустите проект через make
```bash
make run
```

## Запуск тестов

Для запуска тестов выполните следующую команду в терминале:

```bash
go test ./...
```

Или запустите тесты через make
```bash
make test 
```

Эта команда запустит все тесты, находящиеся в проекте, и выведет результаты в консоль.
