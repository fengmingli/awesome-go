CONTROLLER_GEN = $(shell pwd)/bin/controller-gen

.PHONY: controller-gen
controller-gen:
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0)

.PHONY: manifests
manifests: controller-gen codegen
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd

.PHONY: generate
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: codegen
codegen:
	cd $(shell pwd)/hack && bash update-codegen.sh


.PHONY: dependency
dependency:
	go get k8s.io/apimachinery@v0.0.0-20190425132440-17f84483f500 && \
    go get k8s.io/client-go@v0.0.0-20190425172711-65184652c889 && \
    go get k8s.io/code-generator@v0.0.0-20190419212335-ff26e7842f9d


#========================================================================================================

.PHONY: all build clean run check cover lint docker help
BIN_FILE=hello
all: check build
build:
	@go build -o "${BIN_FILE}"
clean:
	@go clean
	@rm --force "xx.out"
test:
	@go test
check:
	@go fmt ./...
	@go vet ./
cover:
	@go test -coverprofile xx.out
	@go tool cover -html=xx.out
run:
	./"${BIN_FILE}"
lint:
	golangci-lint run --enable-all
docker:
	@docker build -t leo/hello:latest .
help:
	@echo "make 格式化够代码 并编译成二进制文件"
	@echo "make build 编译go代码生成二进制文件"
	@echo "make clean 清理中间目标文件"
	@echo "make test 执行测试case"
	@echo "make check 格式化go代码"
	@echo "make cover 检查测试覆盖率"
	@echo "make run 直接运行程序"
	@echo "make lint 执行代码检查"
	@echo "make docker 构建docker镜像"




