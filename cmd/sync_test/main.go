package main

import "time"
import "fmt"

func main() {
	synchronizePeriod := time.Hour
	//fp := uint64(4567062927520364291)
	//fp := 1e9 // 1 second
	//fp := 3600e9 // 1 hour
	fp := 4567062927520364291

	//curr := time.Unix(1596179540, 0)
	////curr := time.Unix(0, 0)
	//var prev time.Time
	//lastI := 0
	//for i := 1; i < 10000; i++ {
	//	prev = curr
	//	curr = prev.Add(1*time.Second)
	//	if prev.UnixNano() == 0 {
	//		continue
	//	}
	//	//log.Println("prev", prev, "curr", curr)
	//	cts := (uint64(curr.UnixNano()) + uint64(fp)) % uint64(synchronizePeriod.Nanoseconds())
	//	pts := (uint64(prev.UnixNano()) + uint64(fp)) % uint64(synchronizePeriod.Nanoseconds())
	//	fmt.Println(cts)
	//	if cts < pts {
	//		//fmt.Println("cts", cts, "pts", pts, "diff", cts-pts, "sync", cts < pts, "i", i, "diffI", i-lastI)
	//		fmt.Println("sync time:", curr, "diffI", i-lastI)
	//		lastI = i
	//	}
	//}

	syncTime := (float64(uint64(synchronizePeriod.Nanoseconds())-(uint64(fp)%uint64(synchronizePeriod.Nanoseconds()))) / 1e9) / 60
	fmt.Println(syncTime)

}
