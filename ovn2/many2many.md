* What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?
  * I got "panic: send on closed channel". We cannot send to the channel when it's closed, but we didn't wait until the goroutines finished sending stuff before channel closed.

* What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
  * It prints out values from Produce function and then crashed in the end "panic: send on closed channel", which means it tries to send all values in the goroutines that aren't finished yet to a closed channel.

* What happens if you remove the statement close(ch) completely?
  * The program runs normally with time: 980.393302ms. It might be OK without close, since in Go when the channel is no longer used it will be garbage collected anyway. But if we really don't want to send anything to the channel then we should close it.

* What happens if you increase the number of consumers from 2 to 4?
  * The program runs normally with time: 514.929453ms. The time is shorter than before because it prints more often than before.

* Can you be sure that all strings are printed before the program stops?
  * It might not print since we only wait for the producers but not for the consumers, we can test this by using time.Sleep() in the end of the main function and see if there are any strings left that are printed out afterwards.
