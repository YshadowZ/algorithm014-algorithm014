学习笔记

# 贪心算法

* 每一步选择当前状态下的最优解
* 适用于全局最优=局部最优的累加的情况
    * 贪心：当下局部最优解的判断
    * 回溯：能够回退
    * 动态规划：最优判断+回退
* 如果一个问题可以使用贪心算法求解，那么贪心算法一定是最优解

# 搜索-遍历

不考虑智能的情况下

* 每个节点都要访问一次
* 每个节点仅访问一次（避免过多无用的访问）
* 对于节点访问顺序的不同，分为
    * 深度优先 DFS
    * 广度优先 BFS
    * 优先级优先

## 示例代码

```golang
DFS (递归)

func dfs (node, visits map[node]bool) {
    if visits[node] {
        return
    }
    visits[node] = true
    dfs(node.left)
    dfs(node.right)
    // 多叉树的情况
    // for _, child := range node.children {
    //     if visits[child] {
    //         continue
    //     }
    //     visits[child] = true
    //     dfs(child, visits)
    // }
}
```

```golang
BFS (队列)
func bfs (graph, start, end) {
    queue := []int
    queue = append(queue, start)
    visits[start] = true
    for len(queue) != 0 {
        node = queue[len(queue) - 1]
        visits[node] = true

        process(node)
        nodes := generate_related_nodes(node)
        queue = append(queue, nodes)
    }
}

# 二分查找

* 目标函数存在单调性
* 存在上下界
* 能够通过索引访问

