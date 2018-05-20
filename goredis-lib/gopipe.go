package main

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
// int PyArg_ParseTuple_String(PyObject *, char**, char**, char**, char**);
// int PyArg_ParseTuple_Connection(PyObject *, char**, long long *);
// int PyArg_ParseTuple_LL(PyObject *, long long *);
// PyObject* Py_String(char *pystring);
import (
	"C"
)
import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/go-redis/redis"
)

var pipe redis.Pipeliner

//export Connect
func Connect(self *C.PyObject, args *C.PyObject) *C.PyObject {
	var ip *C.char
	runtime.GOMAXPROCS(4)
	go http.ListenAndServe("0.0.0.0:8080", nil)
	var port C.longlong
	if C.PyArg_ParseTuple_Connection(args, &ip, &port) == 0 {
		log.Println(0)
		return C.PyLong_FromLong(0)
	}
	ipStr := C.GoString(ip)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", ipStr, port),
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 4,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Panicln(err)
		return C.PyLong_FromLong(0)
	}

	pipe = client.Pipeline()
	return C.PyLong_FromLong(0)
}

//export add_command
func add_command(self *C.PyObject, args *C.PyObject) *C.PyObject {
	var cmd, hashmap, key, value *C.char
	if C.PyArg_ParseTuple_String(args, &cmd, &hashmap, &key, &value) == 0 {
		return C.PyLong_FromLong(0)
	}
	pipe.HSet(C.GoString(hashmap), C.GoString(key), C.GoString(value))
	return C.PyLong_FromLong(0)
}

//export execute
func execute(self *C.PyObject, args *C.PyObject) *C.PyObject {
	_, err := pipe.Exec()
	if err != nil {
		log.Println(err)
		return C.PyLong_FromLong(0)
	}
	err = pipe.Discard()
	if err != nil {
		log.Println(err)
	}
	return C.PyLong_FromLong(0)
}

func main() {}
