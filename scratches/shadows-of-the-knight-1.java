import java.util.*;
import java.io.*;
import java.math.*;

class Direction {
    public int right = 0;
    public int down = 0;

    // parse from string 
    public Direction(String dir) {
        if (dir.contains("U")) {
            down = -1;
        }
        if (dir.contains("D")) {
            down = 1;
        }
        if (dir.contains("L")) {
            right = -1;
        }
        if (dir.contains("R")) {
            right = 1;
        }
    }
}

class BinarySearch {
    public int leftBoundary;
    public int rightBoundary;

    // 1d binary search
    public BinarySearch(int leftBoundary, int rightBoundary){
        this.leftBoundary = leftBoundary;
        this.rightBoundary = rightBoundary;
    }

    // returns 1d coord of next step
    // if direction > 0 then we need to get closer to the right boundary, < 0 - to the left, == 0 - stop moving
    public int nextStep(int currentPos, int direction) {
        if (direction == 0) {
            return currentPos;
        }

        if (direction < 0) {
            this.rightBoundary = currentPos;
        } 
        if (direction > 0) {
            this.leftBoundary = currentPos;
        }
        return (this.leftBoundary + this.rightBoundary) / 2;
    }
}


class Player {

   static int X;
   static int Y;
     

    // https://www.codingame.com/ide/puzzle/shadows-of-the-knight-episode-1
    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        int W = in.nextInt(); // width of the building.
        int H = in.nextInt(); // height of the building.
        int N = in.nextInt(); // maximum number of turns before game over.
        int X0 = in.nextInt();
        int Y0 = in.nextInt();

        X = X0;
        Y = Y0;

        BinarySearch hSearch = new BinarySearch(0, W);
        BinarySearch vSearch = new BinarySearch(0, H);

        // game loop
        while (true) {
            String bombDir = in.next(); // the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
            Direction dir = new Direction(bombDir);

            X = hSearch.nextStep(X, dir.right);
            Y = vSearch.nextStep(Y, dir.down);

            // the location of the next window Batman should jump to.
            System.out.printf("%d %d\n",X, Y);
        }
    }
}
