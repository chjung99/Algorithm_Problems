class RandomizedSet {
    Map<Integer, Integer> locations;
    List<Integer> values;
    int size;

    public RandomizedSet() {
        locations = new HashMap<>();
        values = new ArrayList<>();
        size = 0;
    }
    
    public boolean insert(int val) {
        if (locations.containsKey(val)) {
            return false;
        }

        locations.put(val, size);
        values.add(val);
        size += 1;

        return true;
    }
    
    public boolean remove(int val) {
        if (!locations.containsKey(val)) {
            return false;
        }
        int loc = locations.get(val);
        int lastVal = values.get(size-1);

        Collections.swap(values, loc, size-1);
        locations.put(lastVal, loc);

        locations.remove(val);
        values.remove(size-1);
        size -= 1;

        return true;
    }
    
    public int getRandom() {
        return values.get(new Random().nextInt(values.size()));
    }
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * RandomizedSet obj = new RandomizedSet();
 * boolean param_1 = obj.insert(val);
 * boolean param_2 = obj.remove(val);
 * int param_3 = obj.getRandom();
 */