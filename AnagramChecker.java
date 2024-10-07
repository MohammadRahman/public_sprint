package sprint;

import java.util.Arrays;

public class AnagramChecker {
    public static boolean areAnagrams(String str1, String str2) {
        if(str1 == null || str2 == null){
            return false;
        }
        if (str1.isEmpty() && str2.isEmpty()) {
            return true;
        }
        
        String formatedStr1 = str1.replaceAll("[^a-zA-Z0-9]", "").toLowerCase();
       String formatedStr2 = str2.replaceAll("[^a-zA-Z0-9]", "").toLowerCase();

       if(formatedStr1.length() != formatedStr2.length()){
        return false;
       }
       char[] charArr1 = formatedStr1.toCharArray();
       char[] charArr2 = formatedStr2.toCharArray();

       Arrays.sort(charArr1);
       Arrays.sort(charArr2);

       return Arrays.equals(charArr1, charArr2);

    }
}
