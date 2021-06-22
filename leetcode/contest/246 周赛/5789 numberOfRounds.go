package _46_周赛

import (
	"strconv"
	"time"
)

func numberOfRounds(startTime string, finishTime string) int {
	st, _ := time.Parse("15:04", startTime)
	fin, _ := time.Parse("15:04", finishTime)
	if st.After(fin) {
		fin = fin.Add(time.Hour * 24)
	}
	startMin := startTime[3:]
	sm, _ := strconv.Atoi(startMin)
	switch {
	case sm > 0 && sm < 15:
		st = st.Add(time.Minute * time.Duration(15-sm))
	case sm > 15 && sm < 30:
		st = st.Add(time.Minute * time.Duration(30-sm))
	case sm > 30 && sm < 45:
		st = st.Add(time.Minute * time.Duration(45-sm))
	case sm > 45:
		st = st.Add(time.Minute * time.Duration(60-sm))
	}
	endMin := finishTime[3:]
	fm, _ := strconv.Atoi(endMin)
	switch {
	case fm > 0 && fm < 15:
		fin = fin.Add(-time.Minute * time.Duration(fm))
	case fm > 15 && fm < 30:
		fin = fin.Add(-time.Minute * time.Duration(fm-15))
	case fm > 30 && fm < 45:
		fin = fin.Add(-time.Minute * time.Duration(fm-30))
	case fm > 45:
		fin = fin.Add(-time.Minute * time.Duration(fm-45))
	}

	dur := fin.Sub(st).Minutes()
	return int(dur / 15)
}
