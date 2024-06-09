package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func main() {

// Get the current time in UTC
currentTime := time.Now()

// Get the start of the current week in UTC
currentWeekStart := currentTime.AddDate(0, 0, -int(currentTime.Weekday()))

// Get the end of the previous week (i.e., the start of the current week) in UTC
previousWeekEnd := currentWeekStart.Add(-time.Second)

// Subtract 6 weeks from the end of the previous week in UTC
sixWeeksAgo := previousWeekEnd.AddDate(0, 0, -42) // 42 days in 6 weeks

// Set the time to midnight (start of the day) of the first day of the week in UTC
startTime := sixWeeksAgo.AddDate(0, 0, -int(sixWeeksAgo.Weekday())).Truncate(24 * time.Hour)

// Set the time to 23:59:59 (end of the day) of the previous week end in UTC
endTime := previousWeekEnd.Truncate(24 * time.Hour).Add(24*time.Hour - time.Second)

// Get the Unix timestamp (seconds since epoch) for the start time in UTC
startTimeUnix := startTime.Unix()

// Get the Unix timestamp (seconds since epoch) for the end time of the previous week in UTC
endTimeUnix := endTime.Unix()

fmt.Println("Start Time (6 weeks ago) in UTC:", startTimeUnix)
fmt.Println("End Time (Previous week) in UTC:", endTimeUnix)
}

// func main() {

// 	jewels := "aA"
// 	stones := "aAAbbbb"
// 	result := leetcode.NumJewelsInStones(jewels, stones)
// 	fmt.Println(result)
// }

// type UserData struct {
// 	Name         string
// 	managerLevel map[string]string
// }

// func main() {

// client := redis.NewRedissClient()

// ctx := context.Background()
// client.Do(ctx, client.B().Set().Key("firstkey").Value("firstvalue").Build()).Error()

// // get keys
// stringData, err := client.Do(ctx, client.B().Get().Key("firstkey").Build()).ToString()
// if err != nil {
// 	fmt.Println(err)
// }

// data, e := client.Do(
// 	ctx,
// 	client.
// 		B().
// 		Lpop().
// 		Key("queue:pubmatic").
// 		Build(),
// ).ToString()

// if e != nil {
// 	fmt.Println("Error e", e)
// }
// fmt.Println("hm", data)

// fmt.Println(data)

// decoded_data_key, err := base64.StdEncoding.DecodeString(data)
// if err != nil {
// 	fmt.Println(err)
// }

// arr := strings.Split(string(decoded_data_key), "+;#")

// fmt.Println(arr)
// fcount := len(strings.Split(arr[1], "/")) - 3
// fmt.Println(fcount)
// fsnotify.CallFSNotify()
// processing()
// data := []UserData{
// 	{
// 		Name: "A",
// 		managerLevel: ,
// 	},
// 	{
// 		Name:         "B",
// 		managerLevel: map[string]string{},
// 	},
// 	{
// 		Name:         "C,
// 		managerLevel: map[string]string{},
// 	},
// 	{
// 		Name:         "D",
// 		managerLevel: map[string]string{},
// 	},
// 	{
// 		Name:         "E",
// 		managerLevel: map[string]string{},
// 	},
// 	{
// 		Name:         "G",
// 		managerLevel: map[string]string{},
// 	},
// 	{
// 		Name:         "H",
// 		managerLevel: map[string]string{},
// 	},
// 	{
// 		Name:         "I",
// 		managerLevel: map[string]string{},
// 	},
// }

//  data1 := [][]int{{1, 2}}
// data := [][]int{{1, 2, 4}, {3, 3, 1}}

// data1 := [][]int{{-5}}
// data := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
// data := [][]int{{1, 2}, {2, 3}}

// data := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}

// token := []int{2,1,3}
// token := []int{5, 4, 3, 8}
// token := []int{200, 100}
// token := []int{91, 4, 75, 70, 66, 71, 91, 64, 37, 54}

// score := 20
// nums := []int{3, 4, 5, 2}
// apple := []int{2, 83, 62}

// order := "cba"
// s := "abcd"

// order := "kqep"
// s := "pekeq"

// k := 3
// data := []int{10, 3, 8, 9, 4}
// nums1 := []int{1, 1, 2}
// nums2 := []int{1, 2, 3}
// nums1 = []
// nums2 = [1,2,3], k = 2
// k := 2
// result := leetcode.KSmallestPairs(nums1, nums2, k)

