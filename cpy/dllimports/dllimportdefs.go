
package dllimports

var dllImportDefs = []importTable{
    /*0*/ {"_Py_NoneStruct", 0},
    /*1*/ {"Py_Initialize", 0},
    /*2*/ {"Py_InitializeEx", 0},
    /*3*/ {"Py_IsInitialized", 0},
    /*4*/ {"Py_Finalize", 0},
    /*5*/ {"Py_FinalizeEx", 0},
    /*6*/ {"Py_EncodeLocale", 0},
    /*7*/ {"Py_DecodeLocale", 0},
    /*8*/ {"PyMem_Free", 0},
    /*9*/ {"PyMem_RawFree", 0},
    /*10*/ {"Py_SetProgramName", 0},
    /*11*/ {"Py_GetProgramName", 0},
    /*12*/ {"Py_Main", 0},
    /*13*/ {"Py_BytesMain", 0},
    /*14*/ {"PyRun_AnyFile", 0},
    /*15*/ {"PyRun_SimpleFile", 0},
    /*16*/ {"PyRun_SimpleString", 0},
    /*17*/ {"Py_SetPath", 0},
    /*18*/ {"Py_SetPythonHome", 0},
    /*19*/ {"_Py_fopen_obj", 0},
    /*20*/ {"PyImport_AppendInittab", 0},
    /*21*/ {"Py_IncRef", 0},
    /*22*/ {"Py_DecRef", 0},
    /*23*/ {"PyUnicode_Type", 0},
    /*24*/ {"PyUnicode_FromString", 0},
    /*25*/ {"PyUnicode_GetLength", 0},
    /*26*/ {"PyUnicode_FromFormat", 0},
    /*27*/ {"PyUnicode_Fill", 0},
    /*28*/ {"PyUnicode_WriteChar", 0},
    /*29*/ {"PyUnicode_ReadChar", 0},
    /*30*/ {"PyUnicode_Substring", 0},
    /*31*/ {"PyUnicode_AsUCS4", 0},
    /*32*/ {"PyUnicode_DecodeLocaleAndSize", 0},
    /*33*/ {"PyUnicode_DecodeLocale", 0},
    /*34*/ {"PyUnicode_EncodeLocale", 0},
    /*35*/ {"PyUnicode_FSConverter", 0},
    /*36*/ {"PyUnicode_FSDecoder", 0},
    /*37*/ {"PyUnicode_DecodeFSDefaultAndSize", 0},
    /*38*/ {"PyUnicode_DecodeFSDefault", 0},
    /*39*/ {"PyUnicode_EncodeFSDefault", 0},
    /*40*/ {"PyUnicode_FromWideChar", 0},
    /*41*/ {"PyUnicode_AsWideChar", 0},
    /*42*/ {"PyUnicode_AsWideCharString", 0},
    /*43*/ {"PyUnicode_Decode", 0},
    /*44*/ {"PyUnicode_AsEncodedString", 0},
    /*45*/ {"PyUnicode_Encode", 0},
    /*46*/ {"PyUnicode_DecodeUTF8", 0},
    /*47*/ {"PyUnicode_DecodeUTF8Stateful", 0},
    /*48*/ {"PyUnicode_AsUTF8String", 0},
    /*49*/ {"PyUnicode_AsUTF8AndSize", 0},
    /*50*/ {"PyUnicode_AsUTF8", 0},
    /*51*/ {"PyUnicode_EncodeUTF8", 0},
    /*52*/ {"PyUnicode_DecodeUTF32", 0},
    /*53*/ {"PyUnicode_DecodeUTF32Stateful", 0},
    /*54*/ {"PyUnicode_AsUTF32String", 0},
    /*55*/ {"PyUnicode_EncodeUTF32", 0},
    /*56*/ {"PyUnicode_DecodeUTF16", 0},
    /*57*/ {"PyUnicode_DecodeUTF16Stateful", 0},
    /*58*/ {"PyUnicode_AsUTF16String", 0},
    /*59*/ {"PyUnicode_EncodeUTF16", 0},
    /*60*/ {"PyUnicode_DecodeUTF7", 0},
    /*61*/ {"PyUnicode_DecodeUTF7Stateful", 0},
    /*62*/ {"PyUnicode_EncodeUTF7", 0},
    /*63*/ {"PyUnicode_DecodeUnicodeEscape", 0},
    /*64*/ {"PyUnicode_AsUnicodeEscapeString", 0},
    /*65*/ {"PyUnicode_EncodeUnicodeEscape", 0},
    /*66*/ {"PyUnicode_DecodeRawUnicodeEscape", 0},
    /*67*/ {"PyUnicode_AsRawUnicodeEscapeString", 0},
    /*68*/ {"PyUnicode_EncodeRawUnicodeEscape", 0},
    /*69*/ {"PyUnicode_DecodeASCII", 0},
    /*70*/ {"PyUnicode_AsASCIIString", 0},
    /*71*/ {"PyUnicode_EncodeASCII", 0},
    /*72*/ {"PyUnicode_Concat", 0},
    /*73*/ {"PyUnicode_Split", 0},
    /*74*/ {"PyUnicode_Splitlines", 0},
    /*75*/ {"PyUnicode_Join", 0},
    /*76*/ {"PyUnicode_Tailmatch", 0},
    /*77*/ {"PyUnicode_Find", 0},
    /*78*/ {"PyUnicode_FindChar", 0},
    /*79*/ {"PyUnicode_Count", 0},
    /*80*/ {"PyUnicode_Replace", 0},
    /*81*/ {"PyUnicode_Compare", 0},
    /*82*/ {"PyUnicode_CompareWithASCIIString", 0},
    /*83*/ {"PyUnicode_RichCompare", 0},
    /*84*/ {"PyUnicode_Format", 0},
    /*85*/ {"PyUnicode_Contains", 0},
    /*86*/ {"PyUnicode_InternInPlace", 0},
    /*87*/ {"PyUnicode_InternFromString", 0},
    /*88*/ {"PyObject_Init", 0},
    /*89*/ {"PyObject_InitVar", 0},
    /*90*/ {"_PyObject_New", 0},
    /*91*/ {"_PyObject_NewVar", 0},
    /*92*/ {"PyObject_Del", 0},
    /*93*/ {"PyObject_Length", 0},
    /*94*/ {"PyObject_Str", 0},
    /*95*/ {"PyObject_Bytes", 0},
    /*96*/ {"PyObject_IsSubclass", 0},
    /*97*/ {"PyObject_IsInstance", 0},
    /*98*/ {"PyObject_GetAttr", 0},
    /*99*/ {"PyObject_Call", 0},
    /*100*/ {"PyObject_CallObject", 0},
    /*101*/ {"PyObject_CallNoArgs", 0},
    /*102*/ {"PyObject_ASCII", 0},
    /*103*/ {"PyObject_Hash", 0},
    /*104*/ {"PyObject_IsTrue", 0},
    /*105*/ {"PyObject_Not", 0},
    /*106*/ {"PyObject_Type", 0},
    /*107*/ {"PyObject_HasAttr", 0},
    /*108*/ {"PyObject_Print", 0},
    /*109*/ {"PyObject_HasAttrString", 0},
    /*110*/ {"PyObject_GetAttrString", 0},
    /*111*/ {"PyObject_GenericGetAttr", 0},
    /*112*/ {"PyObject_SetAttr", 0},
    /*113*/ {"PyObject_SetAttrString", 0},
    /*114*/ {"PyObject_DelAttr", 0},
    /*115*/ {"PyObject_DelAttrString", 0},
    /*116*/ {"PyObject_GenericGetDict", 0},
    /*117*/ {"PyObject_GenericSetDict", 0},
    /*118*/ {"PyObject_CallFunction", 0},
    /*119*/ {"PyObject_CallFunctionObjArgs", 0},
    /*120*/ {"PyObject_CallMethod", 0},
    /*121*/ {"PyObject_LengthHint", 0},
    /*122*/ {"PyObject_GetItem", 0},
    /*123*/ {"PyObject_SetItem", 0},
    /*124*/ {"PyObject_DelItem", 0},
    /*125*/ {"PyObject_TypeCheck", 0},
    /*126*/ {"PyObject_Dir", 0},
    /*127*/ {"PyObject_GetIter", 0},
    /*128*/ {"PyLong_Type", 0},
    /*129*/ {"PyLong_Check", 0},
    /*130*/ {"PyLong_CheckExact", 0},
    /*131*/ {"PyLong_AsDouble", 0},
    /*132*/ {"PyLong_AsVoidPtr", 0},
    /*133*/ {"PyLong_AsLong", 0},
    /*134*/ {"PyLong_AsLongLong", 0},
    /*135*/ {"PyLong_AsLongAndOverflow", 0},
    /*136*/ {"PyLong_FromDouble", 0},
    /*137*/ {"PyLong_FromLong", 0},
    /*138*/ {"PyLong_FromLongLong", 0},
    /*139*/ {"PyLong_FromUnsignedLong", 0},
    /*140*/ {"PyLong_FromUnsignedLongLong", 0},
    /*141*/ {"PyLong_AsUnsignedLong", 0},
    /*142*/ {"PyLong_AsSize_t", 0},
    /*143*/ {"PyLong_AsUnsignedLongLong", 0},
    /*144*/ {"PyLong_AsUnsignedLongMask", 0},
    /*145*/ {"PyLong_AsUnsignedLongLongMask", 0},
    /*146*/ {"PyLong_AsLongLongAndOverflow", 0},
    /*147*/ {"PyLong_AsSsize_t", 0},
    /*148*/ {"PyLong_FromSsize_t", 0},
    /*149*/ {"PyLong_FromSize_t", 0},
    /*150*/ {"PyLong_FromString", 0},
    /*151*/ {"PyLong_FromUnicodeObject", 0},
    /*152*/ {"PyLong_FromVoidPtr", 0},
    /*153*/ {"PyImport_Import", 0},
    /*154*/ {"PyModule_Type", 0},
    /*155*/ {"PyModule_Check", 0},
    /*156*/ {"PyModule_CheckExact", 0},
    /*157*/ {"PyModule_NewObject", 0},
    /*158*/ {"PyModule_New", 0},
    /*159*/ {"PyModule_Create2", 0},
    /*160*/ {"PyModule_GetName", 0},
    /*161*/ {"PyModule_GetDict", 0},
    /*162*/ {"PyModule_GetState", 0},
    /*163*/ {"PyModule_GetDef", 0},
    /*164*/ {"PyModule_GetFilenameObject", 0},
    /*165*/ {"PyModule_GetFilename", 0},
    /*166*/ {"PyModule_GetNameObject", 0},
    /*167*/ {"PyModule_AddFunctions", 0},
    /*168*/ {"PyModule_AddIntConstant", 0},
    /*169*/ {"PyModule_AddStringConstant", 0},
    /*170*/ {"PyModule_AddIntMacro", 0},
    /*171*/ {"PyModule_AddStringMacro", 0},
    /*172*/ {"PyModule_AddType", 0},
    /*173*/ {"PyModule_AddObject", 0},
    /*174*/ {"PyModule_AddObjectRef", 0},
    /*175*/ {"PyModuleDef_Init", 0},
    /*176*/ {"create_module", 0},
    /*177*/ {"exec_module", 0},
    /*178*/ {"PyModule_FromDefAndSpec", 0},
    /*179*/ {"PyModule_FromDefAndSpec2", 0},
    /*180*/ {"PyModule_ExecDef", 0},
    /*181*/ {"PyModule_SetDocString", 0},
    /*182*/ {"PyDict_Type", 0},
    /*183*/ {"PyDict_Check", 0},
    /*184*/ {"PyDict_CheckExact", 0},
    /*185*/ {"PyDict_Contains", 0},
    /*186*/ {"PyDict_Copy", 0},
    /*187*/ {"PyDict_SetItem", 0},
    /*188*/ {"PyDict_SetItemString", 0},
    /*189*/ {"PyDict_New", 0},
    /*190*/ {"PyDict_Size", 0},
    /*191*/ {"PyDict_Clear", 0},
    /*192*/ {"PyDict_GetItem", 0},
    /*193*/ {"PyDict_Keys", 0},
    /*194*/ {"PyDict_GetItemString", 0},
    /*195*/ {"PyDict_GetItemWithError", 0},
    /*196*/ {"PyDict_DelItem", 0},
    /*197*/ {"PyDict_DelItemString", 0},
    /*198*/ {"PyDict_SetDefault", 0},
    /*199*/ {"PyDict_Items", 0},
    /*200*/ {"PyDict_Values", 0},
    /*201*/ {"PyDict_Next", 0},
    /*202*/ {"PyDict_Merge", 0},
    /*203*/ {"PyDict_Update", 0},
    /*204*/ {"PyDict_MergeFromSeq2", 0},
    /*205*/ {"PyDict_ClearFreeList", 0},
    /*206*/ {"PyEval_EvalCode", 0},
    /*207*/ {"PyEval_EvalCodeEx", 0},
    /*208*/ {"PyEval_EvalFrame", 0},
    /*209*/ {"PyEval_EvalFrameEx", 0},
    /*210*/ {"PyEval_MergeCompilerFlags", 0},
    /*211*/ {"PyEval_InitThreads", 0},
    /*212*/ {"PyEval_ThreadsInitialized", 0},
    /*213*/ {"PyEval_SaveThread", 0},
    /*214*/ {"PyEval_RestoreThread", 0},
    /*215*/ {"PyEval_GetBuiltins", 0},
    /*216*/ {"PyEval_GetLocals", 0},
    /*217*/ {"PyEval_GetGlobals", 0},
    /*218*/ {"PyEval_GetFrame", 0},
    /*219*/ {"PyEval_SetProfile", 0},
    /*220*/ {"PyFrame_GetBack", 0},
    /*221*/ {"PyFrame_GetCode", 0},
    /*222*/ {"PyFrame_GetLineNumber", 0},
    /*223*/ {"PyTuple_Type", 0},
    /*224*/ {"PyTuple_CheckExact", 0},
    /*225*/ {"PyTuple_Size", 0},
    /*226*/ {"PyTuple_GetItem", 0},
    /*227*/ {"PyTuple_SetItem", 0},
    /*228*/ {"PyTuple_GetSlice", 0},
    /*229*/ {"PyTuple_Check", 0},
    /*230*/ {"PyTuple_New", 0},
    /*231*/ {"_PyTuple_Resize", 0},
    /*232*/ {"PyTuple_ClearFreeList", 0},
    /*233*/ {"PyTuple_Pack", 0},
    /*234*/ {"PyBytes_Type", 0},
    /*235*/ {"PyBytes_Check", 0},
    /*236*/ {"PyBytes_CheckExact", 0},
    /*237*/ {"PyBytes_FromString", 0},
    /*238*/ {"PyBytes_FromStringAndSize", 0},
    /*239*/ {"PyBytes_FromFormat", 0},
    /*240*/ {"PyBytes_FromFormatV", 0},
    /*241*/ {"PyBytes_FromObject", 0},
    /*242*/ {"PyBytes_Size", 0},
    /*243*/ {"PyBytes_AsString", 0},
    /*244*/ {"PyBytes_AsStringAndSize", 0},
    /*245*/ {"PyBytes_Concat", 0},
    /*246*/ {"PyBytes_ConcatAndDel", 0},
    /*247*/ {"_PyBytes_Resize", 0},
    /*248*/ {"PyType_Type", 0},
    /*249*/ {"_PyType_Name", 0},
    /*250*/ {"PyType_GetName", 0},
    /*251*/ {"PyType_GetModule", 0},
    /*252*/ {"PyType_Check", 0},
    /*253*/ {"PyType_CheckExact", 0},
    /*254*/ {"PyType_ClearCache", 0},
    /*255*/ {"PyType_GetFlags", 0},
    /*256*/ {"PyType_Modified", 0},
    /*257*/ {"PyType_HasFeature", 0},
    /*258*/ {"PyType_IS_GC", 0},
    /*259*/ {"PyType_IsSubtype", 0},
    /*260*/ {"PyType_GenericAlloc", 0},
    /*261*/ {"PyType_GenericNew", 0},
    /*262*/ {"PyType_Ready", 0},
    /*263*/ {"PyType_GetSlot", 0},
    /*264*/ {"PyType_FromSpecWithBases", 0},
    /*265*/ {"PyType_FromSpec", 0},
    /*266*/ {"PyList_Type", 0},
    /*267*/ {"PyList_Check", 0},
    /*268*/ {"PyList_CheckExact", 0},
    /*269*/ {"PyList_New", 0},
    /*270*/ {"PyList_SetItem", 0},
    /*271*/ {"PyList_GetItem", 0},
    /*272*/ {"PyList_GetSlice", 0},
    /*273*/ {"PyList_SetSlice", 0},
    /*274*/ {"PyList_Size", 0},
    /*275*/ {"PyList_Insert", 0},
    /*276*/ {"PyList_Append", 0},
    /*277*/ {"PyList_Sort", 0},
    /*278*/ {"PyList_Reverse", 0},
    /*279*/ {"PyList_AsTuple", 0},
    /*280*/ {"PyBool_Type", 0},
    /*281*/ {"PyBool_FromLong", 0},
    /*282*/ {"PyBool_Check", 0},
    /*283*/ {"PyExc_Exception", 0},
    /*284*/ {"PyExc_ValueError", 0},
    /*285*/ {"PyErr_Clear", 0},
    /*286*/ {"PyErr_PrintEx", 0},
    /*287*/ {"PyErr_Print", 0},
    /*288*/ {"PyErr_WriteUnraisable", 0},
    /*289*/ {"PyErr_SetString", 0},
    /*290*/ {"PyErr_SetObject", 0},
    /*291*/ {"PyErr_Format", 0},
    /*292*/ {"PyErr_FormatV", 0},
    /*293*/ {"PyErr_SetNone", 0},
    /*294*/ {"PyErr_BadArgument", 0},
    /*295*/ {"PyErr_NoMemory", 0},
    /*296*/ {"PyErr_SetFromErrno", 0},
    /*297*/ {"PyErr_SetFromErrnoWithFilenameObject", 0},
    /*298*/ {"PyErr_SetFromErrnoWithFilenameObjects", 0},
    /*299*/ {"PyErr_SetFromErrnoWithFilename", 0},
    /*300*/ {"PyErr_SetFromWindowsErr", 0},
    /*301*/ {"PyErr_SetExcFromWindowsErr", 0},
    /*302*/ {"PyErr_SetFromWindowsErrWithFilename", 0},
    /*303*/ {"PyErr_SetExcFromWindowsErrWithFilenameObject", 0},
    /*304*/ {"PyErr_SetExcFromWindowsErrWithFilenameObjects", 0},
    /*305*/ {"PyErr_SetExcFromWindowsErrWithFilename", 0},
    /*306*/ {"PyErr_SetImportError", 0},
    /*307*/ {"PyErr_SyntaxLocationObject", 0},
    /*308*/ {"PyErr_SyntaxLocationEx", 0},
    /*309*/ {"PyErr_SyntaxLocation", 0},
    /*310*/ {"PyErr_BadInternalCall", 0},
    /*311*/ {"PyErr_WarnEx", 0},
    /*312*/ {"PyErr_SetImportErrorSubclass", 0},
    /*313*/ {"PyErr_WarnExplicitObject", 0},
    /*314*/ {"PyErr_WarnExplicit", 0},
    /*315*/ {"PyErr_WarnFormat", 0},
    /*316*/ {"PyErr_ResourceWarning", 0},
    /*317*/ {"PyErr_Occurred", 0},
    /*318*/ {"PyErr_ExceptionMatches", 0},
    /*319*/ {"PyErr_GivenExceptionMatches", 0},
    /*320*/ {"PyErr_Fetch", 0},
    /*321*/ {"PyErr_Restore", 0},
    /*322*/ {"PyErr_NormalizeException", 0},
    /*323*/ {"PyErr_GetExcInfo", 0},
    /*324*/ {"PyErr_SetExcInfo", 0},
    /*325*/ {"PyErr_CheckSignals", 0},
    /*326*/ {"PyErr_SetInterrupt", 0},
    /*327*/ {"PySignal_SetWakeupFd", 0},
    /*328*/ {"PyErr_NewException", 0},
    /*329*/ {"PyErr_NewExceptionWithDoc", 0},
    /*330*/ {"PyException_GetTraceback", 0},
    /*331*/ {"PyException_SetTraceback", 0},
    /*332*/ {"PyException_GetContext", 0},
    /*333*/ {"PyException_SetContext", 0},
    /*334*/ {"PyException_GetCause", 0},
    /*335*/ {"PyException_SetCause", 0},
    /*336*/ {"PyUnicodeDecodeError_Create", 0},
    /*337*/ {"PyUnicodeEncodeError_Create", 0},
    /*338*/ {"PyUnicodeTranslateError_Create", 0},
    /*339*/ {"PyUnicodeDecodeError_GetEncoding", 0},
    /*340*/ {"PyUnicodeEncodeError_GetEncoding", 0},
    /*341*/ {"PyUnicodeDecodeError_GetObject", 0},
    /*342*/ {"PyUnicodeEncodeError_GetObject", 0},
    /*343*/ {"PyUnicodeTranslateError_GetObject", 0},
    /*344*/ {"PyUnicodeDecodeError_GetStart", 0},
    /*345*/ {"PyUnicodeEncodeError_GetStart", 0},
    /*346*/ {"PyUnicodeTranslateError_GetStart", 0},
    /*347*/ {"PyUnicodeDecodeError_SetStart", 0},
    /*348*/ {"PyUnicodeEncodeError_SetStart", 0},
    /*349*/ {"PyUnicodeTranslateError_SetStart", 0},
    /*350*/ {"PyUnicodeDecodeError_GetEnd", 0},
    /*351*/ {"PyUnicodeEncodeError_GetEnd", 0},
    /*352*/ {"PyUnicodeTranslateError_GetEnd", 0},
    /*353*/ {"PyUnicodeDecodeError_SetEnd", 0},
    /*354*/ {"PyUnicodeEncodeError_SetEnd", 0},
    /*355*/ {"PyUnicodeTranslateError_SetEnd", 0},
    /*356*/ {"PyUnicodeDecodeError_GetReason", 0},
    /*357*/ {"PyUnicodeEncodeError_GetReason", 0},
    /*358*/ {"PyUnicodeTranslateError_GetReason", 0},
    /*359*/ {"PyUnicodeDecodeError_SetReason", 0},
    /*360*/ {"PyUnicodeEncodeError_SetReason", 0},
    /*361*/ {"PyUnicodeTranslateError_SetReason", 0},
    /*362*/ {"PyInstanceMethod_Type", 0},
    /*363*/ {"PyInstanceMethod_New", 0},
    /*364*/ {"PyInstanceMethod_Function", 0},
    /*365*/ {"PyCFunction_New", 0},
    /*366*/ {"Py_GetPath", 0},
    /*367*/ {"Py_GetPythonHome", 0},
    /*368*/ {"_py_fopen_obj", 0},
    /*369*/ {"Py_BuildValue", 0},
    /*370*/ {"Py_GetPlatform", 0},
    /*371*/ {"Py_GetVersion", 0},
    /*372*/ {"Py_GetCompiler", 0},
    /*373*/ {"Py_GetBuildInfo", 0},
    /*374*/ {"Py_GetExecPrefix", 0},
    /*375*/ {"Py_GetProgramFullPath", 0},
    /*376*/ {"Py_GetPrefix", 0},
    /*377*/ {"Py_GetCopyright", 0},
    /*378*/ {"Py_SetStandardStreamEncoding", 0},
    /*379*/ {"Py_NewInterpreter", 0},
    /*380*/ {"Py_EndInterpreter", 0},
    /*381*/ {"Py_AddPendingCall", 0},
    /*382*/ {"Py_tracefunc", 0},
    /*383*/ {"Py_CompileString", 0},
    /*384*/ {"Py_CompileStringFlags", 0},
    /*385*/ {"Py_CompileStringObject", 0},
    /*386*/ {"Py_CompileStringExFlags", 0},
    /*387*/ {"Py_eval_input", 0},
    /*388*/ {"Py_file_input", 0},
    /*389*/ {"Py_single_input", 0},
    /*390*/ {"Py_EnterRecursiveCall", 0},
    /*391*/ {"Py_ReprEnter", 0},
    /*392*/ {"Py_ReprLeave", 0},
    /*393*/ {"Py_FatalError", 0},
    /*394*/ {"Py_Exit", 0},
    /*395*/ {"Py_AtExit", 0},
    /*396*/ {"Py_ExitStatusException", 0},
    /*397*/ {"Py_FdIsInteractive", 0},
    /*398*/ {"Py_InitializeFromConfig", 0},
    /*399*/ {"Py_PreInitialize", 0},
    /*400*/ {"Py_PreInitializeFromArgs", 0},
    /*401*/ {"Py_PreInitializeFromBytesArgs", 0},
    /*402*/ {"Py_RunMain", 0},
    /*403*/ {"Py_LeaveRecursiveCall", 0},
    /*404*/ {"PyMethod_Type", 0},
    /*405*/ {"PyMethod_Check", 0},
    /*406*/ {"PyMethod_New", 0},
    /*407*/ {"PyMethod_Function", 0},
    /*408*/ {"PyMethod_Self", 0},
    /*409*/ {"PyState_FindModule", 0},
    /*410*/ {"PyState_AddModule", 0},
    /*411*/ {"PyState_RemoveModule", 0},
}
const (
    Py_NoneStruct = 0
    Py_Initialize = 1
    Py_InitializeEx = 2
    Py_IsInitialized = 3
    Py_Finalize = 4
    Py_FinalizeEx = 5
    Py_EncodeLocale = 6
    Py_DecodeLocale = 7
    PyMem_Free = 8
    PyMem_RawFree = 9
    Py_SetProgramName = 10
    Py_GetProgramName = 11
    Py_Main = 12
    Py_BytesMain = 13
    PyRun_AnyFile = 14
    PyRun_SimpleFile = 15
    PyRun_SimpleString = 16
    Py_SetPath = 17
    Py_SetPythonHome = 18
    Py_fopen_obj = 19
    PyImport_AppendInittab = 20
    Py_IncRef = 21
    Py_DecRef = 22
    PyUnicode_Type = 23
    PyUnicode_FromString = 24
    PyUnicode_GetLength = 25
    PyUnicode_FromFormat = 26
    PyUnicode_Fill = 27
    PyUnicode_WriteChar = 28
    PyUnicode_ReadChar = 29
    PyUnicode_Substring = 30
    PyUnicode_AsUCS4 = 31
    PyUnicode_DecodeLocaleAndSize = 32
    PyUnicode_DecodeLocale = 33
    PyUnicode_EncodeLocale = 34
    PyUnicode_FSConverter = 35
    PyUnicode_FSDecoder = 36
    PyUnicode_DecodeFSDefaultAndSize = 37
    PyUnicode_DecodeFSDefault = 38
    PyUnicode_EncodeFSDefault = 39
    PyUnicode_FromWideChar = 40
    PyUnicode_AsWideChar = 41
    PyUnicode_AsWideCharString = 42
    PyUnicode_Decode = 43
    PyUnicode_AsEncodedString = 44
    PyUnicode_Encode = 45
    PyUnicode_DecodeUTF8 = 46
    PyUnicode_DecodeUTF8Stateful = 47
    PyUnicode_AsUTF8String = 48
    PyUnicode_AsUTF8AndSize = 49
    PyUnicode_AsUTF8 = 50
    PyUnicode_EncodeUTF8 = 51
    PyUnicode_DecodeUTF32 = 52
    PyUnicode_DecodeUTF32Stateful = 53
    PyUnicode_AsUTF32String = 54
    PyUnicode_EncodeUTF32 = 55
    PyUnicode_DecodeUTF16 = 56
    PyUnicode_DecodeUTF16Stateful = 57
    PyUnicode_AsUTF16String = 58
    PyUnicode_EncodeUTF16 = 59
    PyUnicode_DecodeUTF7 = 60
    PyUnicode_DecodeUTF7Stateful = 61
    PyUnicode_EncodeUTF7 = 62
    PyUnicode_DecodeUnicodeEscape = 63
    PyUnicode_AsUnicodeEscapeString = 64
    PyUnicode_EncodeUnicodeEscape = 65
    PyUnicode_DecodeRawUnicodeEscape = 66
    PyUnicode_AsRawUnicodeEscapeString = 67
    PyUnicode_EncodeRawUnicodeEscape = 68
    PyUnicode_DecodeASCII = 69
    PyUnicode_AsASCIIString = 70
    PyUnicode_EncodeASCII = 71
    PyUnicode_Concat = 72
    PyUnicode_Split = 73
    PyUnicode_Splitlines = 74
    PyUnicode_Join = 75
    PyUnicode_Tailmatch = 76
    PyUnicode_Find = 77
    PyUnicode_FindChar = 78
    PyUnicode_Count = 79
    PyUnicode_Replace = 80
    PyUnicode_Compare = 81
    PyUnicode_CompareWithASCIIString = 82
    PyUnicode_RichCompare = 83
    PyUnicode_Format = 84
    PyUnicode_Contains = 85
    PyUnicode_InternInPlace = 86
    PyUnicode_InternFromString = 87
    PyObject_Init = 88
    PyObject_InitVar = 89
    PyObject_New = 90
    PyObject_NewVar = 91
    PyObject_Del = 92
    PyObject_Length = 93
    PyObject_Str = 94
    PyObject_Bytes = 95
    PyObject_IsSubclass = 96
    PyObject_IsInstance = 97
    PyObject_GetAttr = 98
    PyObject_Call = 99
    PyObject_CallObject = 100
    PyObject_CallNoArgs = 101
    PyObject_ASCII = 102
    PyObject_Hash = 103
    PyObject_IsTrue = 104
    PyObject_Not = 105
    PyObject_Type = 106
    PyObject_HasAttr = 107
    PyObject_Print = 108
    PyObject_HasAttrString = 109
    PyObject_GetAttrString = 110
    PyObject_GenericGetAttr = 111
    PyObject_SetAttr = 112
    PyObject_SetAttrString = 113
    PyObject_DelAttr = 114
    PyObject_DelAttrString = 115
    PyObject_GenericGetDict = 116
    PyObject_GenericSetDict = 117
    PyObject_CallFunction = 118
    PyObject_CallFunctionObjArgs = 119
    PyObject_CallMethod = 120
    PyObject_LengthHint = 121
    PyObject_GetItem = 122
    PyObject_SetItem = 123
    PyObject_DelItem = 124
    PyObject_TypeCheck = 125
    PyObject_Dir = 126
    PyObject_GetIter = 127
    PyLong_Type = 128
    PyLong_Check = 129
    PyLong_CheckExact = 130
    PyLong_AsDouble = 131
    PyLong_AsVoidPtr = 132
    PyLong_AsLong = 133
    PyLong_AsLongLong = 134
    PyLong_AsLongAndOverflow = 135
    PyLong_FromDouble = 136
    PyLong_FromLong = 137
    PyLong_FromLongLong = 138
    PyLong_FromUnsignedLong = 139
    PyLong_FromUnsignedLongLong = 140
    PyLong_AsUnsignedLong = 141
    PyLong_AsSize_t = 142
    PyLong_AsUnsignedLongLong = 143
    PyLong_AsUnsignedLongMask = 144
    PyLong_AsUnsignedLongLongMask = 145
    PyLong_AsLongLongAndOverflow = 146
    PyLong_AsSsize_t = 147
    PyLong_FromSsize_t = 148
    PyLong_FromSize_t = 149
    PyLong_FromString = 150
    PyLong_FromUnicodeObject = 151
    PyLong_FromVoidPtr = 152
    PyImport_Import = 153
    PyModule_Type = 154
    PyModule_Check = 155
    PyModule_CheckExact = 156
    PyModule_NewObject = 157
    PyModule_New = 158
    PyModule_Create2 = 159
    PyModule_GetName = 160
    PyModule_GetDict = 161
    PyModule_GetState = 162
    PyModule_GetDef = 163
    PyModule_GetFilenameObject = 164
    PyModule_GetFilename = 165
    PyModule_GetNameObject = 166
    PyModule_AddFunctions = 167
    PyModule_AddIntConstant = 168
    PyModule_AddStringConstant = 169
    PyModule_AddIntMacro = 170
    PyModule_AddStringMacro = 171
    PyModule_AddType = 172
    PyModule_AddObject = 173
    PyModule_AddObjectRef = 174
    PyModuleDef_Init = 175
    create_module = 176
    exec_module = 177
    PyModule_FromDefAndSpec = 178
    PyModule_FromDefAndSpec2 = 179
    PyModule_ExecDef = 180
    PyModule_SetDocString = 181
    PyDict_Type = 182
    PyDict_Check = 183
    PyDict_CheckExact = 184
    PyDict_Contains = 185
    PyDict_Copy = 186
    PyDict_SetItem = 187
    PyDict_SetItemString = 188
    PyDict_New = 189
    PyDict_Size = 190
    PyDict_Clear = 191
    PyDict_GetItem = 192
    PyDict_Keys = 193
    PyDict_GetItemString = 194
    PyDict_GetItemWithError = 195
    PyDict_DelItem = 196
    PyDict_DelItemString = 197
    PyDict_SetDefault = 198
    PyDict_Items = 199
    PyDict_Values = 200
    PyDict_Next = 201
    PyDict_Merge = 202
    PyDict_Update = 203
    PyDict_MergeFromSeq2 = 204
    PyDict_ClearFreeList = 205
    PyEval_EvalCode = 206
    PyEval_EvalCodeEx = 207
    PyEval_EvalFrame = 208
    PyEval_EvalFrameEx = 209
    PyEval_MergeCompilerFlags = 210
    PyEval_InitThreads = 211
    PyEval_ThreadsInitialized = 212
    PyEval_SaveThread = 213
    PyEval_RestoreThread = 214
    PyEval_GetBuiltins = 215
    PyEval_GetLocals = 216
    PyEval_GetGlobals = 217
    PyEval_GetFrame = 218
    PyEval_SetProfile = 219
    PyFrame_GetBack = 220
    PyFrame_GetCode = 221
    PyFrame_GetLineNumber = 222
    PyTuple_Type = 223
    PyTuple_CheckExact = 224
    PyTuple_Size = 225
    PyTuple_GetItem = 226
    PyTuple_SetItem = 227
    PyTuple_GetSlice = 228
    PyTuple_Check = 229
    PyTuple_New = 230
    PyTuple_Resize = 231
    PyTuple_ClearFreeList = 232
    PyTuple_Pack = 233
    PyBytes_Type = 234
    PyBytes_Check = 235
    PyBytes_CheckExact = 236
    PyBytes_FromString = 237
    PyBytes_FromStringAndSize = 238
    PyBytes_FromFormat = 239
    PyBytes_FromFormatV = 240
    PyBytes_FromObject = 241
    PyBytes_Size = 242
    PyBytes_AsString = 243
    PyBytes_AsStringAndSize = 244
    PyBytes_Concat = 245
    PyBytes_ConcatAndDel = 246
    PyBytes_Resize = 247
    PyType_Type = 248
    PyType_Name = 249
    PyType_GetName = 250
    PyType_GetModule = 251
    PyType_Check = 252
    PyType_CheckExact = 253
    PyType_ClearCache = 254
    PyType_GetFlags = 255
    PyType_Modified = 256
    PyType_HasFeature = 257
    PyType_IS_GC = 258
    PyType_IsSubtype = 259
    PyType_GenericAlloc = 260
    PyType_GenericNew = 261
    PyType_Ready = 262
    PyType_GetSlot = 263
    PyType_FromSpecWithBases = 264
    PyType_FromSpec = 265
    PyList_Type = 266
    PyList_Check = 267
    PyList_CheckExact = 268
    PyList_New = 269
    PyList_SetItem = 270
    PyList_GetItem = 271
    PyList_GetSlice = 272
    PyList_SetSlice = 273
    PyList_Size = 274
    PyList_Insert = 275
    PyList_Append = 276
    PyList_Sort = 277
    PyList_Reverse = 278
    PyList_AsTuple = 279
    PyBool_Type = 280
    PyBool_FromLong = 281
    PyBool_Check = 282
    PyExc_Exception = 283
    PyExc_ValueError = 284
    PyErr_Clear = 285
    PyErr_PrintEx = 286
    PyErr_Print = 287
    PyErr_WriteUnraisable = 288
    PyErr_SetString = 289
    PyErr_SetObject = 290
    PyErr_Format = 291
    PyErr_FormatV = 292
    PyErr_SetNone = 293
    PyErr_BadArgument = 294
    PyErr_NoMemory = 295
    PyErr_SetFromErrno = 296
    PyErr_SetFromErrnoWithFilenameObject = 297
    PyErr_SetFromErrnoWithFilenameObjects = 298
    PyErr_SetFromErrnoWithFilename = 299
    PyErr_SetFromWindowsErr = 300
    PyErr_SetExcFromWindowsErr = 301
    PyErr_SetFromWindowsErrWithFilename = 302
    PyErr_SetExcFromWindowsErrWithFilenameObject = 303
    PyErr_SetExcFromWindowsErrWithFilenameObjects = 304
    PyErr_SetExcFromWindowsErrWithFilename = 305
    PyErr_SetImportError = 306
    PyErr_SyntaxLocationObject = 307
    PyErr_SyntaxLocationEx = 308
    PyErr_SyntaxLocation = 309
    PyErr_BadInternalCall = 310
    PyErr_WarnEx = 311
    PyErr_SetImportErrorSubclass = 312
    PyErr_WarnExplicitObject = 313
    PyErr_WarnExplicit = 314
    PyErr_WarnFormat = 315
    PyErr_ResourceWarning = 316
    PyErr_Occurred = 317
    PyErr_ExceptionMatches = 318
    PyErr_GivenExceptionMatches = 319
    PyErr_Fetch = 320
    PyErr_Restore = 321
    PyErr_NormalizeException = 322
    PyErr_GetExcInfo = 323
    PyErr_SetExcInfo = 324
    PyErr_CheckSignals = 325
    PyErr_SetInterrupt = 326
    PySignal_SetWakeupFd = 327
    PyErr_NewException = 328
    PyErr_NewExceptionWithDoc = 329
    PyException_GetTraceback = 330
    PyException_SetTraceback = 331
    PyException_GetContext = 332
    PyException_SetContext = 333
    PyException_GetCause = 334
    PyException_SetCause = 335
    PyUnicodeDecodeError_Create = 336
    PyUnicodeEncodeError_Create = 337
    PyUnicodeTranslateError_Create = 338
    PyUnicodeDecodeError_GetEncoding = 339
    PyUnicodeEncodeError_GetEncoding = 340
    PyUnicodeDecodeError_GetObject = 341
    PyUnicodeEncodeError_GetObject = 342
    PyUnicodeTranslateError_GetObject = 343
    PyUnicodeDecodeError_GetStart = 344
    PyUnicodeEncodeError_GetStart = 345
    PyUnicodeTranslateError_GetStart = 346
    PyUnicodeDecodeError_SetStart = 347
    PyUnicodeEncodeError_SetStart = 348
    PyUnicodeTranslateError_SetStart = 349
    PyUnicodeDecodeError_GetEnd = 350
    PyUnicodeEncodeError_GetEnd = 351
    PyUnicodeTranslateError_GetEnd = 352
    PyUnicodeDecodeError_SetEnd = 353
    PyUnicodeEncodeError_SetEnd = 354
    PyUnicodeTranslateError_SetEnd = 355
    PyUnicodeDecodeError_GetReason = 356
    PyUnicodeEncodeError_GetReason = 357
    PyUnicodeTranslateError_GetReason = 358
    PyUnicodeDecodeError_SetReason = 359
    PyUnicodeEncodeError_SetReason = 360
    PyUnicodeTranslateError_SetReason = 361
    PyInstanceMethod_Type = 362
    PyInstanceMethod_New = 363
    PyInstanceMethod_Function = 364
    PyCFunction_New = 365
    Py_GetPath = 366
    Py_GetPythonHome = 367
    py_fopen_obj = 368
    Py_BuildValue = 369
    Py_GetPlatform = 370
    Py_GetVersion = 371
    Py_GetCompiler = 372
    Py_GetBuildInfo = 373
    Py_GetExecPrefix = 374
    Py_GetProgramFullPath = 375
    Py_GetPrefix = 376
    Py_GetCopyright = 377
    Py_SetStandardStreamEncoding = 378
    Py_NewInterpreter = 379
    Py_EndInterpreter = 380
    Py_AddPendingCall = 381
    Py_tracefunc = 382
    Py_CompileString = 383
    Py_CompileStringFlags = 384
    Py_CompileStringObject = 385
    Py_CompileStringExFlags = 386
    Py_eval_input = 387
    Py_file_input = 388
    Py_single_input = 389
    Py_EnterRecursiveCall = 390
    Py_ReprEnter = 391
    Py_ReprLeave = 392
    Py_FatalError = 393
    Py_Exit = 394
    Py_AtExit = 395
    Py_ExitStatusException = 396
    Py_FdIsInteractive = 397
    Py_InitializeFromConfig = 398
    Py_PreInitialize = 399
    Py_PreInitializeFromArgs = 400
    Py_PreInitializeFromBytesArgs = 401
    Py_RunMain = 402
    Py_LeaveRecursiveCall = 403
    PyMethod_Type = 404
    PyMethod_Check = 405
    PyMethod_New = 406
    PyMethod_Function = 407
    PyMethod_Self = 408
    PyState_FindModule = 409
    PyState_AddModule = 410
    PyState_RemoveModule = 411
)