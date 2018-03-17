package gpu

//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_gpu
//#cgo windows LDFLAGS: -lSDL2 -lSDL2_gpu
//#include <stdlib.h>
//#include "sdl_gpu_wrapper.h"
import "C"
import (
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type Rect C.GPU_Rect

type RendererID C.GPU_RendererID

type RendererEnum int

const (
	RendererUnknown     RendererEnum = 0
	RendererOpenGL1Base RendererEnum = 1
	RendererOpenGL1     RendererEnum = 2
	RendererOpenGL2     RendererEnum = 3
	RendererOpenGL3     RendererEnum = 4
	RendererOpenGL4     RendererEnum = 5
	RendererGLES1       RendererEnum = 11
	RendererGLES2       RendererEnum = 12
	RendererGlES3       RendererEnum = 13
	RendererD3D9        RendererEnum = 21
	RendererD3D10       RendererEnum = 22
	RendererD3D11       RendererEnum = 23
)

const (
	ComparisonNever    = C.GPU_NEVER
	ComparisonLess     = C.GPU_LESS
	ComparisonEqual    = C.GPU_EQUAL
	ComparisonLEqual   = C.GPU_LEQUAL
	ComparisonGreater  = C.GPU_GREATER
	ComparisonNotEqual = C.GPU_NOTEQUAL
	ComparisonGEqual   = C.GPU_GEQUAL
	ComparisonAlways   = C.GPU_ALWAYS
)

const (
	BlendFuncZero             = C.GPU_FUNC_ZERO
	BlendFuncOne              = C.GPU_FUNC_ONE
	BlendFuncSrcColor         = C.GPU_FUNC_SRC_COLOR
	BlendFuncDstColor         = C.GPU_FUNC_DST_COLOR
	BlendFuncOneMinusSrc      = C.GPU_FUNC_ONE_MINUS_SRC
	BlendFuncOneMinusDst      = C.GPU_FUNC_ONE_MINUS_DST
	BlendFuncSrcAlpha         = C.GPU_FUNC_SRC_ALPHA
	BlendFuncDstAlpha         = C.GPU_FUNC_DST_ALPHA
	BlendFuncOneMinusSrcAlpha = C.GPU_FUNC_ONE_MINUS_SRC_ALPHA
	BlendFuncOneMinusDstAlpha = C.GPU_FUNC_ONE_MINUS_DST_ALPHA
)

const (
	BlendEqAdd             = C.GPU_EQ_ADD
	BlendEqSubtract        = C.GPU_EQ_SUBTRACT
	BlendEqReverseSubtract = C.GPU_EQ_REVERSE_SUBTRACT
)

type BlendMode C.GPU_BlendMode

const (
	BlendPresetNormal             = C.GPU_BLEND_NORMAL
	BlendPresetPremultipliedAlpha = C.GPU_BLEND_PREMULTIPLIED_ALPHA
	BlendPresetMultiply           = C.GPU_BLEND_MULTIPLY
	BlendPresetAdd                = C.GPU_BLEND_ADD
	BlendPresetSubtract           = C.GPU_BLEND_SUBTRACT
	BlendPresetModAlpha           = C.GPU_BLEND_MOD_ALPHA
	BlendPresetSetAlpha           = C.GPU_BLEND_SET_ALPHA
	BlendPresetSet                = C.GPU_BLEND_SET
	BlendPresetNormalKeepAlpha    = C.GPU_BLEND_NORMAL_KEEP_ALPHA
	BlendPresetAddAlpha           = C.GPU_BLEND_NORMAL_ADD_ALPHA
	BlendPresetFactorAlpha        = C.GPU_BLEND_NORMAL_FACTOR_ALPHA
)

// TODO: Resume at 210

type Image C.GPU_Image

// Broken
// type TextureHandle c.GPU_TextureHandle

type Context C.GPU_Context

type Target C.GPU_Target

// TODO: L460 GPU_FeatureEnum

const (
	VertexShader   = C.GPU_VERTEX_SHADER
	FragmentShader = C.GPU_FRAGMENT_SHADER
	PixelShader    = C.GPU_PIXEL_SHADER
	GeometryShader = C.GPU_GEOMETRY_SHADER
)

const (
	LanguageNone        = C.GPU_LANGUAGE_NONE
	LanguageARBAssembly = C.GPU_LANGUAGE_ARB_ASSEMBLY
	LanguageGLSL        = C.GPU_LANGUAGE_GLSL
	LanguageGLSLES      = C.GPU_LANGUAGE_GLSLES
	LanguageHLSL        = C.GPU_LANGUAGE_HLSL
	LanguageCG          = C.GPU_LANGUAGE_CG
)

type AttributeFormat C.GPU_AttributeFormat

type Attribute C.GPU_Attribute

type AttributeSource C.GPU_AttributeSource

const (
	ErrorNone                = C.GPU_ERROR_NONE
	ErrorBackendError        = C.GPU_ERROR_BACKEND_ERROR
	ErrorDataError           = C.GPU_ERROR_DATA_ERROR
	ErrorUserError           = C.GPU_ERROR_USER_ERROR
	ErrorUnsupportedFunction = C.GPU_ERROR_UNSUPPORTED_FUNCTION
	ErrorNullArgument        = C.GPU_ERROR_NULL_ARGUMENT
	ErrorFileNotFound        = C.GPU_ERROR_FILE_NOT_FOUND
)

type ErrorObject C.GPU_ErrorObject

const (
	DebugLevel0   = C.GPU_DEBUG_LEVEL_0
	DebugLevel1   = C.GPU_DEBUG_LEVEL_1
	DebugLevel2   = C.GPU_DEBUG_LEVEL_2
	DebugLevel3   = C.GPU_DEBUG_LEVEL_3
	DebugLevelMax = C.GPU_DEBUG_LEVEL_MAX
)

const (
	LogInfo    = C.GPU_LOG_INFO
	LogWarning = C.GPU_LOG_WARNING
	LogError   = C.GPU_LOG_ERROR
)

type Renderer C.GPU_Renderer

func Init(w, h uint16, flags uint32) *Target {
	return (*Target)(unsafe.Pointer(C.GPU_Init(C.Uint16(w), C.Uint16(h), C.Uint32(flags))))
}

func (t *Target) cptr() *C.GPU_Target {
	return (*C.GPU_Target)(unsafe.Pointer(t))
}

func (t *Target) Clear() {
	C.GPU_Clear(t.cptr())
}

func (t *Target) SetColor(color sdl.Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_SetTargetColor(t.cptr(), _c)
}

func (t *Target) Flip() {
	C.GPU_Flip(t.cptr())
}

func (t *Target) RectangleFilled(x1, y1, x2, y2 float32, color sdl.Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_RectangleFilled(t.cptr(),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
		_c)
}
