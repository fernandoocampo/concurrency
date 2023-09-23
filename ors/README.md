# ors

## Or Done pattern

* channels from disparate parts of the application.
* It's unknown if the channel you're reading from has been canceled when you're goroutine is cancelled as well.
* we need wrap our read from the channel with a select statement.