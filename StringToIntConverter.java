package sprint;

import java.util.List;
import java.util.stream.Collectors;

public class StringToIntConverter {
    public List<Integer> convertStringListToIntList(List<String> stringList){
        return stringList.stream()
                .map(Integer:: parseInt)
                .collect(Collectors.toList());
    }
}
