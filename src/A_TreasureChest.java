import java.util.*;

public class A_TreasureChest {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int N = scanner.nextInt();
        char s[] = scanner.next().toCharArray();
        scanner.close();

        int count = 0;
        for (int i = 0; i < N; i++) {
            if(s[i] == '|') {
                count++;
            }
            if(s[i] == '*' && count == 1) {
                System.out.println("in");
                System.exit(0);
            }
            if(count == 2) {
                System.out.println("out");
                System.exit(0);
            }
        }
    }
}

