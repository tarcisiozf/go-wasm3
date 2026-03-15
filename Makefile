dale:
	cp ../wasm4/source/*.h include && cp ../wasm4/build/source/libm3.a lib/darwin/libm3.a && go run examples/sum/sum.go && git add . && git commit -m "update lib" && git push