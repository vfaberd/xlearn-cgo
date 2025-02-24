#ifndef XLEARN_C_API_C_API_H_
#define XLEARN_C_API_C_API_H_

#ifdef __cplusplus
#define XL_EXTERN_C extern "C"
#include <cstdio>
#else
#define XL_EXTERN_C
#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include <string.h>
#endif

#if defined(_MSC_VER) || defined(_WIN32)
#define XL_DLL XL_EXTERN_C __declspec(dllexport)
#else
#define XL_DLL XL_EXTERN_C
#endif

/* Handle to xlearn */
typedef void* XL;
typedef void* DataHandle;

// Say hello to user
XL_DLL int XLearnHello();

// Create xlearn handle
XL_DLL int XLearnCreate(const char *model_type, XL *out);

XL_DLL int XLearnShow(XL *out);

// Free the xLearn handle
XL_DLL int XLearnHandleFree(XL *out);

// Set file path of the test data
XL_DLL int XLearnSetTest(XL *out, const char *test_path);

// Start to predict, this function is for output numpy
XL_DLL int XLearnPredictForMat(XL *out, const char *model_path, 
                               uint64_t *length, const float** out_arr);

// Set string param
XL_DLL int XLearnSetStr(XL *out, const char *key, const char *value);

// Set int param
XL_DLL int XLearnSetInt(XL *out, const char *key, const int value);

// Get int param
XL_DLL int XLearnGetInt(XL *out, const char *key, int *value);

// Set float param
XL_DLL int XLearnSetFloat(XL *out, const char *key, const float value);

// Get float param
XL_DLL int XLearnGetFloat(XL *out, const char *key, float *value);

// Set bool param
XL_DLL int XLearnSetBool(XL *out, const char *key, const bool value);

// Get bool param
XL_DLL int XLearnGetBool(XL *out, const char *key, bool *value);

#endif  // XLEARN_C_API_C_API_H_

