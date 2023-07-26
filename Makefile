
run-tailwind-cli:
	npx tailwindcss -i ./src/style.css -o ./dist/output.css --watch

watch-run:
	nodemon --watch './*' --watch './**/*' -e go,js,html --delay 1s --signal SIGTERM --exec 'go' run ./main.go

build:
	go build -o bin/ssp-webapp ./cmd/main.go

build-release:
	go build -ldflags="-s -w" -o bin/ssp-webapp ./cmd/main.go

