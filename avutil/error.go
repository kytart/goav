// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/imgutils.h>
//#include <stdlib.h>
//static const char *error2string(int code) { return av_err2str(code); }
//static int getAVErrorEAGAIN() { return AVERROR(EAGAIN); }
import "C"
import "errors"

var AvErrorEOF int = int(C.AVERROR_EOF)
var AvErrorEAGAIN int = int(C.getAVErrorEAGAIN())

func ErrorFromCode(code int) error {
	if code >= 0 {
		return nil
	}

	return errors.New(C.GoString(C.error2string(C.int(code))))
}
