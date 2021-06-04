package javastuff;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;

public class lc409 {
	
	public int longestPalindrome(String s) {
		int retVal = 0;
		
		// palindromes require 1 singular and then repeated values afterwards.
		
		HashMap<Character,Integer> map = new HashMap<Character,Integer>();
		
		for(int x = 0; x< s.length(); x++) {
			char ch = s.charAt(x);
			if(map.containsKey(ch)) {
				if(map.get(ch)==1) {
					map.replace(ch, 0);
					retVal+=2;
				}else {
					map.replace(ch, map.get(ch)+1);
				}
				
				
			}else {
				map.put(ch, 1);
			}
		}
		Iterator<?> hmIterator = map.entrySet().iterator();
		boolean hasSingle = false;
		while (hmIterator.hasNext()) {
            Map.Entry mapElement = (Map.Entry)hmIterator.next();
            int mapVal = (int)mapElement.getValue();
            if(mapVal ==1) {
            	hasSingle = true;
            	break;
            }
            
        }
		if(hasSingle) {
			retVal++;
		}
		

		
		return retVal;
		
    }
	
	public int longestPalindromeImproved(String s) {
		int retVal = 0;
		
		// palindromes require 1 singular and then repeated values afterwards.
		char tempArray[] = s.toCharArray();
        
        // sort tempArray
        Arrays.sort(tempArray);
		
		char currVal = tempArray[0];
		int currCount = 1;
		boolean hasSingle = false;
		for(int x = 1; x< tempArray.length; x++) {
			if(currVal == tempArray[x]) {
				currCount++;
				if(currCount%2==0) {
					retVal+=2;
				}
			}else {				
				if(currCount%2 == 1) {
					hasSingle = true;
				}
				currCount = 1;
				currVal = tempArray[x];
			}
			
		}
		if(currCount%2 == 1) {
			hasSingle = true;
		
		}
		if(tempArray.length ==1|| hasSingle) {
			retVal++;
		}
		

		return retVal;
		
    }
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		lc409 test = new lc409();
		System.out.println(test.longestPalindrome("abccccdd"));
		System.out.println(test.longestPalindromeImproved("abccccdd"));
		
	}

}
