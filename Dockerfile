FROM xxx

ADD ./bin/habit-tracking /root/bin/habit-tracking

ENTRYPOINT ["/bin/sh","-c","/root/bin/habit-tracking"]