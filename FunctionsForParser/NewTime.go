package FunctionsForParser

import (
	"strconv"
)

func NewTime(s string) (int, int, int, int, int, int, int) {
	NewYear, NewMonth, NewDays, NewHours, NewMinutes, NewSeconds, NewMSecond := 0, 0, 0, 0, 0, 0, 0

	NewYear, err := strconv.Atoi(s[0:4])
                if err != nil { panic(err) }

	if s[5:6] == "0" {
		M, err := strconv.Atoi(s[6:7])
                if err != nil { panic(err) }
		NewMonth = M
	} else {
                M, err := strconv.Atoi(s[5:7])
                if err != nil { panic(err) }
		NewMonth = M
        }

	if s[8:9] == "0" {
                D, err := strconv.Atoi(s[9:10])
                if err != nil { panic(err) }
		NewDays = D
        } else {
                D, err := strconv.Atoi(s[8:10])
                if err != nil { panic(err) }
                NewDays = D
        }

	if s[11:12] == "0" {
                H, err := strconv.Atoi(s[12:13])
                if err != nil { panic(err) }
                NewHours = H
        } else {
                H, err := strconv.Atoi(s[11:13])
                if err != nil { panic(err) }
                NewHours = H
        }

	if s[14:15] == "0" {
                M, err := strconv.Atoi(s[15:16])
                if err != nil { panic(err) }
                NewMinutes = M
        } else {
                M, err := strconv.Atoi(s[14:16])
                if err != nil { panic(err) }
                NewMinutes = M
        }

	if s[17:18] == "0" {
                S, err := strconv.Atoi(s[18:19])
                if err != nil { panic(err) }
                NewSeconds = S
        } else {
                S, err := strconv.Atoi(s[17:19])
                if err != nil { panic(err) }
                NewSeconds = S
	}

	if s[20:21] == "0" {
		if s[21:22] == "0" {
			if s[22:23] == "0" {
				if s[23:24] == "0" {
					if s[24:25] == "0" {
						if s[25:26] == "0" {
							NewMSecond = 0
						} else {
							NewMSecond, err = strconv.Atoi(s[25:26])
							if err != nil { panic(err) }
						}
					} else {
						NewMSecond, err = strconv.Atoi(s[24:26])
                                                if err != nil { panic(err) }
					}
				} else {
					NewMSecond, err = strconv.Atoi(s[23:26])
                                        if err != nil { panic(err) }
				}
			} else {
				NewMSecond, err = strconv.Atoi(s[22:26])
                      		 if err != nil { panic(err) }
			}
		} else {
			NewMSecond, err = strconv.Atoi(s[21:26])
                        if err != nil { panic(err) }
		}
	} else {
		NewMSecond, err = strconv.Atoi(s[20:26])
	        if err != nil { panic(err) }
	}

	return NewYear, NewMonth, NewDays, NewHours, NewMinutes, NewSeconds, NewMSecond


}