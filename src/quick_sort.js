//关键点 基准点的选择

// 分治 双边循环
function quickSort( arr, startIndex, endIndex) {
    if( startIndex >= endIndex ){
        return;
    }
    const pivotIndex = partition1(arr, startIndex, endIndex);
    quickSort( arr, startIndex, pivotIndex - 1);
    quickSort( arr, pivotIndex + 1, endIndex );
}

function partition( arr, startIndex, endIndex){
    const pivot = arr[startIndex];
    let left = startIndex;
    let right = endIndex;
    while( left !== right ){
        //从右向左找到比基准值小的元素位置
        while( left < right && arr[right] > pivot ){
            right--;
        }
        //从左向右找到比基准值大的元素位置
        while( left < right && arr[left] <= pivot ){
            left++;
        }

        //交互元素 使小于基准的值放到基准值右边 大于基准的值放到左边
        if( left < right ){
            let temp = arr[right];
            arr[right] = arr[left];
            arr[left] = temp;
        }
    }

    //基准值和指针重合点交互
    arr[startIndex] = arr[left];
    arr[left] = pivot;
    return left;
}

// let arr =  [3,1,3,4,5,1,2];
// quickSort( arr, 0, 6);
// console.log(arr);

//单边循环
function partition1( arr, startIndex, endIndex ){
    const pivot = arr[startIndex];
    let mark = startIndex;
    for (let index = startIndex + 1; index <= endIndex; index++) {
        if( arr[index] < pivot ){
            mark++;
            let temp = arr[mark];
            arr[mark] = arr[index];
            arr[index] = temp;
        }
    }
    arr[startIndex] = arr[mark];
    arr[mark] = pivot;
    return mark;
}

//非递归 用栈代替递归
function quickSort2(arr, startIndex, endIndex){
    const stack = [{
        sid: startIndex,
        eid: endIndex
    }];
    while( stack.length > 0 ){
        const {sid, eid} = stack.pop();
        let pivotIndex = partition1( arr, sid, eid );
        if( sid < pivotIndex - 1 ){
            stack.push( {
                sid,
                eid: pivotIndex - 1
            } );
        }
        if( pivotIndex + 1 < eid ){
            stack.push( {
                sid: pivotIndex + 1,
                eid
            } );
        }
    }
}
// let arr =  [3,1,3,4,5,1,2];
// quickSort2( arr, 0, 6);
// console.log(arr);

function createRndArr( len ){
    let result = [];
    for (let index = 0; index < len; index++) {
        result.push( (Math.random() * len)>>0 );
    }
    return result;
}
let arr = createRndArr( 1000 );

console.time("quick_sort");
for (let index = 0; index < 1000; index++) {
    let arr1 = arr.concat([]);
    quickSort2( arr1 );
}
console.timeEnd("quick_sort");