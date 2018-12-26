package util

import (
	"regexp"
	"strings"
	"anonymousFriends/models"
)

var ContentFilter *Filter

func init(){
	ContentFilter = NewFilter()
	ContentFilter.LoadWordDict()
}

// Filter 敏感词过滤器
type Filter struct {
	trie  *Trie
	noise *regexp.Regexp
}

// New 返回一个敏感词过滤器
func NewFilter() *Filter {
	return &Filter{
		trie:  NewTrie(),
		noise: regexp.MustCompile(`[\s&%$@*]+`),
	}
}

// UpdateNoisePattern 更新去噪模式
func (filter *Filter) UpdateNoisePattern(pattern string) {
	filter.noise = regexp.MustCompile(pattern)
}

// LoadWordDict 加载敏感词字典
func (filter *Filter) LoadWordDict() error {
	//content, err := ioutil.ReadFile(path)
	//if err != nil {
	//	return err
	//}
	words := strings.Split(models.SensitiveWords, "\n")
	filter.trie.Add(words...)
	return nil
}

// LoadWordDict 加载敏感词字典
func (filter *Filter) NewLoadWordDict(path string) error {
	// fmt.Println("=====>>获取加载敏感词字典")
	//if baseServer.RedisCache.IsExist("backendDict") {
	//	backendDictTemp := baseServer.RedisCache.Get("backendDict")
	//	backendDictTempBytes, ok := backendDictTemp.([]byte)
	//	if !ok {
	//		resp, err := http.Get(path)
	//		if err != nil {
	//			// handle error
	//		}
	//		defer resp.Body.Close()
	//		body, err := ioutil.ReadAll(resp.Body)
	//		if err != nil {
	//			// handle error
	//		}
	//		var keyDict string
	//		keyDict = string(body)
	//		words := strings.Split(keyDict, ",")
	//		filter.trie.Add(words...)
	//		baseServer.RedisCache.Put("backendDict", keyDict, 60 * 60 * 24 * 365 * time.Second)
	//	}
	//	//fmt.Println("backendDictTemp:", backendDictTemp)
	//	backendDictStr, _ := redis.String(backendDictTempBytes, nil)
	//	// fmt.Println("backendDictStr:", backendDictStr)
	//	words := strings.Split(backendDictStr, ",")
	//	filter.trie.Add(words...)
	//}else{
	//	resp, err := http.Get(path)
	//	if err != nil {
	//		// handle error
	//	}
	//	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		// handle error
	//	}
	//	var keyDict string
	//	keyDict = string(body)
	//	words := strings.Split(keyDict, ",")
	//	filter.trie.Add(words...)
	//	baseServer.RedisCache.Put("backendDict", keyDict, 60 * 60 * 24 * 365 * time.Second)
	//}
	return nil
}

// AddWord 添加敏感词
func (filter *Filter) AddWord(words ...string) {
	filter.trie.Add(words...)
}

// Filter 过滤敏感词
func (filter *Filter) Filter(text string) string {
	return filter.trie.Filter(text)
}

// Replace 和谐敏感词
func (filter *Filter) Replace(text string, repl rune) string {
	return filter.trie.Replace(text, repl)
}

// FindIn 检测敏感词
func (filter *Filter) FindIn(text string) (bool, string) {
	text = filter.RemoveNoise(text)
	return filter.trie.FindIn(text)
}

// RemoveNoise 去除空格等噪音
func (filter *Filter) RemoveNoise(text string) string {
	return filter.noise.ReplaceAllString(text, "")
}
