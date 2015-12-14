package arrayfire

/*
#include <arrayfire.h>
#include <af/defines.h>
*/
import "C"

type (
	Dype         C.af_dtype
	Source       C.af_source
	InterpType   C.af_interp_type
	BorderType   C.af_border_type
	Connectivity C.af_connectivity
	ConvMode     C.af_conv_mode
	ConvDomain   C.af_conv_domain
	MatchType    C.af_match_type
	YccStd       C.af_ycc_std
	CSpace       C.af_cspace_t
	MatProp      C.af_mat_prop
	NormType     C.af_norm_type
	Colormap     C.af_colormap
	ImageFormat  C.af_image_format
	// TODO - causeda  build error:  Backend      C.af_backend
)

const (
	F32 = 0 ///< 32-bit floating point values
	C32 = 1 ///< 32-bit complex floating point values
	F64 = 2 ///< 64-bit complex floating point values
	C64 = 3 ///< 64-bit complex floating point values
	B8  = 4 ///< 8-bit boolean values
	S32 = 5 ///< 32-bit signed integral values
	U32 = 6 ///< 32-bit unsigned integral values
	U8  = 7 ///< 8-bit unsigned integral values
	S64 = 8 ///< 64-bit signed integral values
	U64 = 9 ///< 64-bit unsigned integral values
)

