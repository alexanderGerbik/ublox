// Code generated by "stringer -output=strings_navpvt.go -trimprefix NavPVT -type=NavPVTFixType,NavPVTValid,NavPVTFlags,NavPVTFlags2,NavPVTFlags3"; DO NOT EDIT.

package ubx

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NavPVTNoFix-0]
	_ = x[NavPVTDeadReckoning-1]
	_ = x[NavPVTFix2D-2]
	_ = x[NavPVTFix3D-3]
	_ = x[NavPVTGNSS-4]
	_ = x[NavPVTTimeOnly-5]
}

const _NavPVTFixType_name = "NoFixDeadReckoningFix2DFix3DGNSSTimeOnly"

var _NavPVTFixType_index = [...]uint8{0, 5, 18, 23, 28, 32, 40}

func (i NavPVTFixType) String() string {
	if i >= NavPVTFixType(len(_NavPVTFixType_index)-1) {
		return "NavPVTFixType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NavPVTFixType_name[_NavPVTFixType_index[i]:_NavPVTFixType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NavPVTValidDate-1]
	_ = x[NavPVTValidTime-2]
	_ = x[NavPVTFullyResolved-4]
	_ = x[NavPVTValidMag-8]
}

const (
	_NavPVTValid_name_0 = "ValidDateValidTime"
	_NavPVTValid_name_1 = "FullyResolved"
	_NavPVTValid_name_2 = "ValidMag"
)

var (
	_NavPVTValid_index_0 = [...]uint8{0, 9, 18}
)

func (i NavPVTValid) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _NavPVTValid_name_0[_NavPVTValid_index_0[i]:_NavPVTValid_index_0[i+1]]
	case i == 4:
		return _NavPVTValid_name_1
	case i == 8:
		return _NavPVTValid_name_2
	default:
		return "NavPVTValid(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NavPVTGnssFixOK-1]
	_ = x[NavPVTDiffSoln-2]
	_ = x[NavPVTHeadVehValid-32]
	_ = x[NavPVTCarrSolnFloat-64]
	_ = x[NavPVTCarrSolnFixed-128]
}

const (
	_NavPVTFlags_name_0 = "GnssFixOKDiffSoln"
	_NavPVTFlags_name_1 = "HeadVehValid"
	_NavPVTFlags_name_2 = "CarrSolnFloat"
	_NavPVTFlags_name_3 = "CarrSolnFixed"
)

var (
	_NavPVTFlags_index_0 = [...]uint8{0, 9, 17}
)

func (i NavPVTFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _NavPVTFlags_name_0[_NavPVTFlags_index_0[i]:_NavPVTFlags_index_0[i+1]]
	case i == 32:
		return _NavPVTFlags_name_1
	case i == 64:
		return _NavPVTFlags_name_2
	case i == 128:
		return _NavPVTFlags_name_3
	default:
		return "NavPVTFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NavPVTConfirmedAvai-32]
	_ = x[NavPVTConfirmedDate-64]
	_ = x[NavPVTConfirmedTime-128]
}

const (
	_NavPVTFlags2_name_0 = "ConfirmedAvai"
	_NavPVTFlags2_name_1 = "ConfirmedDate"
	_NavPVTFlags2_name_2 = "ConfirmedTime"
)

func (i NavPVTFlags2) String() string {
	switch {
	case i == 32:
		return _NavPVTFlags2_name_0
	case i == 64:
		return _NavPVTFlags2_name_1
	case i == 128:
		return _NavPVTFlags2_name_2
	default:
		return "NavPVTFlags2(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NavPVTInvalidLlh-1]
}

const _NavPVTFlags3_name = "InvalidLlh"

var _NavPVTFlags3_index = [...]uint8{0, 10}

func (i NavPVTFlags3) String() string {
	i -= 1
	if i >= NavPVTFlags3(len(_NavPVTFlags3_index)-1) {
		return "NavPVTFlags3(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _NavPVTFlags3_name[_NavPVTFlags3_index[i]:_NavPVTFlags3_index[i+1]]
}
