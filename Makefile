.PHONY: build
build: 
	@#rm -rf public/*
	@hugo -b https://questapi.firebaseapp.com
.PHONY: deploy
deploy: build
	@firebase deploy
.PHONY: sandbox
sandbox: 
	@hugo -b /
verifier:
	@cd scripts/verifier && go run main.go