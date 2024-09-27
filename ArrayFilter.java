package sprint;
import java.util.ArrayList;
import java.util.List;
public class ArrayFilter {
    public int[][] filterBySum(int[][] arr, int value) {
        List<int[]> filteredSubArrays = new ArrayList<>();

        for(int[] subArray: arr){
            int sum = 0;
            for(int num: subArray){
                sum+= num;
            }
            if (sum >= value) {
                filteredSubArrays.add(subArray);
            }
        }
        return filteredSubArrays.toArray(new int[filteredSubArrays.size()][]);
    }
}
