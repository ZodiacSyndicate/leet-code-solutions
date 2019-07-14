/*
 * @lc app=leetcode.cn id=208 lang=javascript
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (59.43%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    8.6K
 * Total Submissions: 14.5K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
 */
class TrieNode {
  links = new Array(26)

  _isEnd = false

  containsKey = char => this.links[char.charCodeAt() - 'a'.charCodeAt()] != null

  get = char => this.links[char.charCodeAt() - 'a'.charCodeAt()]

  put = (char, node) =>
    (this.links[char.charCodeAt() - 'a'.charCodeAt()] = node)

  setEnd = () => (this._isEnd = true)

  isEnd = () => this._isEnd
}
/**
 * Initialize your data structure here.
 */
var Trie = function() {
  this.root = new TrieNode()
}

/**
 * Inserts a word into the trie.
 * @param {string} word
 * @return {void}
 */
Trie.prototype.insert = function(word) {
  let node = this.root
  for (let i = 0; i < word.length; i++) {
    const current = word.charAt(i)
    if (!node.containsKey(current)) {
      node.put(current, new TrieNode())
    }
    node = node.get(current)
  }
  node.setEnd()
}

Trie.prototype.searchPrefix = function(word) {
  let node = this.root
  for (let i = 0; i < word.length; i++) {
    const current = word.charAt(i)
    if (node.containsKey(current)) {
      node = node.get(current)
    } else {
      return null
    }
  }
  return node
}

/**
 * Returns if the word is in the trie.
 * @param {string} word
 * @return {boolean}
 */
Trie.prototype.search = function(word) {
  const node = this.searchPrefix(word)
  return node != null && node.isEnd()
}

/**
 * Returns if there is any word in the trie that starts with the given prefix.
 * @param {string} prefix
 * @return {boolean}
 */
Trie.prototype.startsWith = function(prefix) {
  return this.searchPrefix(prefix) != null
}

/**
 * Your Trie object will be instantiated and called as such:
 * var obj = new Trie()
 * obj.insert(word)
 * var param_2 = obj.search(word)
 * var param_3 = obj.startsWith(prefix)
 */
