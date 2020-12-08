package practice

//要对golang map按照value进行排序，思路是不用map，用struct存放key和value
type Url struct {
	Key string
	Value int
}

type UrlList []Url

func MakeMap(urls []string) map[string]int {
	hashTable := make(map[string]int)
	for _,url := range urls {
		if value, ok := hashTable[url]; ok {
			hashTable[url] = value+1
		} else {
			hashTable[url] = 1
		}
	}

	return hashTable
}

func MapToStruct(urlMap map[string]int) UrlList {
	urlStruct := make(UrlList, len(urlMap))
	i:=0
	for k,v := range urlMap {
		urlStruct[i] = Url{k, v}
		i++
	}

	return urlStruct
}

func (url UrlList)QuickSort(start int, end int) {
	if start >= end {
		return
	}
	//获取分区点
	pivot := url.partition(start, end)
	url.QuickSort(start, pivot-1)
	url.QuickSort(pivot+1, end)
}

func (url UrlList)partition(start int, end int) int {

	pivotValue := url[end].Value
	i := start
	for j:=start;j < end;j++ {
		if url[j].Value < pivotValue {
			url[i], url[j] = url[j],url[i]
			i++
		}
	}
	url[i], url[end] = url[end], url[i]

	return i
}