	package leetcode

	import (
		"sort"
	)

	type meetingRoomDetails struct {
		currentMeetingSlot []int `currentMeetingSlot`
		availableAt        int   `availableAt`
		count              int   `count`
		meetingNumber      int   `meetingNumber`
	}

	func MostBooked(n int, meetings [][]int) int {

		// sort the array on the basis of the start time
		sort.Slice(meetings, func(i, j int) bool {
			return meetings[i][0] < meetings[j][0]
		})

		// now create a map which tarck a record for meeting

		meetingRecord := make(map[int]meetingRoomDetails)

		// allot all the meeting room for the first element

		if len(meetings) > n {
			for i := 0; i < n; i++ {
				meetingRecord[i] = meetingRoomDetails{
					currentMeetingSlot: meetings[i],
					availableAt:        meetings[i][1],
					count:              1,
				}
			}
		} else {
			return 0
		}

		for i := n; i < len(meetings); i++ {
			//find out least meeting avalability
			meetingIndex := -1
			minAvailability := 5*1000000 + 1

			// first check if more than two room is avilable on that time priority will be first indext room

			for k, v := range meetingRecord {
				if v.availableAt < meetings[i][0] {
					meetingIndex = k
					break
				}
				if v.availableAt < minAvailability {
					minAvailability = v.availableAt
					meetingIndex = k
				}
			}
			// sort the available room key

			count := meetingRecord[meetingIndex].count + 1
			// now update the meeting index for the next meeting
			// Now check if the available at is delay or not
			diff := meetings[i][1] - meetings[i][0]
			if meetingRecord[meetingIndex].availableAt > meetings[i][0] {
				// updating meeting slot

				start := meetingRecord[meetingIndex].currentMeetingSlot[1]
				end := start + diff
				// meetings[i] = []int{start, end}

				meetingRecord[meetingIndex] = meetingRoomDetails{
					currentMeetingSlot: []int{start, end},
					availableAt:        start + diff,
					count:              count,
				}
			} else {
				meetingRecord[meetingIndex] = meetingRoomDetails{
					currentMeetingSlot: meetings[i],
					availableAt:        meetings[i][0] + diff,
					count:              count,
				}
			}
		}

		// return most booked
		arr := []meetingRoomDetails{}
		for k, v := range meetingRecord {
			v.meetingNumber = k
			arr = append(arr, v)
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i].count > arr[j].count
		})
		return arr[0].meetingNumber
	}
