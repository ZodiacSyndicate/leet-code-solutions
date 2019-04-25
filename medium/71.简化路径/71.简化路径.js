/*
 * @lc app=leetcode.cn id=71 lang=javascript
 *
 * [71] 简化路径
 */
/**
 * @param {string} path
 * @return {string}
 */
var simplifyPath = function(path) {
  const stack = ['']
  const arr = path.split('/')
  for (const str of arr) {
    if (str === '' || str === '.') {
      continue
    } else if (str === '..') {
      stack.pop()
      if (!stack.length) {
        stack.push('')
      }
    } else {
      stack.push(str)
    }
  }
  return stack.length === 1 ? '/' : stack.join('/')
}
