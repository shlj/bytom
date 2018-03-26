.PHONY: clean

TARGET=cSimdTs

LIBS=
UNAME := $(shell uname -s)
ifeq ($(UNAME),Darwin)
	LIBS += lib.dylib
else
	LIBS += lib.so
endif

$(TARGET).so: cSimdTs.cpp
	clang++ -shared -o $@ $^ -std=c++11 -pthread -mavx2 -O3 -fPIC
	cd .. && go test 

clean:
	rm -f *.o *.so *.a