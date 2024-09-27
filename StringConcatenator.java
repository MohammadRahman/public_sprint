package sprint;

public class StringConcatenator {
    public static String concatenate(String... strings) {
        String result = "";
        for (String str : strings) {
            result += str;
        }
        
        return result;
    }
}
