#define Py_LIMITED_API
#include <Python.h>

PyObject * add_command(PyObject *, PyObject*);
PyObject * execute(PyObject *, PyObject *);

int PyArg_ParseTuple_String(PyObject * args, char**cmd, char**hashmap_name, char **key, char **value) {
    return PyArg_ParseTuple(args, "ssss", cmd, hashmap_name, key, value);
}

PyObject* Py_String(char **pystring){
    return Py_BuildValue("s", pystring);
}

static PyMethodDef CpipeMethods[] = {
    {"add_command", add_command, METH_VARARGS, ""},
    {"execute", execute, METH_VARARGS, ""},
    {NULL, NULL, 0, NULL}
};

static struct PyModuleDef cpipemodule = {
   PyModuleDef_HEAD_INIT, "cpipe", NULL, -1, CpipeMethods
};

PyMODINIT_FUNC
PyInit_cpipe(void)
{
    return PyModule_Create(&cpipemodule);
}
