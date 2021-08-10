package main

import (
	"context"
	"fmt"
	"github.com/goh-chunlin/go-onedrive/onedrive"
	"golang.org/x/oauth2"
	"strings"
)

const token = ""

// path:itemId 缓存
// 这个回头再说吧
var cache = map[string]string{}

func init() {
	cache = make(map[string]string)
}

// 他这个设定，一个用户可以用多个drive示例的
// 可以不管，或者设置。并且有默认实例。
// 这就是 drives 和 driveItem 的区别
// 这个确实可以选择 ... 但是他没有名字，待定
func userDrives(client *onedrive.Client) {
	res, err := client.Drives.List(context.Background())
	if err != nil {
		panic(err)
	}

	for _, v := range res.Drives {
		fmt.Println(v.Id)
	}
}

// 获取 client
func getClient(ctx context.Context) *onedrive.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := onedrive.NewClient(tc)

	return client
}

// 好吧，这个版本的废弃了
//func getItemId(ctx context.Context, client *onedrive.Client, path string) string {
//	// 因为 one drive 的标记形式，不是以路径的形式标记
//	// 而是统一 以内部的
//	// 为了达成 zero-abstract， 这是一个类似于不要太多过多消耗来抽象的思想
//	// 所以这里用的 search
//	// 否则就 list 一个一个找了
//	searchRes, err := client.DriveSearch.Search(ctx, path)
//	if err != nil {
//		panic(err)
//	}
//
//	if len(searchRes.DriveItems) <= 0 {
//		// not found
//		// 记得自己封装一下 error
//		return ""
//	}
//
//	// 暂且划分为0 因为假设只有一件
//	return searchRes.DriveItems[0].Id
//}

func getItemId(ctx context.Context, client *onedrive.Client, path string) string {
	item, ok := getItem(ctx, client, path)

	if !ok {
		return ""
	}

	return item.Id
}

// path 一定要是绝对路径，并且不以 / 结尾，就算是文件夹也是，要处理一下
func getItem(ctx context.Context, client *onedrive.Client, path string) (*onedrive.DriveItem, bool) {
	searchItem := func(folderId string) []*onedrive.DriveItem {
		listRes, err := client.DriveItems.List(ctx, folderId)
		if err != nil {
			panic(err)
		}

		return listRes.DriveItems
	}

	paths := strings.Split(path, "/")
	lenp := len(paths)

	// 从 root 也就是 "" 开始
	nextId := ""
	var searchRes *onedrive.DriveItem
	for i := 1; i < lenp; i++ {
		// 当前搜索的目标名字
		nowSearchValue := paths[i]
		list := searchItem(nextId)
		for _, v := range list {
			if v.Name == nowSearchValue {
				nextId = v.Id
				// 最后一个，也就是目标文件名了
				if i == lenp-1 {
					searchRes = v
				}
				break
			}
		}
	}

	if searchRes == nil {
		return nil, false
	}

	return searchRes, true
}

func listItems(ctx context.Context, client *onedrive.Client, path string) (items []*onedrive.DriveItem, count int) {
	// 以 list 为例
	// 要先从 path 获取到对应的 folderId
	itemId := getItemId(ctx, client, path)

	// 后面就一样了
	listRes, err := client.DriveItems.List(ctx, itemId)
	if err != nil {
		panic(err)
	}

	// 返回的是获取到的所有 items 和 数量
	return listRes.DriveItems, listRes.Count
}

func deleteItem(ctx context.Context, client *onedrive.Client, path string) {
	item, ok := getItem(ctx, client, path)
	fmt.Println(item, ok)

	item, err := client.DriveItems.Delete(ctx, "", item.Id)
	if err != nil {
		panic(err)
	}
}


func getObject(ctx context.Context, client *onedrive.Client, path string) {
	item, err := client.DriveItems.Get(ctx, path)
	if err != nil {
		panic(err)
	}
	fmt.Println(item)
}


func main() {
	commonCtx := context.Background()

	client := getClient(commonCtx)

	getObject(commonCtx, client, "t.go")

	//deleteItem(commonCtx, client, "/t.go")
	//fmt.Println(item, ok)

	//userDrives(client)
}
