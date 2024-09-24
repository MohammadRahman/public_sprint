package sprint;

public class ArrayAdder {
    public static int[] concatArrays(int[] arr1, int[] arr2) {
        // solution code here
        int[] newArray = new int[arr1.length + arr2.length];
        System.arraycopy(arr1, 0, newArray, 0, arr1.length);
        System.arraycopy(arr2, 0, newArray, arr1.length, arr2.length);

        return newArray;
    }
}
