package sprint;

public class ArraySorter {
    public double[] sortArray(double[] arr) {
        // solution code here
        int arrLength = arr.length;
        
        for(int i = 0; i < arrLength-1;i++){
            for(int j = 0; j < arrLength -i-1; j++){
                if(arr[j]>arr[j+1]){
                    double temp = arr[j];
                    arr[j]=arr[j+1];
                    arr[j+1]=temp;
                }
            }
        }
        return arr;
    }
}