// nums := []int{1, 2, 3}
// }

type maxHeap []UserData

type MaxHeap interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i].Level > h[j].Level

}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(UserData))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewMaxHeap(arr []UserData) *maxHeap {
	var maxHeap = &maxHeap{}
	*maxHeap = arr
	heap.Init(maxHeap)
	return maxHeap
}

type rowData struct {
	createdBy string
	command   string
	teamName  string
}

type UserData struct {
	Name            string         `json:"Name"`
	Level           int            `json:"Level"`
	CommandInfo     map[string]int `json:"CommandInfo"`
	Reportees       []UserData     `json:"Reportees"`
	CommandCount    int            `json:"commandCount"`
	CommandCountMap map[string]int `json:"commandCountMap"`
}

func processing() string {

	teamData := make(map[string]UserData)
	isUserProcessed := make(map[string]*UserData)

	rowData := []rowData{
		{
			createdBy: "vishnu",
			command:   "DownLoad",
			teamName:  "abc",
		},
		{
			createdBy: "surya",
			command:   "Upload",
			teamName:  "abc",
		},
		{
			createdBy: "vishnu",
			command:   "Upload",
			teamName:  "abc",
		},
		{
			createdBy: "nitish",
			command:   "Upload",
			teamName:  "abcd",
		},
		{
			createdBy: "rohit",
			command:   "Upload",
			teamName:  "abcd",
		},
		{
			createdBy: "surya",
			command:   "Upload",
			teamName:  "abc",
		},
		{
			createdBy: "surya",
			command:   "DownLoad",
			teamName:  "abc",
		},
		{
			createdBy: "surya",
			command:   "Upload",
			teamName:  "abc",
		},
	}

	for _, v := range rowData {
		// maxHeapData := NewMaxHeap([]UserData{})
		currentUser := v.createdBy
		dataFromDB := TestData(currentUser)
		existUser, ok := isUserProcessed[currentUser]
		var managerLevel map[string]string
		if err := json.Unmarshal([]byte(dataFromDB), &managerLevel); err != nil {
			fmt.Println("Error in UnMarshal Data")
		}

		queue := []UserData{}
		// if !ok {
		for k, val := range managerLevel {
			Level, err := strconv.Atoi(k)
			if err != nil {
				fmt.Println("Err")
			}

			commandCount := 0

			// check user command

			user := UserData{
				Level:           Level,
				Name:            val,
				CommandInfo:     map[string]int{},
				Reportees:       []UserData{},
				CommandCount:    commandCount,
				CommandCountMap: map[string]int{},
			}
			// heap.Push(maxHeapData, user)
			queue = append(queue, user)
			// isUserProcessed[val] = &user

		}

		fmt.Println("Queue : ", queue)

		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Level > queue[j].Level
		})

		var cData UserData
		if !ok {
			u := UserData{
				Name:  currentUser,
				Level: len(managerLevel) + 1,
				CommandInfo: map[string]int{
					v.command: 1,
				},
				Reportees:    []UserData{},
				CommandCount: 1,
				CommandCountMap: map[string]int{
					"count": 1,
				},
			}
			isUserProcessed[v.createdBy] = &u

			cData = u

		} else {
			existUser.CommandInfo[v.command]++
			count := 0
			for _, v := range existUser.CommandInfo {
				count += v
			}

			if len(existUser.Reportees) > 0 {
				for _, child := range existUser.Reportees {
					count += child.CommandCountMap["count"]
				}
			}
			existUser.CommandCountMap["count"] = count
			existUser.CommandCount = count

			for _, user := range queue {
				parentUser := isUserProcessed[user.Name]
				parentCount := 0

				for _, v := range parentUser.CommandInfo {
					parentCount += v
				}
				for _, v := range parentUser.Reportees {
					parentCount += v.CommandCountMap["count"]
				}

				parentUser.CommandCountMap["count"] = parentCount
				// update to the map
			}
			continue

		}
		for _, user := range queue {
			// for maxHeapData.Len() > 0 {
			// lastLevelUserData := heap.Pop(maxHeapData).(UserData)
			existUser, ok := isUserProcessed[user.Name]
			if !ok {
				user.Reportees = append(user.Reportees, cData)
				commandCount := 0
				// for _, v := range user.Reportees {
				// 	commandCount += v.CommandCount
				// }

				commandCountMap := 0
				for _, v := range user.Reportees {
					commandCountMap += v.CommandCountMap["count"]
				}

				user.CommandCountMap = make(map[string]int)

				user.CommandCountMap["count"] = commandCountMap
				user.CommandCount = commandCount
				isUserProcessed[user.Name] = &UserData{
					Name:        user.Name,
					Level:       user.Level,
					CommandInfo: user.CommandInfo,
					Reportees:   user.Reportees,
					// CommandCount: user.CommandCount,
					CommandCountMap: user.CommandCountMap,
				}

			} else {

				existUser.Reportees = append(existUser.Reportees, cData)
				commandCount := 0

				for _, v := range existUser.Reportees {
					commandCount += v.CommandCount
				}
				commandCountMap := 0
				for _, v := range existUser.Reportees {
					commandCountMap += v.CommandCountMap["count"]
				}
				existUser.CommandCountMap["count"] += commandCountMap
				existUser.CommandCount = commandCount
				updateUserDataInMap(&cData, teamData, v.teamName, existUser.Name)
				break
			}
			cData = user
		}

		_, isAlreadyPresent := teamData[v.teamName]
		if !isAlreadyPresent {
			teamData[v.teamName] = cData
		}
	}

	if byteData, err := json.Marshal(teamData); err == nil {
		fmt.Println("Team Data", string(byteData))
	}

	return ""
}
func TestData(Name string) string {

	data := map[string]string{
		"vishnu": `{"1": "vipul", "2": "paawan"}`,
		"surya":  `{"1": "vipul", "2": "ankit"}`,
		"nitish": `{"1": "ram", "2": "shyam"}`,
		"rohit":  `{"1": "ram", "2": "shyam", "3":"nitish"}`,
	}
	return data[Name]
}

