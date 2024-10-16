package sprint;

public class OccurrenceCounter {
    public int countOccurrences(int[] arr, int element, int idx){
        if (arr == null) {
            return 0;
        }
        if(idx < 0 || idx >= arr.length){
            return 0;
        }
        if(arr[idx] == element){
            return 1 + countOccurrences(arr, element, idx+1);
        }else{
            return countOccurrences(arr, element, idx+1);
        }
    }
}
