import java.util.*;

public class B_TrickTaking  {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
N
        int N = scanner.nextInt();
        int T = scanner.nextInt();
        int C[] = new int[N];
        int R[] = new int[N];
        for (int i = 0; i < N; i++) {
            C[i] = scanner.nextInt();
        }
        for (int i = 0; i < N; i++) {
            R[i] = scanner.nextInt();
        }
        scanner.close();

        int ans = checkT(C, R, T);
        if (ans != -1) {
            System.out.println(ans);
        } else {
            ans = checkOne(C[0], R[0], C, R);
            System.out.println(ans);
        }
    }

    private static int checkT(int[] C, int[] R, int T) {
        int ans = -1;
        int value = -1;
        for (int i = 0; i < C.length; i++) {
            if (C[i] == T && R[i] > value) {
                value = R[i];
                ans = i+1;
            }
        }
        return ans;
    }

    private static int checkOne(int C1, int R1, int[] C, int[] R) {
        int ans = -1;
        int value = R1;
        for (int i = 1; i < C.length; i++) {
            if (C[i] == C1) {
                if (R[i] > value) {
                    value = R[i];
                    ans = i+1;
                }
            }
        }
        if (ans != -1) {
            return ans;
        } else {
            return 1;
        }
    }
}

