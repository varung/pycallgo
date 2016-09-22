# go build -o libdoubler.so -buildmode=c-shared .
#
import ctypes
lib=ctypes.CDLL("libdoubler.so")
print(lib.DoubleIt(21))

for i in range(100):
    print(lib.Count())
