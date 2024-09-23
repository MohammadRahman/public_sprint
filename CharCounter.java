package sprint;

public class CharCounter {
    public int countOccurrences(String input, char target) {
        // Implement logic here

        int result = 0;
        for(int i = 0; i <input.length(); i++){
            if(Character.toLowerCase(input.charAt(i)) == Character.toLowerCase(target)){
                result++;
            }
        }
        return result;
    }
}
