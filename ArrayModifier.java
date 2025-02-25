package sprint;
import java.util.ArrayList;

public class ArrayModifier {
    public static ArrayList<Double> removeElementsBetween(ArrayList<Double> list, int index1, int index2) {
        // solution code here
        index1 = Math.max(0, Math.min(index1, list.size()));
        index2 = Math.max(0, Math.min(index2, list.size()));

        if (index1 > index2) {
            int temp = index1;
            index1 = index2;
            index2 = temp;
        }
        if (index1 == index2) {
            return list;
        }
        list.subList(index1, index2).clear();
        
        return list;
    }
}
