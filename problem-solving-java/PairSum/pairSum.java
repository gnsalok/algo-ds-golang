import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;

public class pairSum {
    public static void main(String[] args) {
        int a[] = { 1, 2, 3, 4, 5 };
        Map<Integer, List<Integer[]>> map = new HashMap<Integer, List<Integer[]>>();
        for (int i = 0; i < a.length; i++) {
            for (int j = i + 1; j < a.length; j++) {
                Integer key = a[i] + a[j];
                Integer[] pair = { a[i], a[j] };
                if (map.containsKey(key)) {
                    List<Integer[]> list = map.get(key);
                    list.add(pair);
                    map.put(key, list);
                } else {
                    List<Integer[]> list = new ArrayList<Integer[]>();
                    list.add(pair);
                    map.put(key, list);
                }
            }
        }
        for (Entry<Integer, List<Integer[]>> en : map.entrySet()) {
            if (en.getValue().size() > 1) {
                for (Integer[] ab : en.getValue()) {
                    System.out.println(Arrays.toString(ab));
                }
            }
        }

    }
}