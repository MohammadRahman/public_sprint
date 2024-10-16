package sprint;

public class ParenthesesBalanceChecker {
    public boolean isBalanced(String str) {
        if(str == null){
            return false;
        }
        return balanceCheck(str, 0, 0);
    }

    private boolean balanceCheck(String str, int idx, int balance){
        if(idx == str.length()){
            return balance == 0;
        }
        char currChar = str.charAt(idx);

        if(currChar == '('){
            return balanceCheck(str, idx +1, balance +1);
        }
        if(currChar == ')'){
            if(balance <= 0){
                return false;
            }
            return balanceCheck(str, idx+1, balance -1);
        }
        return balanceCheck(str, idx+1, balance);

    }
}
