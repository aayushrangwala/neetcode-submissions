/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

 import(
    "slices"
    "cmp"
 )

func canAttendMeetings(intervals []Interval) bool {
    slices.SortFunc(intervals, func(a, b Interval) int {
        return cmp.Compare(a.start, b.start)
    })

    for index := 1; index < len(intervals); index++ {
        if intervals[index-1].end > intervals[index].start {
            return false
        }
    }

    return true
}
