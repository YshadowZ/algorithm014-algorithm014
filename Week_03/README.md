学习笔记
# 递归

## 代码模板(golang)

```golang
func recursion(level int, params1, params2 ...) {
    // 递归终止条件
    if level > MAX_LEVEL {
        processSomething()
        return
    }
    // 函数的具体处理逻辑
    processLogic()

    // 递归调用
    recursion(level+1, params1, ...)

    // 如果还有其他数据要处理，比如全局变量的处理等
}
```

## 递归思维要点
* 避免人为递归
* 找到最近最简方法，将其拆解成可重复子问题
* 数学归纳法思维



# 分治

## 分治，将大问题拆分成可重复解决的子问题

## 分支代码模板

```golang
func divide_conquer(problem, param1, parm2,...) {
    // 终止条件
    if problem == nil {
        processResult()
        return
    }
    // 处理逻辑
    data := prepare_data(problem)
    subProblems = split_problem(problem, data)

    // 递归处理分解问题
    subResult1 := divide_conquer(subProblems[0], p1, ...)
    subResult2 := divide_conquer(subProblems[1], p1, ...)
    subResult3 := divide_conquer(subProblems[2], p1, ...)

    // 合并结果
    result := process_result(subResult1, subResult2, subResult3)
}
```