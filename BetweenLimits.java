package sprint;

public class BetweenLimits {

    public String findRange(char from, char to){
         // Ensure that the characters are in ascending order
         if (from > to) {
            char temp = from;
            from = to;
            to = temp;
        }
         // Build the result string with all characters between from and to
         StringBuilder result = new StringBuilder();
        for (char c = (char)(from+1); c < to; c++){
            result.append(c);
        }
        return result.toString();
    }
}
