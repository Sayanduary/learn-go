from pathlib import Path

content = r"""# Go Buffered Channels — Simple Explanation

Your Go program demonstrates **buffered channels, goroutines, and channel closing**.

We will understand it using a real-world example: **an email delivery service**.

Imagine you need to send five emails. Instead of sending each one yourself, you put the email addresses into an inbox. A worker takes addresses from the inbox and sends the emails one by one.

## 1. Real-world analogy

| Go concept | Real-world example |
|---|---|
| `main()` | Manager who provides email addresses |
| `emailChan` | Inbox or basket that stores email addresses |
| `emailSender()` | Worker who sends the emails |
| `done` | Completion notification from the worker |
| Goroutine | Worker running independently |
| Buffered channel | Inbox that can hold several email addresses |

## 2. The complete code

```go
package main

import (
    "fmt"
    "time"
)

func emailSender(emailChan chan string, done chan bool) {
    defer func() {
        done <- true
    }()

    for email := range emailChan {
        fmt.Println("Sending Email to", email)
        time.Sleep(time.Second * 1)
    }
}

func main() {
    email_chan := make(chan string, 100)
    done := make(chan bool)

    go emailSender(email_chan, done)

    for i := 0; i < 5; i++ {
        email_chan <- fmt.Sprintf("%d@gmail.com", i)
    }

    fmt.Println("done sending")
    close(email_chan)

    <-done
}
```

## 3. Understand the code step by step

### Step 1: Create a buffered channel

```go
email_chan := make(chan string, 100)
```

This creates a channel that stores strings and has a buffer capacity of 100.

Think of it as an inbox that can hold up to 100 email addresses.

For example, you can put these addresses into it:

```go
email_chan <- "0@gmail.com"
email_chan <- "1@gmail.com"
email_chan <- "2@gmail.com"
```

All three sends can complete without waiting for the worker to receive them because the buffer has enough free space.

**Important:** A send to a buffered channel blocks only when the buffer is full.

### Step 2: Create a `done` channel

```go
done := make(chan bool)
```

This channel is used to notify `main()` when the email worker has finished.

Think of it as a message from the worker saying:

> My work is finished.

This is an **unbuffered channel**, because no capacity was specified.

### Step 3: Start a goroutine

```go
go emailSender(email_chan, done)
```

The `go` keyword starts `emailSender()` in a new goroutine.

A goroutine is a lightweight unit of execution managed by the Go runtime. It allows the email worker to run concurrently with the main function.

Think of it as assigning the email-sending task to a separate worker.

### Step 4: Put five email addresses into the channel

```go
for i := 0; i < 5; i++ {
    email_chan <- fmt.Sprintf("%d@gmail.com", i)
}
```

This loop sends five strings into `email_chan`:

```text
0@gmail.com
1@gmail.com
2@gmail.com
3@gmail.com
4@gmail.com
```

The channel has capacity 100, so all five addresses can be buffered without waiting for the worker to receive them.

The worker may also receive and process addresses while the loop is running. The buffer does not guarantee that all five will be waiting there at the same time.

### Step 5: Print a message

```go
fmt.Println("done sending")
```

This prints:

```text
done sending
```

But does it mean all emails have been sent?

**No.** It means only that the main function has finished sending all five email addresses into the channel.

The worker may still be processing them.

This distinction is one of the most important things to understand about buffered channels.

### Step 6: Close the channel

```go
close(email_chan)
```

This tells the receiver:

> No more email addresses will be sent.

Closing a channel does **not** delete values already stored in its buffer. The worker can still receive and process those values.

After closing a channel, sending another value to it causes a panic.

### Step 7: Understand the `emailSender()` function

```go
func emailSender(emailChan chan string, done chan bool) {
    defer func() {
        done <- true
    }()

    for email := range emailChan {
        fmt.Println("Sending Email to", email)
        time.Sleep(time.Second * 1)
    }
}
```

Let's break it down.

#### A. The `range` loop

```go
for email := range emailChan {
```

This receives values from the channel one at a time.

The loop continues until the channel is closed **and all buffered values have been received**.

The worker processes each email in turn.

#### B. Simulate sending an email

```go
fmt.Println("Sending Email to", email)
time.Sleep(time.Second * 1)
```

The program prints the email address and then waits one second.

The sleep simulates the time required to send an email. It is not actual email delivery.

With five emails, the worker spends about five seconds sleeping in total.

#### C. Understand `defer`

```go
defer func() {
    done <- true
}()
```

`defer` schedules a function to run when `emailSender()` returns.

Once the worker has processed every buffered email, the `range` loop ends. The deferred function then sends `true` into the `done` channel.

Think of it as the worker telling the manager:

> I have finished processing all the emails.

### Step 8: Wait for the worker to finish

```go
<-done
```

This receives a value from the `done` channel.

Because `done` is unbuffered, the main goroutine waits here until the worker sends `true`.

This prevents `main()` from returning before the worker finishes.

---

## 4. Complete execution flow

1. `main()` creates a buffered channel with capacity 100.
2. `main()` creates an unbuffered `done` channel.
3. `main()` starts the `emailSender()` goroutine.
4. `main()` sends five email addresses into `email_chan`.
5. `main()` prints `done sending`.
6. `main()` closes `email_chan`.
7. The worker receives and processes all remaining email addresses.
8. The `range` loop ends because the channel is closed and empty.
9. The deferred function sends `true` into `done`.
10. `<-done` receives that value, allowing `main()` to finish.

## 5. Expected output

The output will look something like this:

```text
done sending
Sending Email to 0@gmail.com
Sending Email to 1@gmail.com
Sending Email to 2@gmail.com
Sending Email to 3@gmail.com
Sending Email to 4@gmail.com
```

**Note:** The exact order of `done sending` and the first email message can vary. The main goroutine and the worker goroutine run concurrently.

The email messages themselves are processed in the order they are received from this single producer and single receiver.

## 6. Buffered channel vs. unbuffered channel

| Unbuffered channel | Buffered channel |
|---|---|
| Capacity is zero | Capacity is specified |
| Sender waits until a receiver is ready | Sender can proceed while buffer space is available |
| Direct handoff between sender and receiver | Can temporarily store values |
| Useful for synchronization | Useful for producer-consumer workflows |

### Real-world example

- **Unbuffered channel:** You hand an envelope directly to a worker. The handoff happens only when both are ready.
- **Buffered channel:** You put envelopes into an inbox. The worker collects and processes them later.

## 7. Important concepts to remember

### Sending data

```go
email_chan <- "hello@example.com"
```

The `<-` operator sends a value into a channel.

### Receiving data

```go
email := <-email_chan
```

The `<-` operator receives
