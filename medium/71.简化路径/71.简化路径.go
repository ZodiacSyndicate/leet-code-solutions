/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */
import "strings"

func simplifyPath(path string) string {
	stack := []string{""}
	arr := strings.Split(path, "/")
	for _, str := range arr {
		if str == "" || str == "." {
			continue
		} else if str == ".." {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = []string{""}
			}
		} else {
			stack = append(stack, str)
		}
	}
	if len(stack) == 1 {
		return "/"
	}
	return strings.Join(stack, "/")
}

