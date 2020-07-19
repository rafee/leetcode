/*
 * @lc app=leetcode id=1344 lang=golang
 *
 * [1344] Angle Between Hands of a Clock
 */

package golang

// @lc code=start
func angleClock(hour int, minutes int) float64 {
	var hourPos float64
	if hour < 12 {
		hourPos = float64(hour * 30)
	}

	hourPos += float64(minutes) / 2
	minutePos := float64(minutes) * 6
	angle := minutePos - hourPos
	if angle < 0 {
		angle = (-angle)
	}
	if angle > 180 {
		angle = (360 - angle)
	}
	return angle
}

// @lc code=end
