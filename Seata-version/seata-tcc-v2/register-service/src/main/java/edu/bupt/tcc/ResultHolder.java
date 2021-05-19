package edu.bupt.tcc;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ResultHolder {
    private static Map<Class<?>, Map<String, String>> map = new ConcurrentHashMap<Class<?>, Map<String, String>>();

    public static void setResult(Class<?> actionClass, String xid, String v) {
        Map<String, String> results = map.get(actionClass);

        if (results == null) {
            synchronized (map) {
                if (results == null) {
                    results = new ConcurrentHashMap<>();
                    map.put(actionClass, results);
                }
            }
        }

        results.put(xid, v);
        System.out.println("setResult");
        System.out.println(map.toString());
        System.out.println(results);
    }

    public static String getResult(Class<?> actionClass, String xid) {
        Map<String, String> results = map.get(actionClass);
        System.out.println("getResult");
        System.out.println(map.toString());
        System.out.println(results);
        if (results != null) {
            return results.get(xid);
        }

        return null;
    }

    public static void removeResult(Class<?> actionClass, String xid) {
        Map<String, String> results = map.get(actionClass);
        if (results != null) {
            results.remove(xid);
        }
    }
}