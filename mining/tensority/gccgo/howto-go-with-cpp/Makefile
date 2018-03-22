.PHONY: clean

TARGET=tensority

$(TARGET): libts.a
	go build -ldflags "-linkmode external -extldflags -static" $(TARGET).go
	rm -f *.o *.so *.a

libts.a: tensority.o sha3.o
	ar rs $@ $^

%.o: %.cpp
	clang++ -o $@ -c $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

# $(TARGET): libts.a
# 	go build $(TARGET).go

# libts.a: tensority.so sha3.so
# 	ar rs $@ $^

# %.so: %.cpp
# 	clang++ -shared -o $@ $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

clean:
	rm -f *.o *.so *.a $(TARGET)
