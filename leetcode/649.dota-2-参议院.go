/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	l := len(senate)
	rs := list.New()
	ds := list.New()

	for i, v := range senate {
		if v == 'R' {
			rs.PushBack(i)
		} else {
			ds.PushBack(i)
		}
	}

	for rs.Front() != nil && ds.Front() != nil {
		if rs.Front().Value.(int) < ds.Front().Value.(int) {
			ds.Remove(ds.Front())
			rs.Front().Value = rs.Front().Value.(int) + l
			rs.MoveToBack(rs.Front())
		} else {
			rs.Remove(rs.Front())
			ds.Front().Value = ds.Front().Value.(int) + l
			ds.MoveToBack(ds.Front())
		}
	}
	if rs.Front() == nil {
		return "Dire"
	}
	return "Radiant"
}

// @lc code=end

