# 변수 정의
IMAGE_NAME := ghcr.io/nobreak-labs/go-myweb
TAG := v2.0
PLATFORMS := linux/amd64,linux/arm64
GITHUB_USERNAME ?= $(shell git config --get user.name)

# 도커 관련 타겟
.PHONY: build build-push ghcr-login clean

# 로컬 빌드
build:
	@echo "Building for platforms: $(PLATFORMS)"
	docker buildx build --platform $(PLATFORMS) -t $(IMAGE_NAME):$(TAG) --load .

# 이미지 푸시
build-push:
	@echo "Building and pushing for platforms: $(PLATFORMS)"
	docker buildx build --platform $(PLATFORMS) -t $(IMAGE_NAME):$(TAG) --push .

# GitHub Container Registry 로그인
ghcr-login:
	@if [ -z "$(CR_PAT)" ]; then \
		echo "Error: CR_PAT environment variable is not set"; \
		exit 1; \
	fi
	@echo "Logging in to GitHub Container Registry..."
	@echo $(CR_PAT) | docker login ghcr.io -u $(GITHUB_USERNAME) --password-stdin

# 로컬 도커 이미지 정리
clean:
	@echo "Cleaning local docker images..."
	docker rmi $(IMAGE_NAME):$(TAG) || true

# 도움말
help:
	@echo "Available targets:"
	@echo "  build        - Build multi-arch docker image locally"
	@echo "  build-push  - Build and push multi-arch docker image to registry"
	@echo "  ghcr-login  - Login to GitHub Container Registry (requires CR_PAT env var)"
	@echo "  clean       - Remove local docker images"
	@echo "  help        - Show this help message"
