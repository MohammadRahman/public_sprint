package sprint;

public class RecursivePalindrome {
    public boolean isPalindrome(String str) {
        if(str == null){
            return false;
        }
        str = removeEmptyString(str);

        if(str.length() <= 1){
            return true;
        }
        char firstChar = Character.toLowerCase(str.charAt(0));
        char lastChar = Character.toLowerCase(str.charAt(str.length() -1));

        if(firstChar == lastChar){
            return isPalindrome(str.substring(1, str.length() -1));
        }else{
            return false;
        }
    }
    private String removeEmptyString(String str){
        StringBuilder formatedStr = new StringBuilder();

        for(char c: str.toCharArray()){
            if(Character.isLetterOrDigit(c)){
                formatedStr.append(c);
            }
        }
        return formatedStr.toString();
    }
}
