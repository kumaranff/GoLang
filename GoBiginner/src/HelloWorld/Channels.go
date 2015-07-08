/*
Channels may be buffered or unbuffered

Unbuffered Channels
- Cannot buffer channel
- The sending routing is locked by the channel until reciving routing take the data
- Once reciving routing gets the data, sending routing is freed.
- Syncronized behavior
myChannel := make(chan int)

Buffered Channels
- Sending channnel is not locked as far as buffered channel can have data
- Asyncroned behavior
myChannel := make(chan int, 5)

If a go routing is waiting for a data from a channel, it's locked as far as
there is no data in the channel, in spite of channel is buffered or unbuffered.
*/