.PHONY: clean

TARGET=tensority

LIBS=
UNAME := $(shell uname -s)
ifeq ($(UNAME),Darwin)
	LIBS += lib.dylib
else
	LIBS += lib.so
endif

# $(TARGET): libts.a
$(TARGET): libts.a
	go build -ldflags "-linkmode external -extldflags -static" $(TARGET).go
	# go build -ldflags "-linkmode external -extldflags -dynamic" $(TARGET).go

libts.a: cSimdTs.o
	ar rs $@ $^
	# sudo mv cSimdTs.so /usr/lib
	# sudo mv sha3.so /usr/lib
	# sudo mv libts.a /usr/lib
	
%.o: %.cpp
	clang++ -o $@ -c $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

# %.so: %.cpp
	# clang++ -shared -o $@ $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

clean:
	rm -f *.o *.so *.a $(TARGET)
