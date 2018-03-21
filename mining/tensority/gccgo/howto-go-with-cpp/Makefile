.PHONY: clean

TARGET=tensority

$(TARGET): libfoo.a
	go build -o $(TARGET)

libfoo.a: foo.o cfoo.o
	ar r $@ $^

%.o: %.cpp
	clang++ -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)
