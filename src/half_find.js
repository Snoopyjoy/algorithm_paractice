function findKPos ( nums, k, startIndex, endIndex ){
    if( startIndex >= endIndex ){
        return -1;
    }

    let middlePos = ((startIndex + (endIndex - startIndex)/2) >> 0);
    let middleVal = nums[middlePos];
    if( middleVal === k ){
        return middlePos;
    }else if( k > middleVal){
        return findKPos ( nums, k, middlePos, endIndex )
    }
    return findKPos ( nums, k, startIndex, middlePos );
}

console.log( findKPos( [1,2,2,3,3,3,4,5], 2, 0, 8 ) );