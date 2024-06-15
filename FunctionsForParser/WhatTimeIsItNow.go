package FunctionsForParser

func WhatTimeItIsNow(x [14]int) ([7] int, bool) {
	var OurTime [7]int
	OurTime[0], OurTime[1], OurTime[2],
        OurTime[3], OurTime[4], OurTime[5], OurTime[6] = x[7], x[8], x[9], x[10], x[11], x[12], x[13]
	key := false
	if x[0] > x[7] {
		key = true
	} else if x[1] > x[8] && x[0] >= x[7] {
                key = true
	} else if x[2] > x[9] && x[1] >= x[8] && x[0] >= x[7] {
                key = true
	} else if x[3] > x[10] && x[2] >= x[9] && x[1] >= x[8] && x[0] >= x[7] {
                key = true
        } else if x[4] > x[11] && x[3] >= x[10] && x[2] >= x[9] && x[1] >= x[8] && x[0] >= x[7] {
                key = true
        } else if x[5] > x[12] && x[4] >= x[11] && x[3] >= x[10] && x[2] >= x[9] && x[1] >= x[8] && x[0] >= x[7] {
		key = true
	} else if x[6] > x[13] && x[5] >= x[12] && x[4] >= x[11] && x[3] >= x[10] && x[2] >= x[9] && x[1] >= x[8] && x[0] >= x[7] {
		key = true
	}

	if key {
		OurTime[0], OurTime[1], OurTime[2],
		OurTime[3], OurTime[4], OurTime[5], OurTime[6] = x[0], x[1], x[2], x[3], x[4], x[5], x[6]
	}

	return OurTime, key
}