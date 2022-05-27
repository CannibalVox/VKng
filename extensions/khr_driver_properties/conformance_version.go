package khr_driver_properties

type ConformanceVersion struct {
	Major    uint8
	Minor    uint8
	Subminor uint8
	Patch    uint8
}

func (v ConformanceVersion) IsAtLeast(other ConformanceVersion) bool {
	if v.Major > other.Major {
		return true
	} else if v.Major < other.Major {
		return false
	}

	if v.Minor > other.Minor {
		return true
	} else if v.Minor < other.Minor {
		return false
	}

	if v.Subminor > other.Subminor {
		return true
	} else if v.Subminor < other.Subminor {
		return false
	}

	return v.Patch >= other.Patch
}
