学习笔记

# Trie树
> Trie树也成为字典树，是一种专门处理字符串匹配的数据结构，用来解决在一组字符串数组中快速查找指定字符串的问题
> Trie树也成为前缀树
> 核心思想是最大限度的减少无谓的字符串比较（空间换时间），提高查询效率

## Trie树的特点
* 根节点不包含字符，除了根节点外，每一个节点都只包含一个字符
* 从根节点到某一节点，路径上经历的所有字符连接起来，为该节点对应的字符串
* 每个节点的所有子节点都包含不同的字符

## Trie树的构造过程
![构造](https://user-gold-cdn.xitu.io/2019/1/2/1680c049b5365880?imageslim)

## Trie树的插入过程
![插入](https://user-gold-cdn.xitu.io/2019/1/2/1680c049b51f4db2?imageslim)

## Trie树的删除
* 删除整个单词
![删除整个单词](https://user-gold-cdn.xitu.io/2019/1/2/1680c049b68802a6?imageslim)
* 删除前缀单词
![删除前缀单词](https://user-gold-cdn.xitu.io/2019/1/2/1680c04a41644ab1?imageslim)
* 删除分支单词
![删除分支单词](https://user-gold-cdn.xitu.io/2019/1/2/1680c04a415421ef?imageslim)

## 应用
* 搜索引擎系统用于文本词频统计
* 字符串检索

# 并查集
并查集主要用于解决一些元素分组的问题，它管理一系列不想交的集合，并支持两种操作：
* 合并(union): 把两个集合合并为一个集合
* 查询(find): 查询两个元素是否在同一个集合中

## 使用场景
是否是亲戚，是否是朋友的和圈相关的问题

## 代码
* 初始化
```golang
/*
    初始化，每个元素自己都是独立的结合，父元素指向自己
    使用一个数组来表示，index为当前元素，value为父元素，若index == value 则此节点为根节点
*/
var arr []int{}

func init() {
    for var i = 0; i < len(arr); i++ {
        arr[i] = i
    }
}
```
* 查询
```golang
// 递归查询根节点
// 判断两个元素是否是同一个集合，只需判断根节点是否是同一个
func find(index int) {
    if arr[index] == index {
        return index
    } else {
        return find(arr[index])
    }
}
```
* 合并
```golang
// 合并
func merge(i, j int) {
    arr[find(i)] = find(j)
}
```

* 路径压缩

```golang
func find(index) {
    if arr[index] == index {
        return index
    } else {
        // 所有的子节点都直接指向根节点
        arr[index] = find(arr[index])
        return arr[index]
    }
}
```

## 资料参考和图片来源
[掘金](https://juejin.im/post/6844903750490914829)