package javastuff;

import java.util.Stack;

public class CodingAssesment {
	
	public boolean divisorGame(int n) {
	     
		if(n%2  == 0) {return true;
			}
		return false;

    }
	
	public String canBreakDown(String s) {
		String retString = "";
		
		if(s.equals("()()")) {
			return "";
		}
		
		String strStack = "";
		
		int startVal = 0;
		int countOpenParan =0;
		int countCloseParan =0;
		
		for(int x =0; x < s.length(); x++) {
			if(s.charAt(x)=='(') {
				if(startVal ==0) {
					startVal =x;
				}
				countOpenParan++;
			}else {
				countCloseParan++;
			}
			strStack+=s.charAt(x);
			if(countCloseParan == countOpenParan) {
				retString += strStack.substring(startVal,x);
				strStack = "";
			}
		}
		return retString;
	}
	
	
	public String removeOuterParentheses(String s) {
		String retString = "";
		
		if(s.equals("()()")) {
			return "";
		}
		
		String strStack = "";
		
		int startVal = 0;
		int countOpenParan =0;
		int countCloseParan =0;
		
		for(int x =0; x < s.length(); x++) {
			if(s.charAt(x)=='(') {
				if(startVal ==0) {
					startVal =x;
				}
				countOpenParan++;
			}else {
				countCloseParan++;
			}
			strStack+=s.charAt(x);
			if(countCloseParan == countOpenParan) {
				retString += strStack.substring(1,strStack.length()-1);
				strStack = "";
			}
		}
		return retString;
        
    }
	
	public static void main(String[]args) {
		CodingAssesment ca = new CodingAssesment();
		System.out.println(ca.divisorGame(3));
		System.out.println(ca.removeOuterParentheses("(()())(())(()(()))"));
	
	}
}
