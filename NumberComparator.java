package sprint;

public class NumberComparator {
    public String whichIsGreater(int n, double f) {
        if(n > f){
            return "Integer";
        }else if((double)n == f){
            return "Equal";
        }else{
            return "Float";
        }
    }
}
