package sprint;

public class Factorial {

    public static int calculateFactorial(int n) {
        // solution code here
        if(n < 0){
            return 0;
        }
        if(n == 0 || n == 1){
            return 1;
        }
        return n * calculateFactorial(n -1);
    }

}