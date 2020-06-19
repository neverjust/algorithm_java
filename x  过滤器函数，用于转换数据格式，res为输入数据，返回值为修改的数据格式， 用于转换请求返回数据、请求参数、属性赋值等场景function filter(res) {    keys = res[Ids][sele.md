```javascript
/**
* 过滤器函数，用于转换数据格式，res为输入数据，返回值为修改的数据格式，
* 用于转换请求返回数据、请求参数、属性赋值等场景
*/

function filter(res) {
    keys = res["Ids"]["selectedRowKeys"];
    var arr = new Array(); 
    for (let j = 1; j < keys.length; j++) {
        arr[j - 1] = parseInt(keys[j])
    } 
    res["Ids"] = arr
    return res
}

```

