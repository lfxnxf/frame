package valid

import "regexp"

// IP验证一个值是否为IP，可验证IP4
func IP(val interface{}) bool {
	// 匹配 IP4
	var ip4Pattern = `((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)`
	ip := regexpCompile(ip4Pattern)
	return isMatch(ip, val)
}
func regexpCompile(str string) *regexp.Regexp {
	return regexp.MustCompile("^" + str + "$")
}
func isMatch(exp *regexp.Regexp, val interface{}) bool {
	switch v := val.(type) {
	case []rune:
		return exp.MatchString(string(v))
	case []byte:
		return exp.Match(v)
	case string:
		return exp.MatchString(v)
	default:
		return false
	}
}
