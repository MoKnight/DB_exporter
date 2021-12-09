/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2021-12-06 15:49:46
 * @LastEditors: MoonKnight
 * @LastEditTime: 2021-12-06 16:50:15
 */

package main

import (
	"fmt"
	"os/exec"
)

func main() {
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

	if count == 1 {
		fmt.Println("no answer")
	} else if count == 2 {
		fmt.Println("right answer")
	} else {
		fmt.Println("no no no")
	}

	fmt.Println(string(out))
}
