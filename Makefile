PATCHES:= $(shell find assets/patches/)

bindata: $(PATCHES) 
	cd assets && go-bindata -o ../rump/bindata/patches.go -pkg bindata patches/
