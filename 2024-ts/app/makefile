.PHONY: install dev backend frontend build test frontend-test lint mongo start all test-all

# Install dependencies
install:
	npm install
	cd frontend && npm install

# Development: backend + frontend
dev:
	@trap 'kill 0' INT TERM EXIT; \
	npm run server & \
	cd frontend && npm run dev & \
	wait

# Backend only
backend:
	npm run server

# Frontend only
frontend:
	cd frontend && npm run dev

# Build React production bundle
build:
	cd frontend && npm run build

# Start MongoDB
mongo:
	npm run mongo

# Backend tests
test:
	npm test

# Frontend tests
frontend-test:
	cd frontend && npm run test

# Frontend lint
lint:
	cd frontend && npm run lint

# Build + MongoDB + production server
start:
	cd frontend && npm run build
	npm run mongo
	npm run server

# Everything needed for production
all: build mongo
	npm run server

# Run backend + frontend tests
test-all: test frontend-test