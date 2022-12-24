
.PHONY: test
test:
	for n in $$(ls services); do \
		(cd "services/$$n" && go test ./...); \
	done

.PHONY: sec-scan
sec-scan:
	for n in $$(ls services); do \
		(cd "services/$$n" && gosec -exclude=G404 ./...); \
	done


.PHONY: test-build
test-build:
	for n in $$(ls services); do \
		(cd "services/$$n" && pwd && go build ); \
	done

.PHONY: tidy
tidy:
	for f in $$(find . -name go.mod); do \ 
		(cd $$(dirname $f); go mod tidy); \
	done

.PHONY: docker-compose
docker-compose:
	docker-compose down --volumes --remove-orphans && \
  	docker-compose rm && \
  	docker-compose build && \
  	docker-compose up