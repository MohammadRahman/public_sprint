package sprint;

public class PalindromeChecker {
    public boolean isPalindrome(String input){
        if(input == null || input.isEmpty()){
            return true;
        }
        String formatedInput = input.replaceAll("[^a-zA-Z0-9]", "").toLowerCase();
        int left = 0;

        int right = formatedInput.length() - 1;

        while(left < right){
            if(formatedInput.charAt(left) != formatedInput.charAt(right)){
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
}
