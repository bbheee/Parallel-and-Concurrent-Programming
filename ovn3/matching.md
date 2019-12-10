* What happens if you remove the go-command from the Seek call in the main function?
  * If you don't have the go-command, the seek call will always print out with the same order. But if you have the go-command, it creates the goroutine and everything will happen concurrently, you will get the different result every time you run the program.

* What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg \*sync.WaitGroup to wg sync.WaitGroup?
  * I got the deadlock error. When you change the declaration, the var wg sync.WaitGroup cannot be referred to as a pointer itself, you have to send its address to the function, in this case &wg. But wg := new(sync.WaitGroup) can be referred to as a pointer. When you remove the \* from sync.WaitGroup, it will just send the copy of the WaitGroup object, which means the WaitGroup in the main thread will not know when wg.Done() is called in Seek. The main thread will wait forever for wg.Wait(), which causes a deadlock.

* What happens if you remove the buffer on the channel match?
  * By removing the buffer on the channel match, I got the deadlock error, and I tried to remove the length of the array with one element, but that works. This works whenever I change the array length to a even number. I guess it's because the buffer is used for the unmatched message, but if you have a even number you will not have any unmatched messages. The main thread will wait for the unmatched message to be read if there is no buffer.

* What happens if you remove the default-case from the case-statement in the main function?
  * I got no one received x's message. However if I change the array length to an even number, then it will cause a deadlock, because the case with name := <-match will block, but there is no default option. If there is no default case when the other case isn't ready, there is no way to exit the statement.
