.PHONY: clean

TARGET=tensority

$(TARGET): libfoo.a
	go build -o $(TARGET)

libfoo.a: foo.o cfoo.o
	ar r $@ $^

%.o: %.cpp
	clang++ -o $@ -c $^ -std=c++11 -pthread -mavx2 -O3

clean:
	rm -f *.o *.so *.a $(TARGET)
