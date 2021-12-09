/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2021-12-05 22:11:19
 * @LastEditors: MoonKnight
 * @LastEditTime: 2021-12-07 00:16:39
 */

package track

import (
	"fmt"
	"os/exec"
)

func GetProcessNum() (result map[string]int) {
	cmd := "ps aux | grep 'bin/rootserver'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	count := 0

	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}

	for _, value := range out {
		if value == '\n' {
			count++
		}
	}

	if count == 3 {
		fmt.Println("right answer")
		result = map[string]int{
			"processnum": 1,
		}
	} else {
		fmt.Println("no answer")
	}

	fmt.Println(string(out))
	return
}