/*
typedef enum {
    afDevice,   ///< Device pointer
    afHost,     ///< Host pointer
} af_source;

#define AF_MAX_DIMS 4

// A handle for an internal array object
typedef void * af_array;

typedef enum {
    AF_INTERP_NEAREST,  ///< Nearest Interpolation
    AF_INTERP_LINEAR,   ///< Linear Interpolation
    AF_INTERP_BILINEAR, ///< Bilinear Interpolation
    AF_INTERP_CUBIC,    ///< Cubic Interpolation
    AF_INTERP_LOWER     ///< Floor Indexed
} af_interp_type;

typedef enum {
    ///
    /// Out of bound values are 0
    ///
    AF_PAD_ZERO = 0,

    ///
    /// Out of bound values are symmetric over the edge
    ///
    AF_PAD_SYM
} af_border_type;

typedef enum {
    ///
    /// Connectivity includes neighbors, North, East, South and West of current pixel
    ///
    AF_CONNECTIVITY_4 = 4,

    ///
    /// Connectivity includes 4-connectivity neigbors and also those on Northeast, Northwest, Southeast and Southwest
    ///
    AF_CONNECTIVITY_8 = 8
} af_connectivity;

typedef enum {

    ///
    /// Output of the convolution is the same size as input
    ///
    AF_CONV_DEFAULT,

    ///
    /// Output of the convolution is signal_len + filter_len - 1
    ///
    AF_CONV_EXPAND,
} af_conv_mode;

typedef enum {
    AF_CONV_AUTO,    ///< ArrayFire automatically picks the right convolution algorithm
    AF_CONV_SPATIAL, ///< Perform convolution in spatial domain
    AF_CONV_FREQ,    ///< Perform convolution in frequency domain
} af_conv_domain;

typedef enum {
    AF_SAD = 0,   ///< Match based on Sum of Absolute Differences (SAD)
    AF_ZSAD,      ///< Match based on Zero mean SAD
    AF_LSAD,      ///< Match based on Locally scaled SAD
    AF_SSD,       ///< Match based on Sum of Squared Differences (SSD)
    AF_ZSSD,      ///< Match based on Zero mean SSD
    AF_LSSD,      ///< Match based on Locally scaled SSD
    AF_NCC,       ///< Match based on Normalized Cross Correlation (NCC)
    AF_ZNCC,      ///< Match based on Zero mean NCC
    AF_SHD        ///< Match based on Sum of Hamming Distances (SHD)
} af_match_type;

typedef enum {
    AF_YCC_601 = 601,  ///< ITU-R BT.601 (formerly CCIR 601) standard
    AF_YCC_709 = 709,  ///< ITU-R BT.709 standard
    AF_YCC_2020 = 2020  ///< ITU-R BT.2020 standard
} af_ycc_std;

typedef enum {
    AF_GRAY = 0, ///< Grayscale
    AF_RGB,      ///< 3-channel RGB
    AF_HSV,      ///< 3-channel HSV
    AF_YCbCr     ///< 3-channel YCbCr
} af_cspace_t;

typedef enum {
    AF_MAT_NONE       = 0,    ///< Default
    AF_MAT_TRANS      = 1,    ///< Data needs to be transposed
    AF_MAT_CTRANS     = 2,    ///< Data needs to be conjugate tansposed
    AF_MAT_CONJ       = 4,    ///< Data needs to be conjugate
    AF_MAT_UPPER      = 32,   ///< Matrix is upper triangular
    AF_MAT_LOWER      = 64,   ///< Matrix is lower triangular
    AF_MAT_DIAG_UNIT  = 128,  ///< Matrix diagonal contains unitary values
    AF_MAT_SYM        = 512,  ///< Matrix is symmetric
    AF_MAT_POSDEF     = 1024, ///< Matrix is positive definite
    AF_MAT_ORTHOG     = 2048, ///< Matrix is orthogonal
    AF_MAT_TRI_DIAG   = 4096, ///< Matrix is tri diagonal
    AF_MAT_BLOCK_DIAG = 8192  ///< Matrix is block diagonal
} af_mat_prop;

typedef enum {
    AF_NORM_VECTOR_1,      ///< treats the input as a vector and returns the sum of absolute values
    AF_NORM_VECTOR_INF,    ///< treats the input as a vector and returns the max of absolute values
    AF_NORM_VECTOR_2,      ///< treats the input as a vector and returns euclidean norm
    AF_NORM_VECTOR_P,      ///< treats the input as a vector and returns the p-norm
    AF_NORM_MATRIX_1,      ///< return the max of column sums
    AF_NORM_MATRIX_INF,    ///< return the max of row sums
    AF_NORM_MATRIX_2,      ///< returns the max singular value). Currently NOT SUPPORTED
    AF_NORM_MATRIX_L_PQ,   ///< returns Lpq-norm

    AF_NORM_EUCLID = AF_NORM_VECTOR_2, ///< The default. Same as AF_NORM_VECTOR_2
} af_norm_type;

typedef enum {
    AF_COLORMAP_DEFAULT = 0,    ///< Default grayscale map
    AF_COLORMAP_SPECTRUM= 1,    ///< Spectrum map
    AF_COLORMAP_COLORS  = 2,    ///< Colors
    AF_COLORMAP_RED     = 3,    ///< Red hue map
    AF_COLORMAP_MOOD    = 4,    ///< Mood map
    AF_COLORMAP_HEAT    = 5,    ///< Heat map
    AF_COLORMAP_BLUE    = 6     ///< Blue hue map
} af_colormap;

typedef enum {
    AF_FIF_BMP          = 0,    ///< FreeImage Enum for Bitmap File
    AF_FIF_ICO          = 1,    ///< FreeImage Enum for Windows Icon File
    AF_FIF_JPEG         = 2,    ///< FreeImage Enum for JPEG File
    AF_FIF_JNG          = 3,    ///< FreeImage Enum for JPEG Network Graphics File
    AF_FIF_PNG          = 13,   ///< FreeImage Enum for Portable Network Graphics File
    AF_FIF_PPM          = 14,   ///< FreeImage Enum for Portable Pixelmap (ASCII) File
    AF_FIF_PPMRAW       = 15,   ///< FreeImage Enum for Portable Pixelmap (Binary) File
    AF_FIF_TIFF         = 18,   ///< FreeImage Enum for Tagged Image File Format File
    AF_FIF_PSD          = 20,   ///< FreeImage Enum for Adobe Photoshop File
    AF_FIF_HDR          = 26,   ///< FreeImage Enum for High Dynamic Range File
    AF_FIF_EXR          = 29,   ///< FreeImage Enum for ILM OpenEXR File
    AF_FIF_JP2          = 31,   ///< FreeImage Enum for JPEG-2000 File
    AF_FIF_RAW          = 34    ///< FreeImage Enum for RAW Camera Image File
} af_image_format;

typedef enum {
    AF_BACKEND_DEFAULT = 0,  ///< Default backend order: OpenCL -> CUDA -> CPU
    AF_BACKEND_CPU     = 1,  ///< CPU a.k.a sequential algorithms
    AF_BACKEND_CUDA    = 2,  ///< CUDA Compute Backend
    AF_BACKEND_OPENCL  = 3,  ///< OpenCL Compute Backend
} af_backend;

// Below enum is purely added for example purposes
// it doesn't and shoudn't be used anywhere in the
// code. No Guarantee's provided if it is used.
typedef enum {
    AF_ID = 0
} af_someenum_t;

#ifdef __cplusplus
namespace af
{
    typedef af_dtype dtype;
    typedef af_source source;
    typedef af_interp_type interpType;
    typedef af_border_type borderType;
    typedef af_connectivity connectivity;
    typedef af_match_type matchType;
    typedef af_cspace_t CSpace;
    typedef af_someenum_t SomeEnum; // Purpose of Addition: How to add Function example
    typedef af_mat_prop trans;
    typedef af_conv_mode convMode;
    typedef af_conv_domain convDomain;
    typedef af_mat_prop matProp;
    typedef af_colormap ColorMap;
    typedef af_norm_type normType;
    typedef af_ycc_std YCCStd;
    typedef af_image_format imageFormat;
    typedef af_backend Backend;
}

#endif
*/
