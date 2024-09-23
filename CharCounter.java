package sprint;

public class CharCounter {
    public int countOccurrences(String input, char target) {
        // Implement logic here
        int result = 0;
        for(int i = 0; i <input.length(); i++){
            if(input.charAt(i) == target){
                result++;
            }
        }
        return result;
    }
}
