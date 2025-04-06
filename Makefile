wasm:
	GOARCH=wasm GOOS=js go build --tags gmath -ldflags="-s -w" -trimpath -o _web/main.wasm ./cmd/game

itchio-wasm: wasm
	cd _web && \
		mkdir -p ../bin && \
		rm -f ../bin/ld57game.zip && \
		zip ../bin/ld57game.zip -r main.wasm index.html wasm_exec.js
