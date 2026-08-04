/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
    // slices.SortFunc(intervals, func(a, b Interval) int {
    //     return cmp.Compare(b.start, a.start)
    // })

    meetings := map[int]int{}
    for _, interval := range intervals {
        meetings[interval.start]++
        meetings[interval.end]--
    }

    times := make([]int, 0, len(meetings))
    for time := range meetings {
        times = append(times, time)
    }

    sort.Ints(times)

    prev := 0
    res := 0
    for _, time := range times {
        prev += meetings[time]
        res = max(prev, res)
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
