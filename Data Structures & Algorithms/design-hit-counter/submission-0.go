type HitCounter struct {
     arr []int
}


func Constructor() HitCounter {
    return HitCounter{
        arr: []int{},
    }
}


func (this *HitCounter) Hit(timestamp int) {
    this.arr = append(this.arr, timestamp)
}


func (this *HitCounter) GetHits(timestamp int) int {
    for len(this.arr) > 0 {
        if timestamp - this.arr[0] >= 300 {
            this.arr = this.arr[1:]
            continue
        }

        break
    }

    return len(this.arr)
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param2 := obj.GetHits(timestamp);
 */
