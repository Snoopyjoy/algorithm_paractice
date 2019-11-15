//堆排序
function heapSort(arr){
    //构建最大堆
    for( let i = (((arr.length-2)/2)>>0); i >= 0; i-- ){
        downAjust( arr, i, arr.length );
    }

    for( let j = arr.length - 1; j > 0; j--){
        //从最大堆中取出堆顶最大值 与最后的元素交换位置
        let temp = arr[j];
        arr[j] = arr[0];
        arr[0] = temp;
        downAjust( arr, 0, j );
    }
}

function downAjust(arr, parentIndex, len){
    let temp = arr[parentIndex];
    //左孩子
    childIndex = 2 * parentIndex + 1;
    while( childIndex < len ){
        //选择较大的孩子
        //如果有右孩子 且右孩子的值大于左节点则定位到右孩子 ???
        if( childIndex + 1 < len && arr[childIndex+1] > arr[childIndex] ){
            childIndex++;
        }
        //父节点大于节点 跳出
        if( temp >= arr[childIndex] ){
            break;
        }

        arr[parentIndex] = arr[childIndex];
        parentIndex = childIndex;
        childIndex = 2* childIndex + 1;
    }
    arr[parentIndex] = temp;
}

// let arr =  [3,1,3,4,5,1,9,2];
// heapSort( arr );
// console.log(arr);

function createRndArr( len ){
    let result = [];
    for (let index = 0; index < len; index++) {
        result.push( (Math.random() * len)>>0 );
    }
    return result;
}
let arr = createRndArr( 1000 );

console.time("heap_sort");
for (let index = 0; index < 1000; index++) {
    let arr1 = arr.concat([]);
    heapSort( arr1 );
}
console.timeEnd("heap_sort");