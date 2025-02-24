package xlearn

/*
#cgo LDFLAGS: -L${SRCDIR} -lxlearn_api -Wl,-rpath,${SRCDIR}
#include "xlearn.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Model struct {
	handle C.XL
}

func Create(modelType string) (*Model, error) {
	var handle C.XL
	cModelType := C.CString(modelType)
	defer C.free(unsafe.Pointer(cModelType))

	ret := C.XLearnCreate(cModelType, &handle)
	if ret != 0 {
		return nil, errors.New("XLearnCreate failed")
	}
	return &Model{handle: handle}, nil
}

func (m *Model) SetTest(testPath string) error {
	if m.handle == nil {
		return errors.New("invalid xLearn handle")
	}
	cTestPath := C.CString(testPath)
	defer C.free(unsafe.Pointer(cTestPath))

	ret := C.XLearnSetTest(&m.handle, cTestPath)
	if ret != 0 {
		return errors.New("XLearnSetTest failed")
	}
	return nil
}

func (m *Model) PredictForMat(modelPath string) ([]float32, error) {
	if m.handle == nil {
		return nil, errors.New("invalid xLearn handle")
	}
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))

	var length C.uint64_t
	var cArr *C.float

	ret := C.XLearnPredictForMat(&m.handle, cModelPath, &length, &cArr)
	if ret != 0 {
		return nil, errors.New("XLearnPredictForMat failed")
	}

	n := int(length)

	cSlice := (*[1 << 30]C.float)(unsafe.Pointer(cArr))[:n:n]
	preds := make([]float32, n)
	for i := 0; i < n; i++ {
		preds[i] = float32(cSlice[i])
	}
	return preds, nil
}

func (m *Model) Free() error {
	if m.handle == nil {
		return nil
	}
	ret := C.XLearnHandleFree(&m.handle)
	if ret != 0 {
		return errors.New("XLearnHandleFree failed")
	}
	m.handle = nil
	return nil
}

func (m *Model) SetInt(key string, value int) error {
	if m.handle == nil {
		return errors.New("invalid xLearn handle")
	}
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	ret := C.XLearnSetInt(&m.handle, cKey, C.int(value))
	if ret != 0 {
		return errors.New("XLearnSetInt failed")
	}
	return nil
}

func (m *Model) GetInt(key string) (int, error) {
	if m.handle == nil {
		return 0, errors.New("invalid xLearn handle")
	}
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var cValue C.int
	ret := C.XLearnGetInt(&m.handle, cKey, &cValue)
	if ret != 0 {
		return 0, errors.New("XLearnGetInt failed")
	}
	return int(cValue), nil
}

func (m *Model) SetFloat(key string, value float32) error {
	if m.handle == nil {
		return errors.New("invalid xLearn handle")
	}
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	ret := C.XLearnSetFloat(&m.handle, cKey, C.float(value))
	if ret != 0 {
		return errors.New("XLearnSetFloat failed")
	}
	return nil
}

func (m *Model) GetFloat(key string) (float32, error) {
	if m.handle == nil {
		return 0, errors.New("invalid xLearn handle")
	}
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var cValue C.float
	ret := C.XLearnGetFloat(&m.handle, cKey, &cValue)
	if ret != 0 {
		return 0, errors.New("XLearnGetFloat failed")
	}
	return float32(cValue), nil
}

func (m *Model) SetBool(key string, value bool) error {
	if m.handle == nil {
		return errors.New("invalid xLearn handle")
	}
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.bool(value)
	ret := C.XLearnSetBool(&m.handle, cKey, cValue)
	if ret != 0 {
		return errors.New("XLearnSetBool failed")
	}
	return nil
}

func (m *Model) GetBool(key string) (bool, error) {
	if m.handle == nil {
		return false, errors.New("invalid xLearn handle")
	}
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var cValue C.bool
	ret := C.XLearnGetBool(&m.handle, cKey, &cValue)
	if ret != 0 {
		return false, errors.New("XLearnGetBool failed")
	}
	return bool(cValue), nil
}