func updateUserDataInMap(user *UserData, teamData map[string]UserData, teamName, parentName string) {

	topLevelUser, ok := teamData[teamName]
	if !ok {
		return
	}

	if topLevelUser.Name == user.Name {
		teamData[teamName] = *user
		return
	}

	// Recursively find and update the parent node in the teamData map
	updateParentNode(user, &topLevelUser, parentName)

	// Update the teamData map with the modified topLevelUser
	teamData[teamName] = topLevelUser
}

// Helper function to recursively update the parent node in the teamData map
func updateParentNode(user *UserData, parent *UserData, parentName string) bool {

	if parent.Name == parentName {
		parent.Reportees = append(parent.Reportees, *user)

		// count := 0
		// commandCount := 0
		// for _, v := range parent.Reportees {
		// 	count += v.CommandCount
		// 	// commandCount += v.CommandCountMap["count"]
		// }

		// parent.CommandCountMap["count"] = commandCount

		// parent.CommandCount = count
		return true
	}

	for i := range parent.Reportees {
		if updateParentNode(user, &parent.Reportees[i], parentName) {
			return true
		}
	}
	return false
}

// // Helper function to update UserData in the map
// func updateUserDataInMap(user *UserData, teamData map[string]UserData, teamName, userName string) {
// 	// If the user is a leaf node (i.e., has no Reportees)

// 	team := teamData[teamName]

// 	if team.Name == userName {
// 		teamData[teamName] = *user
// 		return
// 	} else {
// 		for _, teamUser := range team.Reportees {

// 			// end user
// 			if len(teamUser.Reportees) == 0 {

// 			} else {
// 				for _, reportee := range teamUser.Reportees {
// 					updateUserDataInMap(existUser, teamData, v.teamName, v.createdBy)
// 					updateUserDataInMap(&reportee, teamData)
// 				}
// 			}

// 		}
// 	}
// 	if len(user.Reportees) == 0 {
// 		// Find the parent of the user
// 		parentName := user.Name
// 		for _, teamUser := range teamData {
// 			for _, reportee := range teamUser.Reportees {
// 				if reportee.Name == parentName {
// 					// Update the CommandInfo of the parent
// 					reportee.CommandInfo[user.Name] = user.Level
// 					return
// 				}
// 			}
// 		}
// 	} else {
// 		// If the user has Reportees, recursively update them
// 		for _, reportee := range user.Reportees {
// 			updateUserDataInMap(&reportee, teamData)
// 		}
// 	}
// }
