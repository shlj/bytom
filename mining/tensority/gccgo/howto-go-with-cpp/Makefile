.PHONY: clean

TARGET=tensority

$(TARGET): libfoo.a
	go build -ldflags "-linkmode external -extldflags -static" $(TARGET).go

# $(TARGET): foo.so
# 	go build -o $(TARGET)

libfoo.a: foo.o cfoo.o sha3.o
	ar rs $@ $^

%.o: %.cpp
	clang++ -o $@ -c $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

%.so: %.cpp
	clang++ -o $@ $^ -std=c++11 -pthread -mavx2 -O3 -fPIC

clean:
	rm -f *.o *.so *.a $(TARGET)
