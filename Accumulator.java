package sprint;

public class Accumulator {
    public int accumulate(int num){
        if(num > 0){
            int result = 0;
            for(int i = 0; i <= num; i++){
                result += i;
            }
            return result;
        }
        return 0;
    }
}
