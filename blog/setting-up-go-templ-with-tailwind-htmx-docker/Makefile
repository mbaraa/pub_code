.PHONY: build

BINARY_NAME=spendings

build:
	cd tailwindcss && \
	npx tailwindcss build -i ../static/css/style.css -o ../static/css/tailwind.css -m && \
	cd .. && \
	go mod tidy && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}


dev:
	templ generate --watch --cmd="go run ." & \
	cd tailwindcss && \
	npx tailwindcss build -i ../static/css/style.css -o ../static/css/tailwind.css --watch

clean:
	go clean

