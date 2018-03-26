.PHONY: clean

TARGET=tensority

LIBS=
UNAME := $(shell uname -s)
ifeq ($(UNAME),Darwin)
	LIBS += lib.dylib
else
	LIBS += lib.so
endif

$(TARGET): 
	# clang++ -o loadlib.o -c loadlib.cpp -std=c++11 -pthread -mavx2 -O3 -fPIC
	# ar rs loadlib.a loadlib.o
	# go run -ldflags "-linkmode external -extldflags -static" so.go
	# clang++ -shared -o cSimdTs.so cSimdTs.cpp sha3.c byte_order.c -std=c++11 -pthread -mavx2 -O3 -fPIC
	clang++ -shared -o cSimdTs.so cSimdTs.cpp -std=c++11 -pthread -mavx2 -O3 -fPIC
	go run so.go

# $(TARGET): libts.a
# 	go build -buildmode=shared -ldflags "-linkmode external -extldflags -static" $(TARGET).go
# 	# go build -buildmode=shared -ldflags "-linkmode external -extldflags -dynamic" $(TARGET).go

# libts.a: cSimdTs.o sha3.o
	# ar rs $@ $^
	# sudo mv cSimdTs.so /usr/lib
	# sudo mv sha3.so /usr/lib
	# sudo mv libts.a /usr/lib

# libts.a: cSimdTs.so
# 	ar rs $@ $^
# 	sudo mv cSimdTs.so /usr/lib
# 	sudo mv libts.a /usr/lib
	
# %.o: %.cpp
# 	clang++ -o $@ -c $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

# %.so: %.cpp
	# clang++ -shared -o $@ $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

clean:
	rm -f *.o *.so *.a $(TARGET)
	# sudo rm /usr/lib/cSimdTs.so
	# sudo rm /usr/lib/cSimdTs.o
	# sudo rm /usr/lib/sha3.so
	# sudo rm /usr/lib/sha3.o
	# sudo rm /usr/lib/libts.a
