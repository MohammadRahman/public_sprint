package sprint;

import java.util.ArrayList;
import java.util.List;

public class Combinations {
    public List<String> combN(int n){
        List<String> result = new ArrayList<>();

        if(n<=0){
            return result;
        }
        generateCombinations("", 0, n , result);
        return result;
    }

    private void generateCombinations(String current, int start, int n, List<String>result){
        if(current.length() == n){
            result.add(current);
            return;
        }
        for(int i=start; i <= 9; i++){
            generateCombinations(current +i, i +1, n, result);
        }
    }
}
