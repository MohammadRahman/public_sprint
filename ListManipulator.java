package sprint;

import java.util.List;

public class ListManipulator {

    public List<String> manipulateList(List<String> list) {
        // Handle the empty list case.
        list.remove(list.size() -1);
        list.set(list.size() -1, "The size of the list is " + list.size());
        list.add("last");
        list.set(0, "first");
        
        return list;
    }
}
