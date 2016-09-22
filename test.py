# go build -o libdoubler.so -buildmode=c-shared .
#
import json
from ctypes import *
lib=CDLL("libdoubler.so")
lib.Free.argtypes = c_void_p,
lib.Free.restype = None

lib.HandleRequest.argtypes = c_char_p,
lib.HandleRequest.restype=c_void_p


je = json.dumps({"Operation":"double", "Value":2.0})
res = lib.HandleRequest(je)

print "ptr:", res
s = cast(res, c_char_p).value
print s
lib.Free(res)
#print(s)

