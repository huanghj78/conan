package main

import (
	client "cpfi.client"
)

// 1710308823
func main() {
	path := "/root/cpfi-v2/sequences/test.json"
	client.StartWorking(path)
	// a := []int{1, 2, 3}
	// cmd := fmt.Sprintf("python3 /root/cpfi-v2/systems/opengauss/findLeader.py")
	// output, err := exec.Command("sh", "-c", cmd).Output()
	// fmt.Println(string(output))
	// if err != nil {
	// 	fmt.Println("!!!")
	// 	// os.Exit(1)
	// }

}
