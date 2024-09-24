package sprint;

public class WordCounter {
    public int countWords(String sentence) {
        // Implement logic here
        if(sentence == null || sentence.isEmpty()){
            return 0;
        }
        String[] words = sentence.split("[^a-zA-Z]+");
        int count = 0;

        for(String word: words){
            if(!word.isEmpty()){
                count++;
            }
        }
        return count;
    }
}
