package javastuff;

import java.util.ArrayList;
import java.util.Arrays;

public class Lc1326 {
	
	public int minTaps(int n, int[] ranges) {
		
		ArrayList<int[]> tapRanges = new ArrayList<int[]>();
       
		
		for(int x = 0; x <= n; x++){
			
            tapRanges.add(x, new int[] {x-ranges[x],x+ranges[x]});
            
            
        }
		
		int count = 0;
		int farthestLeft = 0;
		int farthestRight = 0;
		System.out.println(Arrays.deepToString(tapRanges.toArray()));
		while(count< n && farthestRight < n) {
			for(int x =0; x <= n; x++) {
				if(farthestLeft >=tapRanges.get(x)[0] && farthestRight < tapRanges.get(x)[1] ) {
					
					farthestRight = tapRanges.get(x)[1];
					
				}
			}
			count++;
			System.out.println(farthestRight);
			System.out.println(farthestLeft);
			
			// There farthest point is less than n, so the whole garden cannot be reached.
			if(farthestLeft == farthestRight ) {return -1;}
			
			farthestLeft = Math.max(farthestLeft, farthestRight);
		}
	
		
		return count;
    }
	

	
	public static void main(String[]args) {
		
		Lc1326 lcTest = new Lc1326();
		System.out.println(lcTest.minTaps(5, new int[] {3,4,1,1,0,0}));
		
	}
}
