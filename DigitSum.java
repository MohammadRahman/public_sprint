package sprint;

public class DigitSum {
    public static int sumOfDigits(int number) {
        // Implement logic here
        int sum = 0;
        number = absNumber(number);

        while (number > 0) {
            int digit = number % 10;
            sum += digit;
            number /= 10;
        }
        return sum;
    }

    private static int absNumber(int num){
        if(num < 0){
            return -num;
        }
        return num;
    }
}
