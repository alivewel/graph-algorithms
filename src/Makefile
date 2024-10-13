.PHONY: all clean s21_graph test run

all: s21_graph
	cd graphs && go build -o ./graph_app ./cmd/main.go

s21_graph:
	cd graphs && go build -buildmode=archive -o libs/s21_graph.a internal/s21_graph/*.go
	cd graphs && go build -buildmode=archive -o libs/s21_graph_algorithms.a internal/s21_graph_algorithms/*.go

test: s21_graph
	cd graphs && go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

clean:
	rm -f ./graphs/graph_app
	rm -f ./graphs/libs/s21_graph.a
	rm -f ./graphs/libs/s21_graph_algorithms.a
	rm -f ./graphs/data/*.dot
	rm -f ./graphs/coverage.out

run:
	./graphs/graph_app
