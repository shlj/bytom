.PHONY: clean

TARGET=tensority

$(TARGET): libts.a
	go build -ldflags "-linkmode external -extldflags -static" $(TARGET).go
	rm -f *.o *.so

libts.a: cSimdTs.o sha3.o
	ar rs $@ $^

%.o: %.cpp
	clang++ -o $@ -c $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

clean:
	rm -f *.o *.so *.a $(TARGET)
