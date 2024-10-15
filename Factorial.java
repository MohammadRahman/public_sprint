public class Factorial {
    public static int calculateFactorial(int num){
        if(num < 0){
            return 0;
        }
        if(num ==0 || num == 1){
            return 1;
        }
        return num* calculateFactorial(num -1);
    }
}
