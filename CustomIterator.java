package sprint;

import java.util.Iterator;
import java.util.List;
import java.util.NoSuchElementException;

public class CustomIterator {
    private List<Integer> list;
    private int currentIndex;

    public CustomIterator(List<Integer> list) {
        this.list = list;
        this.currentIndex = 0;
    }

    public boolean hasNext() {
        return currentIndex < list.size();
    }

    public Integer next() {
        if (!hasNext()) {
            throw new NoSuchElementException("No more elements in the list");
        }
        return list.get(currentIndex++);
    }
}
