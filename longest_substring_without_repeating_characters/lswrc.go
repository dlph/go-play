package longestsubstringwithoutrepeatingcharacters

// https://leetcode.com/problems/longest-substring-without-repeating-characters
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	m := make(map[string]string)
	var ss string
	for _, r := range s {
		// check if character already seen
		if _, contains := m[string(r)]; contains {
			break
		}

		m[string(r)] = ""
		ss = ss + string(r)
	}

	head := len(ss)
	tail := lengthOfLongestSubstring(s[1:])

	if head > tail {
		return head
	}

	return tail
}
