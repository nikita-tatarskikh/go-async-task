lint:
	@echo ">${YELLOW} Linting source files...${NC}"
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix
	@echo ">${GREEN} Source files fine${NC}"

build:
	@echo ">${YELLOW} Building binary...${NC}"
	go build main.go
	@echo ">${YELLOW} Binary is built${NC}"

run:
	@echo ">${YELLOW} Running binary...${NC}"
	go run main.go
	@echo ">${YELLOW} Binary was executed${NC}"

