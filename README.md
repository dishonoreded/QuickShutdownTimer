# QuickShutDownTimer
this is a quick shutdown timer for windows

# Usage
1. go build -o qst.exe main.go
2. put `qst.exe` into your windows path
3. exec `qst` in the cmd
  * shutdown on a specify time(shutdown on 2006-01-02 15:04)
    ```
    qst -t 2006-01-02 15:04
    ```
  * shutdown after n hours(shutdown after 1.5 hours)
    ```
    qst -a 1.5
    ```
  * cancel shutdown
    ```
    qst -c
    ```
