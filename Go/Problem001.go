package main
import (
    "fmt"
)

/* Project Euler
 * Problem 001
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we
 * get 3, 5, 6, and 9.  The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 and 5 below 1000.
 *
 * Solution:
 * The problem asks for the sum of all of the elements of the union of the
 * multiples of 3 and 5 that are less that 1000.  So we want
 *     ans = Sum[Multiples3 Union Multiples5]
 * which using Set Theory rules can be changed to
 *     ans = Sum[Multiples3] + Sum[Multiples5] - Sum[Multiples3 Intercept Multiples5]
 * which can be further changed to
 *     ans = Sum[Multiples3] + Sum[Multiples5] - Sum[Multiples15]
 * by using the fact that a multiple of 3 and 5 will also be a multiple of 15.
 * This means that there are three sums to compute: the sum of the multiples of
 * 3, the sum of the multiples of 5, and the sum of the multiples of 15.  A
 * multiple of an integer X is any integer M such that M = N*X for any integer
 * N where N >= 0.  Thus the set of multiples for X is [0*X, 1*X, ..., N*X].
 * A multiple can also be written as
 *     M = N * X = Summation from k = 0 to N of X
 * And so the partial sum of the multiples of X is
 *     P = Sum[Summation from k = 0 to N of X]
 * which achieves O(L) time where 
 *     L = (upper bound - lower bound) / X
 * where / denotes integer division.  For our cases, it means we have a O(333),
 * O(200) and O(67) function call.
 * This solution takes advantage of goroutines, a language feature that is
 * similar to threading or asynchronous function calls where the solution is
 * not required right away, and can be fetched later before it is needed.  This
 * allows the program to queue up all three summations, and on a properly
 * configured machine, allow all three to be calculated concurrently.  Then
 * through the use of another language feature, Channels, the main function
 * gets the summed values and computes the answer using them.  While with small
 * bounds like we have, optimizations like these are not likely to improve
 * running time by a significant amount, this program is scalable and with the
 * change of parameters, can compute the summations of much higher values.
 *
 * Possible improvements:
 * The most obvious solution is to make the program take in values from the
 * command line to determine the bounds and which values we are summing the
 * multiples of.
 */

 // NOTE: start and end are inclusive!
func partialsum(start, step, end int, c chan int) {
    sum := 0
    for i := start; i <= end; i += step {
        sum += i
    }
    c <- sum
}

func main() {
    c := make(chan int)             // Channel for sum of factors of 3 and 5
    m := make(chan int)             // Channel for sum of factors of 15
    go partialsum(0, 3, 999, c)     // Sums all of the factors of 3
    go partialsum(0, 5, 999, c)     // Sums all of the factors of 5
    go partialsum(0, 15, 999, m)    // Sums all of the factors of 15
    x, y, z := <-c, <-c, <-m
    ans := x + y - z                // s(m3) + s(m5) - s(m15)
    fmt.Printf("%d", ans)
}